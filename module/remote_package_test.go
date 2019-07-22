package module

import (
	cm "github.com/easierway/concurrent_map"
	"testing"
)

func TestConcurrentMap(t *testing.T) {
	concurrentMap := cm.CreateConcurrentMap(99)
	concurrentMap.Set(cm.StrKey("key"), 10)
	t.Log(concurrentMap.Get(cm.StrKey("key")))
}
