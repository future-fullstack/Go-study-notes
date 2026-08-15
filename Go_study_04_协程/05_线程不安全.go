package main

import (
	"fmt"
	"sync"
)

// sum 是被多个协程共享的全局变量
// sum is a global variable shared by multiple goroutines.
var sum int
var wg sync.WaitGroup

// add 循环执行 sum = sum + 1,模拟对共享变量的写操作
// add repeatedly performs sum = sum + 1, simulating writes to the shared variable.
func add() {
	for i := 1; i <= 100000; i++ {
		sum = sum + 1
	}
	wg.Done()
}

// sub 循环执行 sum = sum - 1,与 add 同时并发操作同一个变量
// sub repeatedly performs sum = sum - 1, concurrently operating on the same variable as add.
func sub() {
	for i := 1; i <= 100000; i++ {
		sum = sum - 1
	}
	wg.Done()
}

func main() {
	wg.Add(2)

	// 两个协程并发读写 sum,彼此没有任何同步,会发生数据竞争
	// The two goroutines read and write sum concurrently with no synchronization, causing a data race.
	go add()
	go sub()

	wg.Wait()
	// 由于数据竞争,最终结果不确定:理论上应等于 0,实际每次运行都不同
	// Due to the data race, the result is non-deterministic: theoretically 0, but it differs on every run.
	fmt.Println(sum) // 输出的这个数字会一直变,这里的线程是不安全的	This value keeps changing; the code is thread-unsafe.
}
