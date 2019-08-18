package main

import "fmt"

func hello(somewords chan string) {
	fmt.Println("say something.")
	somewords <- "I like you." // 往信道中写入数据（一个字符串） ，如果注释掉这句，没有程序向信道中写入数据，那么程序会出现panic
}

func main() {
	var somewords = make(chan string) // 初始化一个信道
	go hello(somewords)
	<-somewords // 从信道中读取数据，允许从信道中读取数据赋值给空变量。如果没有数据往somewords协程中写入的话，主协程就会一直处于阻塞状态。
	fmt.Println("main routine.")
}
