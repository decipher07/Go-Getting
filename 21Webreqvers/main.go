package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const myurl = "http://3.109.139.40:3000/"

func main() {
	fmt.Println("Hello World")
	// PerformGetRequest()
	performPostRequest()
}

func PerformGetRequest() {
	myurl := myurl + "ping"
	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	fmt.Println("The response status code is ", response.StatusCode)
	fmt.Println("Content Length is: ", response.ContentLength)

	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	byteContent, _ := responseString.Write(content)

	// fmt.Println(string(content))
	fmt.Println("ByteCount is: ", byteContent)
	fmt.Println(responseString.String())

	defer response.Body.Close()
}

func performPostRequest() {
	pathToPost := myurl + "payment/makepayment"

	requestBody := strings.NewReader(`
		{
			"userId": "98e3c010-1573-4d44-a23a-323a0dd25f56",
			"month": 11
		}
	`)

	response, err := http.Post(pathToPost, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(content))
}
