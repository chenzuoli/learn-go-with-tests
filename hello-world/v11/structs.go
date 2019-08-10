package main

import "fmt"

type abc struct { // struct数据结构
	x int
	y int
}

func main() {
	fmt.Println(abc{1, 2}) // 大括号赋值，2个字段均须赋初始值
	var v = abc{3, 4}      // 必须给struct赋给一个变量才能进行访问，不能直接访问  abc.x   abc.y
	fmt.Println(v.x)
	fmt.Println("---------")
	var p = &v
	fmt.Println(p)
	fmt.Println(p.x)    // 结构体指针可以直接访问x  y的值，而不需要写成(*p).x
	fmt.Println((*p).x) // 这样也行，更简便的方式就是p.x
}
