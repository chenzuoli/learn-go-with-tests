package main

import (
	"fmt"
	"sync"
)

type rect struct {
	lenth int
	width int
}

func (a rect) area(wg *sync.WaitGroup) int {
	defer wg.Done() // this code style is really readable and easy to understand.
	if a.lenth < 0 {
		fmt.Println("rectagular's length is not allowed to be under zero.")
		//wg.Done()
		return 0
	}
	if a.width < 0 {
		fmt.Println("rectagular's width is not allowed to be under zero.")
		//wg.Done()
		return 0
	}
	fmt.Println("rectagular's area is ", a.lenth*a.width)
	//wg.Done()
	return a.lenth * a.width
}

func main() {
	var wg sync.WaitGroup
	var recta1 = rect{4, 5}
	var recta2 = rect{2, 3}
	var recta3 = rect{10, 30}
	var a = []rect{recta1, recta2, recta3}
	for _, recta := range a {
		wg.Add(1)
		go recta.area(&wg)
	}
	wg.Wait()
	fmt.Println("all routines is finished.")
}
