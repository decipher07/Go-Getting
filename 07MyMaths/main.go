package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	fmt.Println("Hello World")

	/* Random from maths */
	/* rand.Seed(time.Now().Unix())
	fmt.Println(rand.Intn(5) + 1) */

	/* Random from Crypto */
	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
