package main

import "fmt"

func main() {
	var array [10]int // 定义数据长度为10，元素为int类型的数字，未初始化，默认全部为0值
	fmt.Println(array)
	array[0] = 1
	array[4] = 3
	fmt.Println(array)
	fmt.Println(array[4])
	fmt.Println("-------")
	arr := [4]string{"chen", "zuo", "li"} // 定义长度为4的string类型数据，前3个有初始化值
	fmt.Println(arr)
	fmt.Println("-------")
	var chip = array[1:5] // 切片，相比数组来说更灵活，更常用，语法结构：array[start:end], 包含start，不包含end，即左闭右开方式。∏
	fmt.Println(chip)
}
