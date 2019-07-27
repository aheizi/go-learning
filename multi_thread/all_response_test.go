package multi_thread

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTaskA(i int) string {
	time.Sleep(time.Millisecond * 50)
	return fmt.Sprintf("This result id from %d.", i)
}

func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string, numOfRunner)
	for i := 0; i < numOfRunner; i++ {
		go func(i int) {
			ret := runTaskA(i)
			ch <- ret
		}(i)
	}
	finalRet := ""
	for j := 0; j < numOfRunner; j++ {
		finalRet += <-ch + "\n"
	}
	return finalRet
}

func TestAllResponse(t *testing.T) {
	t.Log("Before:", runtime.NumGoroutine())
	t.Log(AllResponse())
	time.Sleep(time.Second * 1)
	t.Log("After", runtime.NumGoroutine())
}
