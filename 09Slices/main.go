package main

import (
	"fmt"
	"sort"
)

func main() {

	var fruitList = []string{"Apple", "Peach", "Banana"}
	// fmt.Printf("Type is %T \n", fruitList)

	// fmt.Println(fruitList)
	fruitList = append(fruitList, "Mango", "Grapes")
	// fmt.Println(len(fruitList))
	// fmt.Println(fruitList)
	fruitList = append(fruitList[:3])
	// fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 123
	highScores[1] = 987
	highScores[2] = 456
	highScores[3] = 354

	highScores = append(highScores, 12, 412, 1223)

	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	// fmt.Println(highScores)

	course := []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(course)
	index := 2
	course = append(course[:index], course[index+1:]...)
	fmt.Println(course)
}
