package main

import "fmt"

func main() {

	fmt.Println("MAPS IN GO LANG")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	languages["Go"] = "Golang"

	fmt.Println(languages)
	delete(languages, "Go")

	for _, value := range languages {
		fmt.Printf("For Key v, value is %v\n", value)
	}
}
