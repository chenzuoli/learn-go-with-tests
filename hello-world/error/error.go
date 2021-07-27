package main

import (
	"fmt"
	"os"
)

func main() {
	var file, err = os.Open("/text.txt") // open a not exists file.
	if err != nil {
		fmt.Println("open file error.", err)
		return
	}
	fmt.Println(file.Name(), " open successfully.")
}
