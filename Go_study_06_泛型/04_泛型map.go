package main

import "fmt"

// main 演示泛型在 map 类型上的应用
// main demonstrates using generics on map types.
func main() {
	// map 的 key 只能是基本数据类型
	// The key of a map must be a basic (comparable) type.
	// MyMap 使用两个类型参数:T 约束 key,必须为 int 或 string;K 约束 value,可为任意类型
	// MyMap uses two type parameters: T constrains the key to int or string, K constrains the value to any type.
	type MyMap[T int | string, K any] map[T]K
	var myMap = MyMap[string, int]{ // 实例化 key 为 string、value 为 int	Instantiate with string keys and int values.
		"Age": 12,
	}
	fmt.Println(myMap)
}
