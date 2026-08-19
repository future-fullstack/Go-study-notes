package main

import "fmt"

// Student  定义结构体
type Student struct {
	Name   string
	Age    int
	Gender string
}

// Update 定义 Update 方法
func (stu *Student) Update(name string) {
	fmt.Printf("%p\n", stu)
	stu.Name = name
}

// "查询位置"这个动作查询的是变量所在内存位置
// The act of "querying the location" refers to querying the memory address where the variable resides.
// 可以将内存看成一个小房间，定义变量就是给房间取名字或者说贴门牌号
// Memory can be viewed as a small room. Defining a variable is equivalent to naming the room or attaching a nameplate to it.
// 操作 就是走进这个门牌号对应的房间，改变房间内部的东西（比如换掉家具、改名字）
// Performing an operation corresponds to entering the room designated by the nameplate and modifying its internal contents (e.g., replacing furniture or changing its name).
// 通过下述代码可以看出无论是这个房间的什么细节，只要是查询内存地址，那么就都是一样的
// The following code demonstrates that regardless of the specific detail of the room being queried, the queried memory addresses are identical.

func main() {
	stu := Student{Name: "Mike", Age: 17, Gender: "male"}
	fmt.Printf("%p\n", &stu)
	fmt.Printf("%p\n", &stu.Age)
	stu.Update("John")
	fmt.Println(stu.Name)
}
