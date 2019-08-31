package main

import (
	"fmt"
	"runtime/debug"
)

// recover function only recover the one routine that panicked call it.
func recoverFullName() {
	if r := recover(); r != nil {
		fmt.Println("recoverd from ", r)
		debug.PrintStack() // print the stack strace.
	}
}

func printFullName(firstName *string, lastName *string) {
	defer recoverFullName()
	if firstName == nil {
		panic("########## runtime error: first name is empty")
	}
	if lastName == nil {
		panic("########## runtime error: last name is empty")
	}
	fmt.Println("########## full name is ", firstName, lastName)
}

func main() {
	firstName := "chen"
	printFullName(&firstName, nil)
	fmt.Println("########## print the full name done.")
}
