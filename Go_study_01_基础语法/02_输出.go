package main

import "fmt"

func main() {
	//fmt.Println自带换行
	//fmt.Println automatically append a newline
	fmt.Println("Hello", "nihao")
	fmt.Println("你好")

	//格式化输出默认没换行符,想加换行符自己加(\n)
	//Formatted does have a newline by default;add \n yourself if you need
	//字符串
	//String
	fmt.Printf("%s is so handsome \n", "Whm")

	//整数 d
	//Integer placeholder %d
	fmt.Printf("%d\n", 5)

	//小数 f .2表示保留两位小数
	//Float placeholder %f;.2 specifies two decimal place
	fmt.Printf("%f.2\n", 5.8)

	//查询未知类型 T
	//%T reveals the date type of the variable
	fmt.Printf("%T\n", 5.8)

	//任意类型v
	//%v for any type
	fmt.Printf("%v %v \n", "你好", "我是谁")

	//打印空字符串的两种方式v 和 #v
	//%v vs %#v for printing an empty string
	fmt.Printf("%v\n", "")
	fmt.Printf("%#v\n", "")

	//变量赋值
	//variable assignment
	var f = fmt.Sprintf("%v\n", "我是f")
	fmt.Println(f)

}
