package main

import (
	"fmt"
	"time"
)

// done 是一个无缓冲信道,用于在主协程和子协程之间传递"完成"信号
// done is an unbuffered channel used to signal completion between the main and the child goroutine.
var done = make(chan struct{})

// event 模拟一个耗时操作,执行完毕后关闭 done 信道通知外界
// event simulates a time-consuming task, closing the done channel to notify the outside when finished.
func event() {
	fmt.Println("开始执行")
	time.Sleep(3 * time.Second) // 模拟耗时 3 秒	Simulate a 3-second task.
	fmt.Println("执行结束")
	close(done) // 关闭信道表示事件已完成	Closing the channel signals the event is done.
}

func main() {
	go event() // 启动子协程执行事件	Launch the child goroutine to run the event.

	// 关于 select 的细节 :select 不是循环，只要进入一个分支就相当于主线程进入岔路
	// Note about select: it is not a loop; once a branch is entered, the main thread goes down that path.
	select {
	case <-done: // 若 done 先关闭,说明事件在超时前完成	If done closes first, the event finished in time.
		fmt.Println("协程执行完毕")
	case <-time.After(2 * time.Second): // 若 2 秒定时器先触发,说明事件超时	If the 2-second timer fires first, the event timed out.
		fmt.Println("超时")
		return // 超时后直接结束程序	Return to end the program on timeout.
	}
}
