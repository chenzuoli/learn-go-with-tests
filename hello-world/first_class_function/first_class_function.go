package main

import "fmt"

func main() {
	a := func() {
		fmt.Println("I am a first class function.")
	}
	a() // run the function

	fmt.Println("First class function %T test.", a)
	fmt.Printf("%T", a) // get the variable a's type, it is a function.

	fmt.Println("--------")
	func(n string) {
		fmt.Println("i am a anonymous function ", n)
	}("chenzuoli")

}
