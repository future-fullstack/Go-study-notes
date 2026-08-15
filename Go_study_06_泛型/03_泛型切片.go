package main

import "fmt"

// main 演示泛型在切片类型上的应用
// main demonstrates using generics on slice types.
func main() {
	// MySlice 定义了一个约束为 T 的切片类型:元素只能是 int 或 string
	// MySlice defines a slice type parametrized by T: elements may only be int or string.
	type MySlice[T int | string] []T
	var mySlice = MySlice[int]{1, 2, 3} // 实例化为 int 切片	Instantiates an int slice.
	fmt.Println(mySlice[0] + 1)         // 因为元素确定是 int,可以直接做运算	Since elements are int, arithmetic works directly.
}
