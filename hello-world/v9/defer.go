package main

import "fmt"
import "strconv"

func main() {
	defer fmt.Println("world") // defer语句会将函数推迟到外层函数执行完成后执行
	fmt.Println("hello\n------------")
	test_defer(1)
	fmt.Println("-------------")
	defer_stack()
}


func test_defer(param int) int {
	defer fmt.Println(param)
	param = param + 1
	fmt.Println("param's value is " + strconv.Itoa(param))
	return param
}


// defer栈，推迟的函数会被压入一个栈中，当外层函数返回时，被推迟的函数会按照后进先出的原则进行执行
func defer_stack(){
	fmt.Println("counting.")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done.")
}

