package main

import "fmt"
import "runtime"
import "time"

const spanish = "Spanish"
const french = "French"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "

// Hello returns a personalised greeting in a given language
func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	prefix := englishHelloPrefix

	switch language { // switch 后可以接变量
	case french: // 只运行选定的case，而非case后所有的case
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	}

	return prefix + name
}

func run_os() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os { // 初始化变量os
	case "darwin":
		fmt.Println("OS X.")
		fallthrough // 跳过执行下一个case
	case "linux1":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

func switch_with_no_conditions() { // 没有条件的switch语句相当于switch true, 那么case语句中也必须使用bool类型的判断条件
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}

	switch {
	case 1 < 1:
		fmt.Println("1")
	case 2 > 1:
		fmt.Println("2")
	}

}

func main() {
	fmt.Println(Hello("world", "Spanish"))
	fmt.Println("-----------------")
	run_os()
	fmt.Println("-----------------")
	switch_with_no_conditions()
}
