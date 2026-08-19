package main

import (
	"fmt"
)

func main() {
	/*
		 for循环使用方法
		Usage of the for loop
		for 初始化 ; 停止条件 ; 每次循环时对 i 的操作
		for initialization;condition;post-operation
	*/

	// e.g. 从 1 加到 100
	// e.g. sum from 1 to 100
	var sum1 int
	for i := 1; i <= 100; i = i + 1 {
		sum1 = sum1 + i
	}
	fmt.Println(sum1)

	/* 	 Dead loop
	// 死循环
	// Infinite loop
	for i := 1; true; i = i + 1 {
		fmt.Println(time.Now())
		time.Sleep(2 * time.Second) //时间停止两秒
	}

	// 死循环特殊写法 1
	// Special infinite loop syntax
	for true {
		fmt.Println(time.Now())
		time.Sleep(2 * time.Second)
	}

	// 死循环特殊写法 2
	// Special infinite loop syntax
	for {
		fmt.Println(time.Now())
		time.Sleep(2 * time.Second)
	}
	*/

	/*
		小知识； 关于 time.sleep() 函数
		Tip:About the time.sleep() function
		time.sleep(d*time.Duration)
	*/

	// while 模式的循环 (Go 中没有while循环)
	// While-style loop(Go doesn't have a while loop)
	var sum2 int
	var k int
	for k < 100 {
		k = k + 1
		sum2 = sum2 + k
	}
	fmt.Println(sum2)

	// do while 模式
	//Loop that execute at least once
	var i = 1
	var sum = 0
	for {
		sum = sum + i
		i++
		if i == 101 {
			break
		}
	}
	fmt.Println(sum)

	// slice map 的遍历方式
	// Iterating over slice and map
	slc := []int{1, 2, 3}
	emap := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range emap {
		fmt.Println(k, v)
	}
	for i, v := range slc {
		fmt.Println(i, v)
	}

	// 9*9 乘法表的实现
	//Implementation of the 9*9 multiplication table
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%dx%d=%d\t", i, j, i*j)
		}
		fmt.Println()
	}

	/*
		break 		循环直接结束
		break		Exits the loop
		continue 	跳过本次循环
		continue		Skip to the next iteration
	*/
}
