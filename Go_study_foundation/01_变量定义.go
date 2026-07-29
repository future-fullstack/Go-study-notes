package main

import "fmt"

// 函数外定义全局变量不能用短声明符号(:=)
//Short variable declaration(:=) cannot be used outside a function to declare global variables

// 全局变量不用不会报错，也不会消失，但是函数内部不用已经声明的变量就会报错
// Unused global variables won't cause compilation errors,nor will they disappear
// However,declared variables inside function MUST be used,otherwise it will trigger compilation errors

func main() {
	//方法一：先声明、再赋值
	//Method 1:Dlare first,then assign
	var name1 string
	name1 = "WHM"
	fmt.Println(name1)

	//方法二：声明并赋值
	//Method 2:Delare and assign simultaneously
	var name2 string = "WHM"
	fmt.Println(name2)

	//方法三：省略类型直接赋值
	//Method 3:type inference （short declaration）
	var name3 = "WHM"
	fmt.Println(name3)

	//方法四：声明并赋值（短声明符号）	Method 4:declare and assign simultaneously (short declaration operator\symbol)
	name4 := "WHM"
	fmt.Println(name4)

	//定义多个变量
	//declare variables simultaneously(like tuple unpacked in math)
	b1, b2 := 10, 20
	fmt.Println(b1, b2)

}

const Version = "1.0.1" //常量定义方式（定义之后无法修改）	consent declaration (immutable after definition)

//命名规范：想要实现跨包访问变量首字母要大写
//Naming convention :To be exported (accessible across package),the variable name must start with a capital letter
