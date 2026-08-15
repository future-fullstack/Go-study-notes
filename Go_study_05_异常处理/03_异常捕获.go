package main

import (
	"fmt"
	"runtime/debug"
)

// read 使用 defer + recover 捕获运行时 panic(数组越界),使程序崩溃后仍能恢复执行
// read uses defer + recover to catch a runtime panic (index out of range) so the program can recover.
func read() {
	defer func() {
		err := recover() // recover 捕获 panic 的值	recover captures the value of the panic.
		if err != nil {  // 若非空,说明确实发生了 panic	If non-nil, a panic did occur.
			fmt.Println(err)                   // 打印 panic 信息	Print the panic message.
			fmt.Println(string(debug.Stack())) // 打印错误信息的堆栈	Print the stack trace of the error.
		}
	}()
	var list = []int{1, 2}
	fmt.Println(list[2]) // 访问越界,触发 panic	Access out of range, triggering a panic.
}

func main() {
	read() // 即使 read 内 panic,也不会导致整个程序中断	Even if read panics, the whole program is not interrupted.

	// 正常逻辑
	fmt.Println("正常逻辑")
}
