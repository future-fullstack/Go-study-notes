package main

import "testing"

// 一般单元测试:直接调用 Add 判断返回结果是否符合预期
// Basic unit test: call Add directly and check whether the result matches the expectation.
func TestAdd1(t *testing.T) {
	result := Add(3, 4) // 调用被测函数 Add,传入 3 和 4	Call the function under test with 3 and 4
	if result != 7 {    // 结果不是 7 则测试失败	Fail the test if the result isn't 7
		t.Errorf("测试失败")
		return
	}
	t.Logf("测试通过")
}

// 子测试(适用于测试样例不多的情况下):用 t.Run 给每个样例起名字,失败时定位更清晰
// Subtests (suitable when there aren't many test cases): name each case with t.Run so failures are easier to locate.
func TestAdd2(t *testing.T) {
	t.Run("add1", func(t *testing.T) {
		if Add(1, 2) != 3 { // 期望结果为 3	Expect the result to be 3
			t.Logf("测试失败")
			return
		}
	})

	t.Run("add2", func(t *testing.T) {
		if Add(3, 4) != 7 { // 期望结果为 7	Expect the result to be 7
			t.Logf("测试失败")
			return
		}
	})
}

// 子测试PLUS(适用于测试样例很多的情况下):把样例放进切片循环执行,新增用例只需加一行
// Subtests PLUS (suitable when there are many test cases): put the cases in a slice and run them in a loop; adding a case takes one line.
func TestAdd3(t *testing.T) {
	cases := []struct {
		Name           string // 子测试名称	Name of the subtest
		A, B, Expected int    // 入参 A、B 与期望结果 Expected	Inputs A, B and the expected result
	}{
		{"add1", 1, 2, 3},
		{"add2", 4, 8, 12},
		{"add3", -1, 2, 1},
	}
	for _, c := range cases { // 遍历所有测试样例	Iterate over all the test cases
		t.Run(c.Name, func(t *testing.T) { // 以样例名为名创建子测试	Run a subtest named after the case
			if c.A+c.B != c.Expected { // 实际结果不等于期望结果则测试失败	Fail if the actual result doesn't equal the expected one
				t.Errorf("测试失败")
			}
		})
	}
}
