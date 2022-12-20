package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	greeter()

	result := adder(2, 4)
	fmt.Println("The result is ", result)

	fmt.Println("The Proresults are ", proAdder(1, 2, 3, 4, 41, 41))
}

func greeter() {
	fmt.Println("Namastey Golang")
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(vals ...int) int {
	total := 0
	// for i := range vals {
	// 	total += vals[i]
	// }

	// for d := 0; d < len(vals); d++ {
	// 	total += vals[d]
	// }

	// d := 0
	// for d < len(vals) {
	// 	total += vals[d]
	// 	d++
	// }

	// for _, val := range vals {
	// 	total += val
	// }

	return total
}
