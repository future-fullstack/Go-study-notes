package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	// 属性,类型,结构体tag(标签)
	// Field, type, and struct tag.
	// 结构体标签的作用是：当转换成 JSON 时规定属性的键名
	// Struct tags specify the key names when marshalling to JSON.
	Name     string `json:"name"`
	Age      int    `json:"age,omitempty"` // omitempty: 字段为空值时则在 JSON 中忽略	Exclude field if it holds its zero value.
	Password string `json:"-"`             // "-": 转换为 JSON 时忽略该字段	Exclude field entirely from JSON output.
}

func main() {
	user := User{Name: "WHM", Age: 18, Password: "123456"} // 初始化结构体实例	Initialize the struct instance.
	fmt.Println(user)                                       // 直接打印结构体	Print the struct directly.
	byteDtae, _ := json.Marshal(user)                       // 将结构体序列化为 JSON	Marshal the struct to JSON.
	fmt.Println(string(byteDtae))                           // 输出 JSON:注意 password 被忽略	Name:注意 JSON 中不含 password,因为有 json:"-" 标签。
	fmt.Printf(string(byteDtae))
}
