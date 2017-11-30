package main

import (
	"github.com/mercadolibre/golang-restclient/rest"
	"fmt"
	"net/http"
	"encoding/json"
	"errors"
	"sync"
)

const MAX_ITEMS_PER_QUERY = 100
const ITEM_FETCHERS_PER_QUERY = 20

type SearchResult struct {
	Paging PageInfo     `json:"paging"`
	Results []SeachItem `json:"results"`
}

type PageInfo struct {
	Total int `json:"total"`
	Offset int `json:"offset"`
	Limit int `json:"limit"`
}

type SeachItem struct {
	Id string `json:"id"`
}

type ItemInfo struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Price float64 `json:"price"`
	Location SellerAddress `json:"seller_address"`
}

type SellerAddress struct {
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func searchItems(query string) ([]*ItemInfo,error){
	var inputChannel chan string = make(chan string, 20)
	var outputChannel chan *ItemInfo = make(chan *ItemInfo, 20)
	wgProducers := &sync.WaitGroup{}
	wgConsumers := &sync.WaitGroup{}

	result:=make([]*ItemInfo,0)

	// item info producers
	for i:=0;i< ITEM_FETCHERS_PER_QUERY;i++ {
		//increment running producers
		wgProducers.Add(1)
		go itemFetcher(inputChannel, outputChannel, wgProducers)
	}

	//item info consumer
	go func() {
		//iterate until outputChannel is closed
		for itemInfo:=range outputChannel {
			result=append(result,itemInfo)
			//decrements pending items
			wgConsumers.Done()
		}
	}()

	var err error = nil
	offset:=0
	total:=MAX_ITEMS_PER_QUERY

	for offset <  total {
		url:=fmt.Sprintf("https://api.mercadolibre.com/sites/MLA/search?q=%s&offset=%v", query,offset)
		resp := rest.Get(url)
		if resp.Response != nil && resp.Response.StatusCode == http.StatusOK {
			var searchResult SearchResult
			err=json.Unmarshal(resp.Bytes(), &searchResult)
			if err!=nil {
				//interrupt search
				break
			}
			if(searchResult.Paging.Total<MAX_ITEMS_PER_QUERY){
				total=searchResult.Paging.Total
			}
			for _,r := range searchResult.Results {
				//increments pending items
				wgConsumers.Add(1)
				inputChannel <- r.Id
			}
			offset+=searchResult.Paging.Limit
		}else {
			//interrupt search
			err= errors.New(fmt.Sprintf("Error calling to %s",url))
			break
		}
	}

	//items ids have been retrieved. input channel should be closed in order to let producers finish
	close(inputChannel)
	//wait until all item info producers have finished
	wgProducers.Wait()
	//wait until all item info are received by consumer
	wgConsumers.Wait()
	//close output channel in order to finish consumer goroutine
	close(outputChannel)

	return result,err
}

//Items Info producer
func itemFetcher(inputChannel <-chan string, outputChanel chan<- *ItemInfo,wgProducers *sync.WaitGroup) {
	//iterate until inputChannel is closed
	for id:=range inputChannel {
		resp := rest.Get(fmt.Sprintf("https://api.mercadolibre.com/items/%s?attributes=id,title,price,seller_address.latitude,seller_address.longitude", id))
		var itemInfo ItemInfo
		if resp.Response != nil && resp.Response.StatusCode == http.StatusOK {
			err:=json.Unmarshal(resp.Bytes(), &itemInfo)
			if err!=nil {
				itemInfo.Id=id
			}
		}else {
			itemInfo.Id=id
		}
		outputChanel<-&itemInfo
	}
	//release producers waitgroup
	wgProducers.Done()
}