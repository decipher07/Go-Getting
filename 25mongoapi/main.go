package main

import (
	"fmt"
	"log"
	"net/http"

	router "github.com/decipher07/mongoapi/routers"
)

func main() {
	fmt.Println("MONGODB API")
	r := router.Router()
	log.Fatal(http.ListenAndServe(":3000", r))
	fmt.Println("Server is running at port : 3000")
}
