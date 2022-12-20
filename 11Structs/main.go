package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var dj User = User{"DJ", "dsaf@gmail.com", false, 12}
	fmt.Println(dj)
	fmt.Printf("DJ is %+v\n", dj)
	fmt.Printf("DJ is %v and %v\n", dj.Name, dj.Age)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
