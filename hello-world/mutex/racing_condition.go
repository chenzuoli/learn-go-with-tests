package main

import (
	"fmt"
	"sync"
)

var i = 0

// use the limited buffered channel to solve the racing condition problem.
// we should choose the right way to solve the racing condition problem looking at the situations if you haven's use the channel, then you should use the mutex.
func increment(wg *sync.WaitGroup, channel chan bool) {
	channel <- true
	i = i + 1
	<-channel
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var channel = make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go increment(&wg, channel)
	}
	wg.Wait()
	fmt.Println("i's value is ", i)
}
