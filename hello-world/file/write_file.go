package main

import (
	"fmt"
	"log"
	"os"
)

// 写入到项目根目录下了，文件名为lines.txt
func main() {
	file, e := os.Create("lines.txt")
	if e != nil {
		log.Fatal(e)
		return
	}
	var lines = []string{"Welcome to the world of Go.", "Go is a compiled language.", "It is easy to learn Go."}
	for _, line := range lines {
		_, e = fmt.Fprintln(file, line)
		if e != nil {
			log.Fatal(e)
			return
		}
	}
	e = file.Close()
	if e != nil {
		log.Fatal(e)
		return
	}
	fmt.Println("write file successful.")
}
