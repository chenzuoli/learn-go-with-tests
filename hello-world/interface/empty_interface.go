package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("type is  %T, value is %v\n", i, i)
}

func main() {
	s := "Hello world."
	i := 1
	v := struct {
		name   string
		gender bool
	}{
		"chenzuoli",
		true,
	}
	describe(s)
	describe(i)
	describe(v)
}
