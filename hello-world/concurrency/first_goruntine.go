package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Hello World.")
}

func main() {
	go sayHello()
	//time.Sleep(1*time.Second)
	fmt.Println("main function.") // you usually get the result "main function." and no string "Hello World.", because go runtine is not run when the main runtine is dead.
}
