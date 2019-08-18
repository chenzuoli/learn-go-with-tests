package main

import "fmt"

func sendData(send chan<- int) { // single direction channel: to channel， 简称唯送信道
	fmt.Println("send data to channel.")
	send <- 10
}

func main() {
	var channel = make(chan int) // 双向通道
	go sendData(channel)         // 在方法内变为单向通道
	data := <-channel            // 又转变为双向通道，读取数据
	fmt.Println("get data from channel: ", data)
}
