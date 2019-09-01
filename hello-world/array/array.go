package main

import "fmt"

func main() {
	// 定义数据长度为10，元素为int类型的数字，未初始化，默认全部为0值
	var array [10]int
	fmt.Println(array)
	array[0] = 1
	array[4] = 3
	fmt.Println(array)
	fmt.Println(array[4])

	// 定义长度为4的string类型数据，前3个有初始化值
	fmt.Println("-------")
	arr := [4]string{"chen", "zuo", "li"}
	fmt.Println(arr)

	// 切片，相比数组来说更灵活，更常用，语法结构：array[start:end], 包含start，不包含end，即左闭右开方式.[)
	fmt.Println("-------")
	var chip = array[1:5]
	fmt.Println(chip)

	// 切片并不存储任何数据，它只是描述了底层数组的一个片段，更改切片的元素会修改底层数组的元素，与它共享底层数组的切片也会观测到这些变化。
	fmt.Println("-------")
	arr = [4]string{"i", "love", "you", "forever"}
	fmt.Println(arr)
	var subarr1 = arr[0:2]
	var subarr2 = arr[1:3]
	fmt.Println(subarr1, subarr2)
	subarr2[1] = "wangying"
	fmt.Println(arr, subarr1, subarr2)

	// 切片文法类似于没有长度的数组文法，它会在底层构建一个相同的数组文法，然后构建一个引用它的切片
	fmt.Println("--------")
	x := []int{1, 2, 3, 4, 5}
	y := []bool{true, false, true, false}
	z := []struct {
		i int
		j bool
	}{
		{1, true},
		{2, false},
		{3, true},
	}
	fmt.Println(x, y, z)
	fmt.Println(x[1:], y[:3], z[:]) // 上下边界写法

	// 切片的长度与容量，长度就是切片所包含的元素个数，容量是从切片的第一个元素开始到其底层数组的元素末尾的个数
	fmt.Println("---------")
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	printSlice(s)
	s = s[:0] // 缩减长度
	printSlice(s)
	s = s[:6] // 增加长度
	printSlice(s)
	s = s[2:] // 缩减容量
	printSlice(s)

	// 使用内置函数make创建切片函数
	ints := make([]int, 5)
	printSlice(ints)
}

func printSlice(a []int) {
	fmt.Printf("a's length is %s, a's cap is %s, a's value is %v\n", len(a), cap(a), a)
}
