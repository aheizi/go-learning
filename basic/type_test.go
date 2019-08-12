package basic

import "testing"

// go 不支持隐式类型转换（需要显式类型转换）
type MyInt int64

func TestImplicit(t *testing.T) {
	var a int = 1
	var b int64
	//b = a
	b = int64(a)
	t.Log(a, b)
}

// go语言不支持指针的运算，比如ptr + 1，是不支持的
func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a

	//aPtr = aPtr + 1
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

// string的初始化是空字符串，所以在判断字符串的时候，不是判断 s == nil，而是s == ""
func TestString(t *testing.T) {
	var s string
	t.Log("*" + s + "*")
	t.Log(len(s))

	//if (s == nil) {
	//}
}
