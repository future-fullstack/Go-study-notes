package main

import (
	"fmt"
	"sync"
)

func main() {
	// sync.Map 是标准库提供的并发安全 map,适合并发读写场景
	// sync.Map is a concurrency-safe map from the standard library, suited for concurrent reads and writes.
	var m = sync.Map{}

	// 协程 1:无限循环向 sync.Map 写入数据
	// Goroutine 1: writes data to the sync.Map in an infinite loop.
	go func() {
		for {
			m.Store(1, "张三") // Store 方法用于写入键值对	The Store method writes a key-value pair.
		}
	}()

	// 协程 2:无限循环从 sync.Map 读取数据
	// Goroutine 2: reads data from the sync.Map in an infinite loop.
	go func() {
		for {
			val, ok := m.Load(1) // Load 方法用于读取键对应的值,返回值和该键是否存在	The Load method reads the value for a key, returning the value and whether the key exists.
			fmt.Println(val, ok)
		}
	}()

	select {} // 空 select 会永久阻塞主协程,让上面的协程一直运行	An empty select blocks the main goroutine forever, keeping the goroutines running.
}
