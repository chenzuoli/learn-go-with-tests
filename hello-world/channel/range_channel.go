package main

import "fmt"

func into_channel(channel chan int) {
	for i := 1; i < 10; i++ {
		channel <- i
	}
	close(channel) // close channel. if you don't close the channel, then the program will be panic into deadlock.
}

func main() {
	var channel chan int = make(chan int)
	go into_channel(channel)
	for v := range channel { // when the channel is closed, then the for cicle will be break out.
		fmt.Println(v)
	}
}
