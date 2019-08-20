package main

import "fmt"

func main() {
	var channel = make(chan string, 3)
	channel <- "1"
	channel <- "2"
	fmt.Println(len(channel), cap(channel))
	fmt.Println(<-channel)
	fmt.Println(len(channel), cap(channel))
	fmt.Println(<-channel)
	fmt.Println(len(channel), cap(channel))
}
