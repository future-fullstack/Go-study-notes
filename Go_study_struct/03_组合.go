package main

import "fmt"

type Animal struct {
	Name string
}

// Cat 按照下述方式可以实现结构体的组合
type Cat struct {
	Animal
}

// 结构体组合后的赋值方法
func main() {
	animal := Animal{Name: "KK"}
	cat := Cat{Animal: animal}
	fmt.Println(cat.Animal.Name)
}
