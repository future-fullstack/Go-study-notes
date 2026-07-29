package main

// Code 自定义类型在网络请求场景下的应用
// 对于自定义类型优点的探讨详情放下下一节 "06_自定义类型与类型别名"
type Code int // 定义 Code 类型,底层类型是 int 类型	 Code defines a type with an underlying int type.

func (c Code) GetCodeMsg() (Code, string) {
	switch c {
	case SuccessCode:
		return c, "success"
	case ServiceErrCode:
		return c, "service error"
	case NetworkErrCode:
		return c, "network error"
	}
	return 0, ""
}

// 定义全局变量
const (
	SuccessCode    Code = 0    // 一切正常	Successful operation.
	ServiceErrCode Code = 1001 // 服务错误	Internal service error.
	NetworkErrCode Code = 1002 // 网络错误	Network communication failure.
)

func webServe(name string) (Code, string) {
	if name == "1" {
		return ServiceErrCode.GetCodeMsg()
	}
	if name == "2" {
		return NetworkErrCode.GetCodeMsg()
	}
	return SuccessCode.GetCodeMsg()
}

func main() {
}
