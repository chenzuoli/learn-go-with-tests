package main

import "fmt"

type person string

// defer stack feature: Last in First out(LIFO)
func (a person) printName() {
	for _, letter := range []rune(a) {
		defer fmt.Print("%c", letter)
	}
}

func main() {
	var p = person("12345")
	fmt.Println(p)
	p.printName()
}
