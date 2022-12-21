package main

import (
	"fmt"
	"net/url"
)

const myurl = "http://3.109.139.40:5173/login?myform=something&otherform=nothing"

func main() {
	fmt.Println("Hello World ")
	// fmt.Println(myurl)

	result, _ := url.Parse(myurl)

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Path)
	// fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("The type of query is %T\n", qparams)
	fmt.Println(qparams["myform"])

	for key, values := range qparams {
		fmt.Println("The Key is ", key, " and value is ", values)
	}

	partsOfUrl := &url.URL{
		Scheme:   "http",
		Host:     "3.109.40.5173",
		Path:     "/login",
		RawQuery: "myform=something",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
