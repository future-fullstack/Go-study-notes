package main

import (
	"fmt"
	"os"
)

// init 会在 main 之前自动执行,这里尝试读取一个不存在的文件
// init runs automatically before main; here it tries to read a file that does not exist.
func init() {
	_, err := os.ReadFile("1234")
	if err != nil {
		// 也可以使用 log.Fatalln("错误了") 打印日志后直接退出(这里暂时注释掉)
		// Alternatively, log.Fatalln("错误了") logs the message and exits directly (commented out for now).
		//log.Fatalln("错误了")
		panic("错误了") // panic 会打印堆栈信息并立即终止程序,程序被中断	panic prints a stack trace and terminates the program immediately, interrupting it.
	}
}

func main() {
	// 因为 init 中发生了 panic,main 的这行代码永远不会执行,程序已被中断
	// Because panic occurred in init, this line in main never runs; the program is already interrupted.
	fmt.Println("main")
}
