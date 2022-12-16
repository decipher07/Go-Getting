package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "HELLO WORLD"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("RATING: ")

	input, err := reader.ReadString('\n')
	fmt.Println("Thanks for the rating, ", input)
	fmt.Println(err)

}
