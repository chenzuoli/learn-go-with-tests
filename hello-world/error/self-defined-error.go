package main

import (
	"errors"
	"fmt"
	"math"
)

// calculate the circle area

// catch the error and define the error, use the errors module.
func area(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("the given radius is less than zero.")
		//return 0, fmt.Errorf("the given radius %s is less than zero.", radius)
	}
	return math.Pi * radius * radius, nil
}

func main() {
	radius := -10.0
	area, err := area(radius)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("the circle's area is ", area)
	}
}
