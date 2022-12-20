package main

import "fmt"

func main() {
	fmt.Println("Hello world")

	// days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

	// fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for _, day := range days {
	// 	fmt.Printf("Index is v, and value is %v\n", day)
	// }

	rouge_value := 1

	for rouge_value < 10 {
		if rouge_value == 2 {
			fmt.Println("Hello World")
		}

		if rouge_value == 5 {
			goto lco
		}

		fmt.Println("The value is ", rouge_value)
		rouge_value++
	}

lco:
	fmt.Println("Hello WORLD LCO")
}
