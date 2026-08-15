package main

import (
	"fmt"
	"sync"
)

var sum1 int
var wg1 sync.WaitGroup
var lock sync.Mutex // 创建一把锁,用于保护共享变量	Create a lock to protect the shared variable.

// 注意下述函数的内部,都存在 lock.Lock() 和 lock.Unlock()
// Note that both functions below contain lock.Lock() and lock.Unlock().
// 因为下述函数会用 go 启动为协程,两个协程必然有先后拿到锁:先拿到的先执行、再解锁;
// 一个函数未解锁前,另一个函数不能进入临界区执行
// Because both functions are launched as goroutines with go, they acquire the lock one after another:
// whoever gets the lock first runs first and then unlocks; until a function unlocks,
// the other function cannot enter the critical section.
func add1() {
	lock.Lock() // 加锁:同一时刻只允许一个协程进入临界区	Lock: only one goroutine may enter the critical section at a time.
	for i := 1; i <= 100000; i++ {
		sum1 = sum1 + 1
	}
	lock.Unlock() // 解锁:放行下一个协程	Unlock: let the next goroutine in.
	wg1.Done()
}

func sub1() {
	lock.Lock()
	for i := 1; i <= 100000; i++ {
		sum1 = sum1 - 1
	}
	lock.Unlock()
	wg1.Done()
}

func main() {
	wg1.Add(2)

	go add1()
	go sub1()

	wg1.Wait()
	// 互斥锁保证了对 sum1 的读写互斥执行,结果稳定为 0
	// The mutex guarantees exclusive access to sum1, so the result is stable at 0.
	fmt.Println(sum1) // 加锁后数字稳定不变(始终为 0),线程安全	With the lock, the value stays stable (always 0); thread-safe.
}
