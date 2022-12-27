package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels in Go")

	myChan := make(chan int, 1)
	wg := &sync.WaitGroup{}

	wg.Add(2)

	/* <-chan means This is going out of channel */
	/* chan<- means This is going into the channel */

	/* Receive Only == Properly Focus on the Channel Arrows */
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-myChan

		fmt.Println(val)
		fmt.Println(isChannelOpen)

		fmt.Println(<-myChan)
		fmt.Println(<-myChan)

		wg.Done()
	}(myChan, wg)

	/* Send Only == Properly Focus on the Channel arrows */
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myChan <- 1
		myChan <- 2
		close(myChan)
		wg.Done()
	}(myChan, wg)

	wg.Wait()
}
