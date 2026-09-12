package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("http://127.0.0.1:801") // 向本地 HTTP 服务发起 GET 请求	Send a GET request to the local HTTP server.
	if err != nil {
		fmt.Println(err) // 打印请求错误	Print the request error.
		return           // 请求失败时提前结束	Exit early when the request fails.
	}
	byteData, _ := io.ReadAll(response.Body) // 读取响应体中的全部字节	Read all bytes from the response body.
	fmt.Println(string(byteData))            // 把响应字节转换成字符串并打印	Convert the response bytes to a string and print them.
}
