package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:81") // 主动连接服务端的 81 端口,成功则返回 conn,如拨号	Dial the server's port 81; on success return conn, like dialing a number.
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		var byteData = make([]byte, 1024) // 申请 1024 字节缓冲,用来装本次读到的数据	Allocate a 1024-byte buffer to hold the data read this time.
		n, err := conn.Read(byteData)     // 阻塞读取服务端数据,返回实际读到的字节数 n	Block reading server data, returning the bytes actually read as n.
		if err == io.EOF {
			break // EOF 表示对方已 Close 连接无更多数据,跳出循环收尾	EOF means the peer closed the connection with no more data; break the loop.
		}
		fmt.Println(string(byteData[:n]))
	}
}
