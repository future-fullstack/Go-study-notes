package main

import (
	"fmt"
	"reflect"
)

// getType 通过反射读取任意对象的类型,再用 switch 判断其底层 Kind(种类)
// getType reads the type of any object via reflection, then uses a switch to inspect its underlying Kind.
func getType(obj any) {
	t := reflect.TypeOf(obj) // 获取对象的类型信息	Obtain the object's type info.
	switch t.Kind() {        // 判断类型属于哪一种 Kind	Determine which Kind the type belongs to.
	case reflect.Struct:
		fmt.Println("Struct")
	case reflect.String:
		fmt.Println("String")
	case reflect.Int:
		fmt.Println("Int")
	}

}

func main() {
	getType("sdasdasdad") // 传入字符串,底层 Kind 是 String	Pass a string, whose Kind is String.
	getType(213)          // 传入整数,底层 Kind 是 Int	Pass an integer, whose Kind is Int.
}
