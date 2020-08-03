package main

import "testing"

func TestAddNum(t *testing.T) {
	res := AddNum(10)
	if res != 55 {
		t.Fatalf("AddNum(10)执行错误,期望值=%v 实际值=%v\n", 55, res)
	}
	t.Log("AddNum(10)执行正确...")
}

func TestHello(t *testing.T) {
	t.Log("test hello")
}