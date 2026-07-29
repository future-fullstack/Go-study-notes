package main

import (
	"fmt"
	"time"
)

// select 用于监听多个 channel 的操作,类似于 switch,但专门用于 channel
// select listens for operations on multiple channels, similar to switch but dedicated to channels.
// 当某个 channel 就绪时,执行对应的 case;多个就绪则随机选择一个
// When a channel is ready, its corresponding case executes; if multiple are ready, one is picked at random.

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	// 启动协程向 ch1 发送数据
	// Launch a goroutine to send data into ch1.
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "one"
	}()

	// 启动协程向 ch2 发送数据
	// Launch a goroutine to send data into ch2.
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "two"
	}()

	// 使用 select 同时监听两个 channel
	// Use select to listen on both channels simultaneously.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("received from ch1:", msg1)
		case msg2 := <-ch2:
			fmt.Println("received from ch2:", msg2)
		}
	}

	// select 配合 default 实现非阻塞通信
	// select with default enables non-blocking communication.
	select {
	case msg := <-ch1:
		fmt.Println(msg)
	default:
		// 没有 channel 就绪时执行 default,不会阻塞
		// default executes when no channel is ready, avoiding a block.
		fmt.Println("no channel ready")
	}
}
