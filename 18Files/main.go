package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Hello World")

	content := "This needs to go in a file - Deepankar"

	file, err := os.Create("myfile.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("The length is ", length)

	defer file.Close()

	readFile("myfile.txt")
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(filename string) {
	dataBytes, err := ioutil.ReadFile(filename)

	checkNilError(err)

	fmt.Println("Get the content as : ", string(dataBytes))
}
