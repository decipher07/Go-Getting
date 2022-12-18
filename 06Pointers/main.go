package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// var ptr *int
	// fmt.Println("The pointer is ", ptr)

	myNumber := 23
	var ptr = &myNumber

	fmt.Println("The Value of ptr is ", ptr)
	fmt.Println("The Value of ptr is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("The value of number is ", myNumber)
}
