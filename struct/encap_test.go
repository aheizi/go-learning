package _struct

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id string
	Name string
	Age string
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{"0", "Bob", "21"}
	e1 := Employee{Id:"1", Name:"Jo", Age:"22"}
	e2 := new(Employee)
	e2.Age = "23"

	t.Log(e)
	t.Log(e1)
	t.Log(e2)
	t.Logf("s is %T", e)
	t.Logf("s is %T", e2)
}

func (e Employee) StringStruct() string {
	fmt.Printf("Address struct is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%s", e.Id, e.Name, e.Age)
}

func (e *Employee) StringPointer() string {
	fmt.Printf("Address pointer is %x\n", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%s", e.Id, e.Name, e.Age)
}

func TestStructOpt(t *testing.T) {
	e := Employee{"01", "Lisa", "21"}
	fmt.Printf("Native Address is %x\n", unsafe.Pointer(&e.Name))

	t.Log("pointer: ", e.StringPointer())

	t.Log("pointer: ", e.StringStruct())
}