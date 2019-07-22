package _struct

import (
	"testing"
)

type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {

}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\"hello world!\")"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}


// Go接口最佳实践
// 1. 倾向于使用小的接口定义
// 2. 较大的接口，可以由很多小接口组合而成
// 3. 只依赖于必要功能的最小接口