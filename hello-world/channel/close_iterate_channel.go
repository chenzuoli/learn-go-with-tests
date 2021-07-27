package main

import "fmt"

func write_channel(channel chan int) {
	for i := 1; i < 10; i++ {
		channel <- i
	}
	close(channel) // close channel.
}

func main() {
	var channel chan int = make(chan int)
	go write_channel(channel)
	for { // iterate the channel.
		data, ok := <-channel
		if ok {
			fmt.Println(data, ok)
		} else {
			break
		}
	}
}
