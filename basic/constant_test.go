package basic

import "testing"

const (
	Readable = 1 << iota
	Write
	Executable
)

// 可以用来定义mask
func TestConstantTry(t *testing.T) {
	a := 1
	t.Log(a&Readable == Readable, a&Write == Write, a&Executable == Executable)
}


