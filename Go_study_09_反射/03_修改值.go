package main

import (
	"fmt"
	"reflect"
)

// setValue 通过反射修改变量的值:obj 必须传指针,Elem() 解引用后才能真正写入
// setValue modifies a variable via reflection: obj must be a pointer, and Elem() dereferences it so the value can actually be written.
func setValue(obj any, value any) {
	v1 := reflect.ValueOf(obj)   // obj 是指针的 Value	The Value of obj (a pointer).
	v2 := reflect.ValueOf(value) // 待写入的新值的 Value	The Value of the new value to write.
	// Elem() 取指针指向的实体;若其 Kind 与新值不一致则不修改直接返回
	// Elem() follows the pointer; if its Kind differs from the new value, bail out without modifying.
	if v1.Elem().Kind() != v2.Kind() {
		return
	}
	switch v1.Elem().Kind() {
	case reflect.String:
		v1.Elem().SetString(value.(string)) // 写入字符串	Write the string.
	case reflect.Int:
		v1.Elem().SetInt(v2.Int()) // 写入整数	Write the integer.
	}

}

func main() {
	var name = "FSE"
	var age = 19
	setValue(&name, "fse") // 传入指针才能被修改	Pass a pointer so it can be modified.
	setValue(&age, 20)
	fmt.Println(name, age) // 已被修改成小写与 20	Both are now changed to lowercase "fse" and 20.
}
