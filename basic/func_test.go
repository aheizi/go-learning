package basic

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 多返回值
func TestReturnMultiValue(t *testing.T) {
	a, _ := returnMultiValue()
	t.Log(a)
}

func returnMultiValue() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

// 返回函数，装饰器模式
func TestFn(t *testing.T) {
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}

func Sum(ops ...int) int {
	ret := 0
	for _, i := range ops {
		ret += i
	}
	return ret
}

// 可变参数
func TestVarParam(t *testing.T) {
	t.Log(Sum(1, 2, 3))
	t.Log(Sum(1, 2, 3, 4, 5))
}

func clear() {
	fmt.Println("clear resources.")
}

func TestDefer(t *testing.T) {
	defer clear()
	fmt.Println("start")
	panic("err")
	// 这里不会被执行到
	fmt.Println("end")
}