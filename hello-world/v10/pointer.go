package main

import "fmt"
import "strconv"

// pointer指针测试
func main() {
	var i, j, k = 32, 34, 26
	fmt.Println("i's value is " + strconv.Itoa(i) + " and j's value is " + strconv.Itoa(j))
	var p = &i // 定义指针
	fmt.Println(p)
	fmt.Println(*p) // 获取指针指向的值
	fmt.Println("---------")
	var q *int
	q = &k
	*q = j
	fmt.Println(q)
	fmt.Println(*q)
	fmt.Println(k)
}
