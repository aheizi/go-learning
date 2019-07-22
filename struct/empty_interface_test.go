package _struct

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	//if i,ok := p.(int); ok {
	//	fmt.Println("Integer: ", i)
	//	return
	//}
	//if s,ok := p.(string); ok {
	//	fmt.Println("String: ", s)
	//	return
	//}
	//fmt.Println("Unknow type.")

	switch v:=p.(type) {
	case int:
		fmt.Println("Integer: ", v)
	case string:
		fmt.Println("String: ", v)
	default:
		fmt.Println("Unknow type.")
	}
}

func TestEmptyInterfaceAssertion(t *testing.T) {
	DoSomething(10)
	DoSomething("10")
}
