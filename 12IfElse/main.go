package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular"
	} else if loginCount > 10 {
		result = "Nothing"
	} else {
		result = "What not"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

	if num := 13; num < 10 {
		fmt.Println("Number Less than 10")
	} else {
		fmt.Println("Number greater than 10")
	}
}
