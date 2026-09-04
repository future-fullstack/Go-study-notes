package main

import (
	"fmt"
	"reflect"
)

// Student 用于演示:遍历结构体字段时,可以同时拿到字段名、类型、标签等信息
// Student is used to demonstrate: when iterating struct fields we can grab the field name, type, struct tag, etc.
type Student struct {
	Name  string `json:"name"`
	Age   int
	IsMan bool
}

// ParseJson 用反射遍历结构体的所有字段,并打印字段的元信息与值
// ParseJson iterates over all fields of a struct via reflection, printing each field's metadata and value.
func ParseJson(obj any) {
	v := reflect.ValueOf(obj)           // 值容器,用 NumField / Field(i) 访问字段	Value wrapper; access fields via NumField / Field(i).
	t := reflect.TypeOf(obj)            // 类型信息,用 Field(i) 拿到字段的描述	Type info; Field(i) gives each field's description.
	for i := 0; i < v.NumField(); i++ { // NumField() 得到字段数量	NumField() returns how many fields there are.
		fmt.Println(t.Field(i).Name, t.Field(i).Type, t.Field(i).Tag, t.Field(i).Tag.Get("json")) // 字段名/类型/标签原文/按 json 键取值	Field name/type/raw tag/tag value looked up by "json".
		fmt.Println(v.Field(i))                                                                   // 字段对应的实际值	The field's actual value.
	}
}

func main() {
	s := Student{
		Name:  "WHM",
		Age:   20,
		IsMan: true,
	}
	ParseJson(s) // 注意:这里传的是值而非指针,遍历字段足够用	Note: passing a value (not pointer) is enough for iterating fields.
}
