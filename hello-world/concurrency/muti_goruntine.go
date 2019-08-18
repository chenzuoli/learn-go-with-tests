package main

import (
	"fmt"
	"time"
)

func number() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Microsecond)
		fmt.Println(i)
	}
}
func alphabet() {
	arr := []string{"a", "b", "c", "d", "e"}
	for i := 1; i < len(arr); i++ {
		time.Sleep(400 * time.Microsecond)
		fmt.Println(arr[i])
	}
}

func main() {
	go number()
	go alphabet()
	time.Sleep(3000 * time.Microsecond)
	fmt.Println("main rountine terminated.")
}
