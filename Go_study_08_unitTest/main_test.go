package main

import (
	"fmt"
	"os"
	"testing"
)

// setup 是所有测试开始前的准备函数:在 TestMain 中统一调用
// setup runs before all tests: it is called once from TestMain.
func setup() {
	fmt.Println("测试前")
}

// 测试函数
// 注：函数内部无内容的时候默认测试成功
// A test function. Note: a test with no body passes by default.
func TestAdd(t *testing.T) {
	fmt.Println("测试中")
	t.Errorf("测试失败")
}

// teardown 是所有测试结束后的清理函数:在 TestMain 中统一调用
// teardown runs after all tests: it is called once from TestMain.
func teardown() {
	fmt.Println("测试后")
}

// 测试入口函数(只要执行测试函数就会开始执行 TestMain 函数)
// The test entry point (TestMain runs as soon as any test function is executed).
func TestMain(m *testing.M) {
	fmt.Println("hello world")
	setup()         // 运行测试前的准备工作	Run the pre-test setup
	code := m.Run() // 开始测试(测试函数会依次执行,并返回测试信息 code )	Run the tests in order and get the exit code
	teardown()      // 运行测试后的清理工作	Run the post-test teardown
	os.Exit(code)   // 结束进程并向操作系统返回退出码（0代表通过，非0代表失败）	Exit the process with the exit code (0 = pass, non-zero = fail)
}
