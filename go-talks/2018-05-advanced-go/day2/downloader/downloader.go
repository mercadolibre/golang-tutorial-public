package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		log.Printf("usage: %s [concurrency] [url] [output]", os.Args[0])
		return
	}

	workers, err := strconv.ParseInt(os.Args[1], 10, 64)
	if err != nil {
		log.Fatalf("error parsing concurrency param: %v", err)
	}

	resourceURL, err := url.Parse(os.Args[2])
	if err != nil {
		log.Fatalf("error parsing url param: %v", err)
	}

	out, err := os.OpenFile(os.Args[3], os.O_RDWR|os.O_CREATE|os.O_EXCL, os.ModePerm)
	if err != nil {
		log.Fatalf("error opening file for writing: %v", err)
	}

	res, err := http.Head(resourceURL.String())
	if err != nil {
		log.Fatalf("error requesting HEAD of file: %v", err)
	}

	if res.StatusCode >= 300 {
		log.Fatalf("unexpected status code received from server: %s", res.Status)
	}

	acceptRange := res.Header.Get("accept-ranges")
	if acceptRange != "bytes" && workers > 1 {
		log.Fatalf("remote server does not accept range downloads")
	}

	contentLength, err := strconv.ParseInt(res.Header.Get("content-length"), 10, 64)
	if err != nil || contentLength == 0 {
		log.Fatalf("remote server content-length is invalid")
	}

	chunkSize := contentLength / workers
	for i := int64(0); i < workers; i++ {
		chunkStart, chunkEnd := chunkSize*i, (chunkSize*i)+chunkSize-1
		if i+1 == workers {
			chunkEnd = contentLength
		}

		rangeHeader := fmt.Sprintf("bytes=%d-%d", chunkStart, chunkEnd)
		_ = rangeHeader
	}

	_ = workers
	_ = resourceURL
	_ = out
	_ = chunkSize
}
