package constant_test

import (
	"testing"
)

const(
	// 简洁的定义常量,Monday = 1,后面依次加一
	Mongday = iota + 1
	Tuesday
	Wednesday
)

const(
	Readable = 1 << iota
	Writable
	Exceutable
)

func TestConstantTry(t *testing.T) {
	t.Log(Mongday, Tuesday)
}

func TestConstantTry1(t *testing.T) {
	a := 7  // 0111
	t.Log(a&Readable==Readable, a&Writable==Writable, a&Exceutable==Exceutable)
}