package main

import "fmt"

// type switch
func findType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Printf("I am a string and my value is %s\n", i.(string))
	case int:
		fmt.Printf("I am a int and my value is %d\n", i.(int))
	default:
		fmt.Println("Unknown type.")
	}
}

func main() {
	s := "chenzuoli"
	i := 1
	v := struct {
		name   string
		gender bool
	}{
		name:   "chenzuoli",
		gender: true,
	}
	findType(s)
	findType(i)
	findType(v)
}
