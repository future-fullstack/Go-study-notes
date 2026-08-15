// 向上抛的核心含义:底层函数不对错误进行处理,而是把错误向上层返回,由上层统一处理
// The essence of propagating errors upward: low-level functions do not handle errors;
// they return the error to the upper layer, which handles it centrally.
package main

import (
	"errors"
	"fmt"
)

// div 是底层函数:除数为 0 时返回错误,而不是自己处理
// div is the low-level function: when the divisor is 0 it returns an error instead of handling it itself.
func div(a, b int) (res int, err error) {
	if b == 0 {
		err = errors.New("除数不能为0") // 构造并向上返回一个错误	Construct and return an error upward.
		return                     // 返回错误(商为零值 0)	Return the error (the quotient is the zero value 0).
	}
	res = a / b
	return // 正常情况:返回商和 nil 错误	Normal case: return the quotient and a nil error.
}

// sever 是上层函数:调用 div 后检查错误,若错误非空则继续向上抛
// sever is the upper-level function: it calls div, and if the error is non-nil it keeps propagating it upward.
func sever() (res int, err error) {
	res, err = div(3, 0)
	if err != nil {
		return // 把 div 的错误继续向上抛,不执行后续逻辑	Propagate div's error upward and skip the rest.
	}
	// 只有错误为 nil 时才继续执行其他逻辑
	// Only when the error is nil do we continue with the rest of the logic.
	res++
	res += 2

	return
}

func main() {
	res, err := sever()
	if err != nil {
		fmt.Println(err) // 最上层最终打印错误并退出	The top level finally prints the error and exits.
		return
	}
	fmt.Println(res)
}
