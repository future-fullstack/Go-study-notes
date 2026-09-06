package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:81") // 把 "ip:端口" 解析成 TCP 地址对象,首返回值基本成功故丢弃	Resolve the "ip:port" string into a TCP address object; the first result rarely fails so it is dropped.
	listen, err := net.ListenTCP("tcp", addr)            // 在本机 81 端口监听,占住端口等待客户端,如"开机等电话"	Listen on port 81, holding it open for clients, like "waiting by the phone".
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("TCP server bind addr : ", addr.String())
	for {
		conn, err := listen.Accept() // 阻塞等待一个客户端连入,返回一条已建立的连接,每次循环处理一个	Block until a client connects, returning one established conn; each loop serves one client.
		if err != nil {
			break // 出错则跳出循环结束(演示用无重试)	On error, break the loop and exit (demo, no retry).
		}
		fmt.Println(conn.RemoteAddr())
		conn.Write([]byte("hello world")) // 向该客户端发送字节数据	Send the byte data to this client.
		time.Sleep(2 * time.Second)
		conn.Close() // 关闭连接,如同挂断电话	Close the connection, like hanging up the phone.
	}
}
