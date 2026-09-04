package main

import (
	"fmt"
	"reflect"
	"strings"
)

// User 带结构体标签 big:标了 big:"-" 的字段会被反射单独处理(转大写)
// User carries struct tags named "big": fields tagged big:"-" are handled specially (upper-cased) during reflection.
type User struct {
	Name1 string `big:"-"`
	Name2 string
}

// SetStruct 用反射遍历结构体字段,把带 big 标签且值非空的字段转成大写
// SetStruct iterates the struct's fields via reflection and upper-cases any field that carries a non-empty "big" tag.
func SetStruct(obj any) {
	v := reflect.ValueOf(obj).Elem()    // 指针解引用后的实体 Value,才能 SetString	The dereferenced Value (so SetString is possible).
	t := reflect.TypeOf(obj).Elem()     // 对应类型,用于读取字段标签	The matching type, used to read field tags.
	for i := 0; i < v.NumField(); i++ { // 遍历所有字段	Iterate over all fields.
		value := v.Field(i)              // 当前字段的可写 Value	The writable Value of the current field.
		fmt.Println(value)               // 打印字段原值	Print the field's original value.
		big := t.Field(i).Tag.Get("big") // 读取该字段 big 标签的值	Read the "big" tag value of this field.
		if big == "" {                   // 没有标 big 标签就跳过(不处理)	Skip fields without the "big" tag.
			continue
		}
		value.SetString(strings.ToUpper(value.String())) // 把字段值改成大写	Replace the field value with its upper-case version.
	}

}

func main() {
	s := User{
		Name1: "name1",
		Name2: "name2",
	}
	SetStruct(&s) // 传指针以便修改	Name1 带 big:"-" 会被改成大写,Name2 保持原样.
	fmt.Println(s)
}
