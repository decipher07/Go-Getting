package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url string = "https://lco.dev"

func main() {
	fmt.Println("Hello World")

	response, err := http.Get(url)

	manageError(err)
	fmt.Printf("The Response is of type %T\n", response)

	defer response.Body.Close()

	databytes, err := ioutil.ReadAll(response.Body)

	manageError(err)
	content := string(databytes)

	fmt.Println("The Content is ", content)
}

func manageError(err error) {

	if err != nil {
		panic(err)
	}
}
