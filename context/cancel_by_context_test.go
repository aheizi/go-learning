package context

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCanceled(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func TestChannel(t *testing.T) {
	ctx, cancle := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCanceled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 50)
			}
			fmt.Println(i, "Canceled")
		}(i, ctx)
	}
	cancle()
	time.Sleep(time.Second * 1)
}
