package client

import (
	"go-learning/src/module/service"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(service.GetFibonacci(10))
	t.Log(service.Square(10))
}


