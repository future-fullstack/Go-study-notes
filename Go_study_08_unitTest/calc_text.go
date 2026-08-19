package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(3, 4)
	if result != 7 {
		t.Errorf("测试失败")
		return
	}
	t.Logf("测试通过")
}
