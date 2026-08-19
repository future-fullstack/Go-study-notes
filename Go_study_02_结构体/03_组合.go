package main

import "fmt"

// Animal 是一个基础结构体,只有 Name 字段
// Animal is a base struct with only a Name field.
type Animal struct {
	Name string
}

// Cat 按照下述方式可以实现结构体的组合:把 Animal 作为匿名字段嵌入,使得 Cat 自动拥有 Animal 的字段和方法
// Structural embedding (composition): by embedding Animal as an anonymous field,
// Cat automatically inherits Animal's fields and methods.
type Cat struct {
	Animal
}

// 结构体组合后的赋值方法
// main shows how to assign values after embedding.
func main() {
	animal := Animal{Name: "KK"}          // 分别创建 Animal 实例	Create an Animal instance.
	cat := Cat{Animal: animal}            // 把 animal 作为嵌入字段赋值给 Cat	Assign animal to Cat's embedded field.
	fmt.Println(cat.Animal.Name)          // 通过嵌入字段访问 Name	Access Name through the embedded field.
	fmt.Println(cat.Name)                 // 也可直接访问被提升的字段	Name is promoted, so it can be accessed directly.
}
