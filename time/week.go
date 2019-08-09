package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(today)
	fmt.Println(time.Saturday)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Print("Tomorrow is ")
		fmt.Println(today + 1)
	case today + 2:
		fmt.Print("In two days is ")
		fmt.Println(today + 2)
	default:
		fmt.Println("Too far away.")
	}
}
