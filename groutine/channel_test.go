package groutine

import (
	"fmt"
	"testing"
	"time"
)

func isCanceled(ch chan struct{}) bool {
	select {
	case <-ch:
		return true
	default:
		return false
	}
}

func cancel_1(ch chan struct{}) {
	ch <- struct{}{}
}

func cancel_2(ch chan struct{}) {
	close(ch)
}

func TestChannel(t *testing.T) {
	cancelChen := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelChen chan struct{}) {
			for {
				if isCanceled(cancelChen) {
					break
				}
				time.Sleep(time.Millisecond * 50)
			}
			fmt.Println(i, "Canceled")
		}(i, cancelChen)
	}
	cancel_2(cancelChen)
	time.Sleep(time.Second * 1)
}
