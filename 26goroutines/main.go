package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	/* 	go greeter("Hello")
	   	greeter("World") */

	websiteList := []string{
		"http://lco.dev",
		"http://google.com",
		"http://github.com",
		"http://fb.com",
	}

	for _, web := range websiteList {
		wg.Add(1)
		go getStatusCode(web)
	}

	wg.Wait()
}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}

func getStatusCode(endpoint string) {
	res, _ := http.Get(endpoint)
	fmt.Printf("%d Status Code for %s\n", res.StatusCode, endpoint)
	defer wg.Done()
}
