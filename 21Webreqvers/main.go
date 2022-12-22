package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Hello World")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "http://3.109.139.40:3000/ping"

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

}
