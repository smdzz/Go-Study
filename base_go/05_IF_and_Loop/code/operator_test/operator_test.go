package operator_test

import "testing"

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	// 长度不同,无法进行比较
	//t.Log(a == c)
	t.Log(a == d)
}

const(
	Readable = 1 << iota
	Writable
	Exceutable
)

func TestConstantTry1(t *testing.T) {
	// 按位清零示意
	a := 7  // 0111
	t.Log(a)
	// a 0111  Readable 0001   清零之后变为 0110
	a = a &^ Readable
	t.Log(Readable, Exceutable)
	t.Log(a)
	a = a &^ Exceutable
	t.Log(a)
	t.Log(a&Readable==Readable, a&Writable==Writable, a&Exceutable==Exceutable)
}