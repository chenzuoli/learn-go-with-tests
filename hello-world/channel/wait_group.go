package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("start goroutine ", i)
	time.Sleep(time.Second * 2)
	fmt.Println("end goroutine ", i)
	wg.Done()
}

func main() {
	num := 3
	var wg sync.WaitGroup
	for i := 0; i < num; i++ {
		wg.Add(1)
		go process(i, &wg) // 这里需要传递指针过去，不然会是waitgroup的副本
	}
	wg.Wait()
	fmt.Println("all routines finished.")
}
