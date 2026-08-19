package main

import "fmt"

// Student 结构体的定义(类似于 python 的类)
// Student defines this structure,analogous to a class in python
// 但是在 Go 中定义结构体时是是仅仅定义属性的，相关函数(方法)是外部定义
// However,in Go,struct are defined solely for attributes;associated functions(methods) are defined externally

type Class struct {
	Name string
}
type Student struct {
	Class
	Name string
}

// Study 结构体方法
func (s Student) Study() {
	fmt.Printf("%s is studying!And he is in %s\n", s.Name, s.Class.Name)
}

func main() {
	s1 := Student{ // 使用键值对方式初始化结构体	Initialize the struct with key-value pairs.
		Class: Class{Name: "classroom1"},
		Name:  "WHM"}
	fmt.Printf("%T\n", s1) // 打印结构体类型	Print the struct's type.
	s1.Study()             // 调用结构体方法	Call the struct's method.

}
