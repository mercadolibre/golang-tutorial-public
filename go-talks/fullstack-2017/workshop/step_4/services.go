package main

import (
	"github.com/mercadolibre/golang-restclient/rest"
	"fmt"
	"net/http"
	"encoding/json"
	"errors"
)

const MAX_ITEMS_PER_QUERY = 100

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
}

func searchItems(query string) ([]*ItemInfo,error){
	result:=make([]*ItemInfo,0)

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
				result=append(result,&ItemInfo{r.Id})
			}
			offset+=searchResult.Paging.Limit
		}else {
			//interrupt search
			err= errors.New(fmt.Sprintf("Error calling to %s",url))
			break
		}
	}

	return result,err
}