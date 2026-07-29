package main

import (
	"fmt"
	"time"
)

// 延时计算求和
// Sum calculation with delayed execution
// Go 语言中的匿名函数使用场景常常对应 python 的闭包
// Anonymous functions in Go are often used in scenarios similar to python closures
func awaitAdd(awaitSecond int) func(...int) int {
	var add = func(numList ...int) int {
		time.Sleep(time.Duration(awaitSecond) * time.Second)
		var sum int
		for _, i := range numList {
			sum += i
		}
		return sum
	}
	return add
}

// 值传递和引用传递
// Value vs. Pointer (Reference) Passing
// 值传递：传递的是原变量的一个副本（拷贝），不会影响原变量
// Pass by value: Passes a copy of the original variable. Modifications do not affect the original.
func Copy(name string) {
	fmt.Println("in funcC %p\n", &name)
}

// 引用传递：
// Pass by reference (using pointers)
func Set(name *string) {
	fmt.Println("in funcS %p\n", name)
	*name = " 我现在变了 "
}
func main() {
	//add2 := awaitAdd(2)
	//t1 := time.Now()
	//fmt.Println(add2(1, 2, 3))
	//t2 := time.Now()
	//fmt.Println(t2.Sub(t1))

	// 结合上述函数定义我们知道：参数定义的内存地址 与 传递参数到函数内部之后"参数"的内存地址不同，所以传参时是将参数拷贝了一份传递进去
	// As inferred from the function definitions above, the memory address of the declared variable differs from the address of the parameter inside the function. This demonstrates that arguments are passed as copies.
	var name string = " 你好呀，我是最初的值 "
	fmt.Printf("variable %p\n", &name)
	fmt.Printf("func %p\n", Copy)
	Copy(name)
	Set(&name)
	fmt.Printf(name)

}
