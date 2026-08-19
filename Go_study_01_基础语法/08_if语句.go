package main

import "fmt"

// 以年龄判断为例展示Go语言的判断语句
// Demonstrates Go's conditional statement using age as an example
func main() {
	var age int
	fmt.Println("请输入年龄：")
	fmt.Scan(&age)
	////中断式	卫语句(最推荐，逻辑最清晰)	Early return guard clause(Most recommend,clearest logic)
	//if age <= 0 {
	//	fmt.Println("未出生")
	//	return
	//}
	//if age < 18 {
	//	fmt.Println("未成年")
	//	return
	//}
	//if age <= 35 {
	//	fmt.Println("青年")
	//	return
	//}
	//fmt.Println("中年")

	////嵌套写法	Nested approach
	//if age < 18 {
	//	if age <= 0 {
	//		fmt.Println("未出生")
	//	} else {
	//		fmt.Println("未成年")
	//	}
	//} else {
	//	if age <= 35 {
	//		fmt.Println("青年")
	//	} else {
	//		fmt.Println("中年")
	//	}
	//}

	////多条件式(也很推荐)		Compound condition(also recommend)
	//if age <= 0 {fmt.Println("未出生")}
	//if age < 18 && age >0 {fmt.Println("未成年")}
	//if age <= 35 && age >= 18{fmt.Println("青年")}
	//if age >35 {fmt.Println("中年")}

	//多条件式中的条件判断符号		Logic operators in multi-condition check
	//&&	都为真才是真			Logic AND
	//||	有一真就是真			Logic OR
	//!		取反					Logic NOT

	//逻辑短路	Short-circuit evaluation
	// && 第一个条件如果是 false，后面的条件就不会去走了	&&If the first condition is false,the rest will not be evaluated
	// || 第一个条件如果是 true，后面的条件就不会去走了		||If the first condition is true,the rest will not be evaluated
}
