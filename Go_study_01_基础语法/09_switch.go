package main

import "fmt"

func main() {
	fmt.Println("请输入年龄：")
	var age int
	fmt.Scan(&age)
	// 用法1
	// Method 1
	// 在 switch 语句中，一旦满足其中一个条件直接停止继续向下运行
	// In a switch statement,execution stops immediately once a case matched
	switch {
	case age <= 0:
		fmt.Println("未出生")
	case age < 18:
		fmt.Println("未成年")
		// fallthrough:加上这个命令之后程序在满足该条件下仍然会继续向下运行
		// fallthrough:Adding fallthrough forces program to continue to the next case even after a case is matched
	case age <= 35:
		fmt.Println("青年")
	default:
		fmt.Println("中年")

	}

	fmt.Println("请输入星期(数字)：")
	var day int
	fmt.Scan(&day)
	// 用法二		注意：case 在这种使用方式之下 case 后面只能跟具体的值
	// Method 2	Note:In this usage,case must be followed by a specific value
	switch day {
	case 1:
		fmt.Println("今天是星期一")
	case 2:
		fmt.Println("今天是星期二")
	case 3:
		fmt.Println("今天是星期三")
	case 4:
		fmt.Println("今天是星期四")
	case 5:
		fmt.Println("今天是星期五")
	case 6, 7:
		fmt.Println("今天是周末")
	}

}
