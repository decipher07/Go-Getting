package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup /* This must be a pointer */
var mut sync.Mutex    /* This must be a pointer */
var signals = []string{"test"}

func main() {
	/* 	go greeter("Hello")
	   	greeter("World") */

	websiteList := []string{
		"http://lco.dev",
		"http://google.com",
		"http://github.com",
		"http://fb.com",
		"http://instagram.com",
	}

	for _, web := range websiteList {
		wg.Add(1)
		go getStatusCode(web)
	}

	wg.Wait()
	fmt.Println(signals)
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

	/* Mutex Lock */
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()
	/* Mutex Unlock */

	defer wg.Done()
}
