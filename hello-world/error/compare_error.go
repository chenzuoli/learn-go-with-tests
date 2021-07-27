package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	matches, err := filepath.Glob("[")
	if err != nil && err == filepath.ErrBadPattern {
		fmt.Println(err)
		return
	}
	fmt.Println(matches)
}
