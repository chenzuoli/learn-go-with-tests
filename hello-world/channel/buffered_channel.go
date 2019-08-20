package main

import "fmt"

func main() {
	var buffered_channel = make(chan string, 2)
	buffered_channel <- "1"
	fmt.Println("add first string.")
	buffered_channel <- "2"
	fmt.Println("add second string.")
	//buffered_channel <- "3" // there is a deadlock because the channel is full.
	fmt.Println(<-buffered_channel) // read one element from buffered channel
}
