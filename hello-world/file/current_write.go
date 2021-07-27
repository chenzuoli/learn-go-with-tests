package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sync"
)

// 使用channel信道来解决并发写入文件的问题，多协程生产数据到信道，单信道写入数据

func produce(data chan int, wg *sync.WaitGroup) {
	data <- rand.Intn(999)
	wg.Done()
}

func consume(done chan bool, data chan int) {
	file, e := os.Create("hello-world/file/concurrent.txt")
	if e != nil {
		log.Fatal(e)
		done <- false
		return
	}
	for randInd := range data {
		_, e = fmt.Fprintln(file, randInd)
		if e != nil {
			log.Fatal(e)
			done <- false
			return
		}
	}
	done <- true
}

func main() {
	done := make(chan bool, 1)
	data := make(chan int, 100)
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ { // 必须先将group容量扩大到100，然后再生产数据，生产一条，容量减少1，不然会造成死锁。
		wg.Add(1)
		go produce(data, &wg)
	}
	go consume(done, data)
	go func() {
		wg.Wait()
		close(data)
	}()
	flag := <-done
	if flag {
		fmt.Println("write successful.")
	} else {
		fmt.Println("write failed.")
	}
}
