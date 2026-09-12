package main

import (
	"fmt"
	"net/http"
)

// Index 处理根路径请求并返回简单的文本响应
// Index handles requests for the root path and returns a simple text response.
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)         // 打印客户端请求的路径	Print the path requested by the client.
	w.Write([]byte("Hello World!")) // 向客户端写回文本响应	Write a text response back to the client.
}

func main() {
	http.HandleFunc("/", Index)                                 // 把根路径交给 Index 处理	Route the root path to Index.
	fmt.Println("web server listen addr: http://127.0.0.1:801") // 提示服务监听地址	Print the address where the server is listening.
	http.ListenAndServe("127.0.0.1:801", nil)                   // 启动 HTTP 服务并持续监听请求	Start the HTTP server and keep listening for requests.
}
