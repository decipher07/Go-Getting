package main

import "fmt"

const NumberOfUsers int = 3000

func main() {

	fmt.Println("Hello World From Variables")

	var username string = "DJ"
	fmt.Println(username)
	fmt.Printf("The Variable is of Type %T\n", username)
	fmt.Println("")

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("The Variable is of Type %T\n", isLoggedIn)
	fmt.Println("")

	var smallValue uint8 = 255
	fmt.Println(smallValue)
	fmt.Printf("The Variable is of Type %T\n", smallValue)
	fmt.Println("")

	var smallFloat float32 = 255.2646565686862353535
	fmt.Println(smallFloat)
	fmt.Printf("The Variable is of Type %T\n", smallFloat)
	fmt.Println("")

	var anotherVariable string
	fmt.Println(anotherVariable)
	fmt.Printf("The Variable is of Type %T\n", anotherVariable)
	fmt.Println("")

	var website = "www.google.com"
	fmt.Println(website)

	fmt.Println(NumberOfUsers)
}
