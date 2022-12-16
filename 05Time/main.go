package main

import (
	"fmt"
	"time"
)

func main() {
	welcome := "WELCOME "
	fmt.Println(welcome)

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("02-01-2006 15:04:05 Monday"))

	createdDate := time.Date(2022, time.January, 12, 23, 11, 31, 00, time.UTC)
	fmt.Println(createdDate)
}
