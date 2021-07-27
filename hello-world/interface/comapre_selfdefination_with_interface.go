package main

import "fmt"

type Person interface {
	createName() string
}

type Student struct {
	name   string
	school string
}

type Engineer struct {
	name    string
	company string
}

func (name Student) createName() string {
	return "chenyi"
}
func (name Engineer) createName() string {
	return "chenzuoli"
}

func findPersonType(person Person) {
	switch person.(type) {
	case Student:
		fmt.Printf("He is a %T, and he is in %v\n", person, person)
	case Engineer:
		fmt.Printf("He is a %T, and he is in %v\n", person, person)
	default:
		fmt.Printf("Unknown this guy.\n")
	}
}

func main() {
	var student Student = Student{"", "first school"}
	var engineer Engineer = Engineer{"", "first company"}
	findPersonType(student)
	findPersonType(engineer)
}
