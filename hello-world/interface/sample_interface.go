package main

import (
	"fmt"
)

// interface defination
type VowelsFinder interface {
	FindVowels() []rune
}

type MyString string

// implement the interface
func (ms MyString) FindVowels() []rune {
	var vowels []rune
	for _, rune := range ms {
		if rune == 'a' || rune == 'e' || rune == 'i' || rune == 'o' || rune == 'u' {
			vowels = append(vowels, rune)
		}
	}
	return vowels
}

func main() {
	ms := MyString("chenzuoli")
	var findvowels VowelsFinder
	findvowels = ms
	fmt.Printf("vowels are %c", findvowels.FindVowels())
}
