package main

import (
	"fmt"
	"time"
)

func server1(channel chan string) {
	time.Sleep(time.Second * 4)
	channel <- "come from server1."
}
func server2(channel chan string) {
	time.Sleep(time.Second * 2)
	channel <- "come from server2."
}

func main() {
	var ch1 = make(chan string)
	var ch2 = make(chan string)
	go server1(ch1)
	go server2(ch2)
	for {
		time.Sleep(time.Second * 1)
		select {
		case s1 := <-ch1:
			fmt.Println(s1)
			return
		case s2 := <-ch2:
			fmt.Println(s2)
			return
		default:
			fmt.Println("no response.")
		}
	}
}
