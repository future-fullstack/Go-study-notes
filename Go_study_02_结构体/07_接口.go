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

// sing 接收一个接口类型参数:任何实现了 Sing() 方法的类型都可以传入,这里实现了多态
// sing takes an interface-typed parameter: any type that implements Sing() can be passed in, achieving polymorphism.
func sing(c SingInterface) {
	c.Sing()
}

func main() {
	ch := Chicken{Name: "ik"}  // 创建 Chicken 实例	Create a Chicken instance.
	ca := Cat{Name: "咪咪"}    // 创建 Cat 实例	Create a Cat instance.
	sing(ch)                   // 因为 Chicken 实现了 Sing(),可以作为接口传入	Chicken implements Sing(), so it satisfies the interface.
	sing(ca)                   // Cat 同样满足接口	Cat also satisfies the interface.
}
