package main

// 测试if else条件语句；

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// 这里开始就不能使用 v 了
	return lim
}

func main() {
	fmt.Println(pow(3, 2, 10),pow(3, 3, 20))  // go中会从内到外执行，执行完pow(3, 2, 10)后，再执行pow(3, 3, 20)，然后再执行fmt.Println(9, 20)
	fmt.Println(9, 20)
}
