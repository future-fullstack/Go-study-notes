package main

import "fmt"

// 简单来说，接口就是一份“行为契约”，它定义了一组方法签名
// 很抽象对吧，下面是关于接口的例子

// SingInterface 创建一个接口，SingInterface 是接口名字
// Sing 是它所包含的方法签名
type SingInterface interface {
	Sing()
}

// Chicken  Cat  定义两个结构体

type Chicken struct {
	Name string
}

type Cat struct {
	Name string
}

// 定义两个结构体的方法
// 很简单能从下述方法定义中发现这两个结构体有相同的方法 Sing
// 并且两个方法不接收参数且无返回值
// 那么这个方法(Sing)就被接口包含
// 能使用接口内方法签名的类型就是接口类型
// 具体应用如下

func (c Chicken) Sing() {
	fmt.Println(c.Name, "is singing")
}
func (c Cat) Sing() {
	fmt.Println(c.Name, "is singing")
}

// Define function
func sing(c SingInterface) {
	c.Sing()
}

func main() {
	ch := Chicken{Name: "ik"}
	ca := Cat{Name: "咪咪"}
	sing(ch)
	sing(ca)
}
