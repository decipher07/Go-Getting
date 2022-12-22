package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Hello World")
	EncodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJs", 299, "learncodeonline.in", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "learncodeonline.in", "321abc", []string{"full stack", "js"}},
		{"Angular", 499, "learncodeonline.in", "321abc", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Println(string(finalJson))
}
