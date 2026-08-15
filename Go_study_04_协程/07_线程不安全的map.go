package main

import "fmt"

func main() {
	// 创建一个普通的 map,必须初始化后才能向 map 中写入值
	// Create a normal map; it must be initialized before values can be written to it.
	var m = make(map[int]string)

	// 协程 1:无限循环向 map 写入数据
	// Goroutine 1: writes data to the map in an infinite loop.
	go func() {
		for {
			m[1] = "张三"
		}
	}()

	// 协程 2:无限循环从 map 读取数据
	// Goroutine 2: reads data from the map in an infinite loop.
	// 普通 map 不支持并发读写,两个协程同时操作会直接报 fatal error:concurrent map read and map write
	// A plain map does not support concurrent access; simultaneous reads and writes trigger
	// a fatal error: "concurrent map read and map write".
	go func() {
		for {
			fmt.Println(m[1])
		}
	}()

	select {} // 空 select 会永久阻塞主协程,让上面的协程一直运行	An empty select blocks the main goroutine forever, keeping the goroutines running.
}
