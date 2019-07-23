package groutine

import (
	"fmt"
	"testing"
	"time"
)

func SelectService() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func AsyncSelectService() chan string {
	retCh := make(chan string, 1)
	go func() {
		ret := SelectService()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited")
	}()
	return retCh
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncSelectService():
		t.Log(ret)
	case <- time.After(time.Millisecond * 100):
		t.Error("time out")
	}
}
