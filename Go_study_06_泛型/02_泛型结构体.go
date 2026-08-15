package main

import (
	"encoding/json"
	"fmt"
)

// Response 是一个泛型结构体:T 可以是任意类型(any),由 Data 字段承载
// 常用于封装统一的 API 返回格式,而具体的数据类型由调用方指定
// Response is a generic struct: T can be any type, carried by the Data field.
// It is commonly used to wrap a unified API response format, with the concrete type specified by the caller.
type Response[T any] struct {
	Code int    `json:"code"` // 状态码	Status code.
	Msg  string `json:"msg"`  // 提示信息	Message.
	Data T      `json:"data"` // 具体数据(类型由调用方决定)	Actual payload (type decided by the caller).
}

// main 将同一个泛型结构体实例化为不同数据类型
// main instantiates the same generic struct with different data types.
func main() {
	type User struct { // 定义用户类型	Define the User type.
		Name string `json:"name"`
	}

	type UserInfo struct { // 定义用户信息类型	Define the UserInfo type.
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	//user := Response{
	//	Code: 0,
	//	Msg:  "success",
	//	Data: User{
	//		Name: "Mike",
	//	},
	//}
	//byteData, _ := json.Marshal(user)
	//fmt.Println(string(byteData))
	//
	//userInfo := Response{
	//	Code: 0,
	//	Msg:  "success",
	//	Data: UserInfo{
	//		Name: "Mike",
	//		Age:  20,
	//	},
	//}
	//byteData, _ = json.Marshal(userInfo)
	//fmt.Println(string(byteData))

	var userResponse Response[User]
	json.Unmarshal([]byte(`{"code":0,"msg":"success","data":{"name":"Mike","age":20}}`), &userResponse)
	fmt.Println(userResponse)

	var userInfoResponse Response[UserInfo]
	json.Unmarshal([]byte(`{"code":0,"msg":"success","data":{"name":"Mike","age":20}}`), &userInfoResponse)
	fmt.Println(userInfoResponse)
	// {"code":0,"msg":"success","data":{"name":"Mike","age":20}}
}
