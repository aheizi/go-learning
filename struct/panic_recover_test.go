package _struct

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover from ", err)
		}
	}()

	fmt.Println("start.")
	panic(errors.New("Something Wrong!"))
}
