package main

import (
	"fmt"
)

// Hello 是一个最简单的函数:无参数、无返回值	Hello is the simplest function: no params, no return value.
func Hello() {
	fmt.Println("Hello")
}

// param1 演示带一个参数的函数	param1 demonstrates a function with one parameter.
func param1(id int) {
	fmt.Println(id)
}

// param2 演示带多个不同类型参数的函数	param2 demonstrates a function with multiple parameters of different types.
func param2(id int, username string) {
	fmt.Println(id, username)
}

// 参数类型相同时可以只给最后一个参数声明类型
// When parameters share the same type,the type can be declared only on the last one
func param3(id, age int) {
	fmt.Println(id, age)
}

// add1 通过传入切片来计算总和	 add1 computes the sum via a slice parameter.
// 法 1
// Method 1
func add1(numberList []int) {
	var sum int
	for _, number := range numberList {
		sum += number
	}
	fmt.Println(sum)
}

// add2 使用可变参数(...int),可以接收任意数量的 int,更灵活
// add2 uses variadic parameters (...int) to accept any number of ints, which is more flexible.
// 法2
// Method 2
func add2(numberList ...int) {
	var sum int
	for _, number := range numberList {
		sum += number
	}
	fmt.Println(sum)
}

// 返回值	Returns
// 无返回值	Returns nothing
func ri() {
	return
	// 没有返回值		Returns nothing
}

// 有 1 个返回值(注意声明方式)
// Sigle return value example(note the return value syntax)
func r2() int {
	return 8
}

// 有 2 个返回值
// Return two value
func r3() (int, bool) {
	return 1, false
}
// 提前声明要返回的值的类型
// Named return value
func r4() (ok bool, num int) {
	ok = false
	num = 8
	// 如果 return 后面跟具体值，会直接覆盖掉函数体内之前给 ok 和 num 赋的值
	// Explicit return values will override the named return variables
	return
}

// 小知识：Go 语言中函数内部不能再创建一般函数
// Note:Go doesn't support nested named function

// 匿名函数
// Anonymous function
// 想要在函数体内创建函数需要借助到特殊函数-----匿名函数
// To define a function inside a function,you must use an anonymous function(closure)
// 外部的匿名函数创建
// Anonymous function declared at the package level
var setAge = func(age int) {
	fmt.Println(age)
}

func main() {
	add2(1, 23)

	// 匿名函数
	// Anonymous function
	// 注意，这是在函数内部
	// Note:this is inside function
	// 下面的办法可以实现类似于函数嵌套的效果，但是这个和一般函数不同
	// This achieves a similar effect to nested functions,but it's not a regular (named) function declaration
	var setName = func(name string) {
		fmt.Println(name)
	}
	setName("a")
	setAge(23)

	// 高阶函数
	// High-order function
	// e.g.
	fmt.Println("请输入数字")
	fmt.Println("1.没吃饭")
	fmt.Println("2.吃饭了")
	var num int
	fmt.Scan(&num)

	// 将多个函数放入 map,用函数作为值,实现按需调用(函数也是类型)
	// Store functions as map values to dispatch by key (functions are first-class values).
	var m = map[int]func(){
		1: food,
		2: nofoof,
	}

	// 取出并调用对应的函数;ok 为 false 说明该键不存在
	// Fetch and call the matching function; ok is false if the key does not exist.
	fun, ok := m[num]
	if ok {
		fun()
	}
}

// food 根据用户选择提示去吃饭	 food suggests going to eat.
func food() {
	fmt.Println("快去吃饭")
}

// nofoof 根据用户选择提示去散步	 nofoof suggests taking a walk.
func nofoof() {
	fmt.Println("快去散步")
}
