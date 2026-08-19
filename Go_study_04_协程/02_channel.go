package main

import (
	"fmt"
	"sync"
	"time"
)

// moneyChan 定义信道,用于协程间通信
// moneyChan defines a channel for communication between goroutines.
// 格式: name(信道名) chan(关键字) type(传入信道的数据类型)
// Syntax: name(chan name) chan(keyword) type(data type passed through the channel)
var moneyChan = make(chan int)

// pay 模拟消费行为,并将消费金额发送到信道中
// pay simulates a shopping activity and sends the spent amount into the channel.
func pay(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s 开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 购物结束\n", name)
	moneyChan <- money // 将金额写入信道	Send the amount into the channel.
	wait.Done()
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)

	go pay("Mike", 10, &wg)
	go pay("John", 2, &wg)
	go pay("Amn", 17, &wg)

	// 在另一个协程中等待所有支付完成,然后关闭信道
	// Wait for all payments to complete in a separate goroutine, then close the channel.
	go func() {
		wg.Wait()
		close(moneyChan) // 关闭信道,通知接收方不再有数据	Close the channel to signal receivers that no more data will be sent.
	}()

	// 关闭 1: 手动检查信道是否已关闭
	// Method 1: Manually check whether the channel is closed.
	//for {
	//	money, open := <-moneyChan
	//	fmt.Println(money, open)
	//	if !open {
	//		break
	//	}
	//}

	// 关闭 2: 使用 range 自动遍历信道(更简洁)
	// Method 2: Use range to iterate over the channel automatically (more concise).
	var moneyList = []int{}
	for money := range moneyChan {
		moneyList = append(moneyList, money)
	}

	fmt.Println(len(moneyChan)) // channel 是有长度的(缓冲区大小)	Channel has a length (buffer size).
	fmt.Println("购买完成", moneyList)
}
