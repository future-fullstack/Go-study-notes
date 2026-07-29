package main

import "fmt"

// defer 函数用来控制函数执行顺序：从上到下先把 非defer 命令执行一遍，再根据不同 defer 函数距离 return 的距离不同、按照距离由近到远依次执行 defer 函数
// The 'defer' statement controls execution order:non-deferred statements execute first from top to bottom.Then deferred function calls are executed in LIFO(Last-In-First-Out) order before the function return
// defer 函数只能使用该 defer 函数之前声明的变量
// A deferred function can only access variables declared before the 'defer' statement itself
// defer 函数主要用来资源清理
// The primary use case for 'defer' is resource cleanup(e.g.,closing files or releasing locks)
func main() {
	fmt.Println(1)
	//	defer fmt.Println(a)	这行会报错
	var a string = "hello"
	defer fmt.Println("defer 1", a)
	fmt.Println(2)
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3")
	return
}
