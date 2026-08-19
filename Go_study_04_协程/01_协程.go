package main

import (
	"fmt"
	"sync"
	"time"
)

// eat 模拟一个耗时的操作，使用 goroutine 并发执行
// eat simulates a time-consuming operation executed concurrently via goroutine.
func eat(name string, wait *sync.WaitGroup) {
	fmt.Println(name, "is eating")
	time.Sleep(1 * time.Second)
	wait.Done() // 通知 WaitGroup 该协程已完成	Notify WaitGroup that this goroutine is done.
}

func main() {
	var wait sync.WaitGroup
	wait.Add(3)                // 设置计数器，等待 3 个协程完成	Set counter to wait for 3 goroutines to complete.
	startime := time.Now()     // 记录开始时间	Record start time.
	go eat("Mike", &wait)      // 启动协程	Launch a goroutine.
	go eat("John", &wait)      // 启动协程	Launch a goroutine.
	go eat("Jack", &wait)      // 启动协程	Launch a goroutine.
	wait.Wait()                // 阻塞主协程直到所有协程完成	Block the main goroutine until all goroutines finish.
	fmt.Println(time.Since(startime)) // 打印总耗时	Print total elapsed time.
}
