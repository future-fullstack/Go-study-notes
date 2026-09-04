package main

import (
	"fmt"
	"reflect"
)

// getValue 与 getType 类似,但这里拿到的是"值容器" reflect.Value,可以直接取出底层数据
// getValue is similar to getType, but here we get a reflect.Value wrapper, from which the underlying data can be pulled out directly.
func getValue(obj any) {
	v := reflect.ValueOf(obj) // 获取对象的值容器	Obtain the object's reflect.Value wrapper.
	switch v.Kind() {         // 判断值的 Kind	Determine the Kind of the value.
	case reflect.Struct:
		fmt.Println("Struct")
	case reflect.String:
		fmt.Println("String", v.String()) // v.String() 取出真正的字符串	Extract the actual string via v.String().
	case reflect.Int:
		fmt.Println("Int", v.Int()) // v.Int() 取出真正的整数	Extract the actual integer via v.Int().
	}

}

func main() {
	getValue("Hello world") // 打印 String + 字符串内容	Print "String" plus the string content.
	getValue(123)           // 打印 Int + 整数内容	Print "Int" plus the integer content.
}
