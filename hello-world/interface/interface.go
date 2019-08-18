package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X float64
	Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f)

	v := Vertex{3, 4}

	a = f  // a MyFloat实现了Abser interface
	a = &v // a *Vertex实现了Abser interface

	fmt.Println(a)
}
