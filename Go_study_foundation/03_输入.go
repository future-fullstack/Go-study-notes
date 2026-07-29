package main

import (
	"fmt"
)

//fmt.Scan(&) 在给内存划定区域的同时给用户创造了可输入的接口
//fmt.Scan(&) allocates memory space and creates an input interface for user

func main() {
	fmt.Print("请输入你的名字：")
	var name string
	fmt.Scan(&name)
	fmt.Println(name)
	fmt.Print("请输入你的年龄：")
	var age int
	fmt.Scan(&age)
	fmt.Println(age)

	//n,err可以接收fmt.Scan的返回值，其中n代表接收数据数量，err错误类型
	//n,err can receive the return values for fmt.Scan,where n represents the number of items reads,and err represents the error type
	var num int
	n, err := fmt.Scan(&num)
	fmt.Println(n, err, num)

}
