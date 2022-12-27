package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello World from Mutex")

	score := []int{}
	wg := &sync.WaitGroup{}
	m := &sync.RWMutex{}

	wg.Add(1)
	/* R One */
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One R")
		m.Lock()
		score = append(score, 1)
		m.Unlock()
		wg.Done()
	}(wg, m)

	wg.Add(1)
	/* R Two */
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two R")
		m.Lock()
		score = append(score, 2)
		m.Unlock()
		wg.Done()
	}(wg, m)

	wg.Add(1)
	/* R Three */
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R")
		m.Lock()
		score = append(score, 3)
		m.Unlock()
		wg.Done()
	}(wg, m)

	wg.Add(1)
	/* R Three */
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Four R")
		m.RLock()
		fmt.Println(score)
		m.RUnlock()
		wg.Done()
	}(wg, m)

	wg.Wait()
	fmt.Println(score)
}
