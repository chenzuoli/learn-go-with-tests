package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	fphr := flag.String("fpath", "test.txt", "file path to read from")
	flag.Parse()

	file, err := os.Open(*fphr)
	// handle the error
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	lines := bufio.NewScanner(file)
	for lines.Scan() {
		fmt.Println(lines.Text())
	}
	err = lines.Err()
	if err != nil {
		log.Fatal(err)
	}
}
