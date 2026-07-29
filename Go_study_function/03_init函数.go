package main

import "fmt"

// init 函数最先被执行
// The init() function executes before main().
// 不同 init 函数之间谁在前谁先被执行
// If multiple init() functions exist,they execute in order they are defined(lexical order)
// 该函数不接收参数，没有返回值
// It takes no arguments and return no values
// 该函数在实际应用之中主要用于初始化
// In practice,it is primarily used for package-level initialization

// 注意下列函数运行顺序
// Note the execution order of follow functions
func init() {
	fmt.Println("Num1")
}

func init() {
	fmt.Println("Num2")
}

func init() {
	fmt.Println("Num3")
}

func main() {
	fmt.Println("Num4")
}
