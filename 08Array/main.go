package main

import (
	"fmt"
)

func main() {
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Grapes"

	fmt.Println("The fruit list is ", fruitList)
	fmt.Println("The fruit list is len ", len(fruitList))

	var veglist = [5]string{"1221", "5", "4"}
	fmt.Println("Veglist is ", veglist)
	fmt.Println("Veglist is ", len(veglist))
}
