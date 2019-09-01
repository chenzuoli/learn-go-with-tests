package main

import "fmt"

// 自定义函数类型
type add func(a int, b int) int

// 函数作为参数传递
func sample(a func(a int, b int) int) {
	fmt.Println(a(60, 7))
}

// 函数作为返回值
func return_func() func(a int, b int) int {
	f := func(a int, b int) int {
		return a + b
	}
	return f
}

func main() {
	var a add = func(a int, b int) int {
		return a + b
	}
	s := a(1, 2)
	fmt.Println(s)
	fmt.Println("------")
	f := func(a int, b int) int { // set the 'f' function as the parameter of the function 'sample'
		return a + b
	}
	sample(f)
	fmt.Println("------")
	t := return_func()
	fmt.Println(t(49, 1)) // function as the return type
}
