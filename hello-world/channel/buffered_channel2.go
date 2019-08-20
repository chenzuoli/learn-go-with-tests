package main

import (
	"fmt"
	"time"
)

func write(channel chan int) {
	for i := 0; i < 5; i++ {
		channel <- i
		fmt.Printf("write %d to channel.\n", i)
	}
}

func main() {
	var channel = make(chan int, 2)
	go write(channel)
	time.Sleep(time.Second * 2)
	for i := 1; i <= 5; i++ {
		a := <-channel
		fmt.Printf("read from channel: %d\n", a)
		time.Sleep(time.Second * 2)
	}
}
