package main

import (
	"fmt"
	"sync"
)

var x = 0

func increment(wg *sync.WaitGroup, mutex *sync.Mutex) { // must transfor the pointer of the mutex.
	mutex.Lock()
	x = x + 1
	mutex.Unlock()
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg, &mutex)
	}
	wg.Wait()
	fmt.Printf("x's value is %d", x)
}
