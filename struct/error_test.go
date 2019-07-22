package _struct

import (
	"errors"
	"fmt"
	"strconv"
	"testing"
)

var LessThanTwoError = errors.New("n should be not less then 2")
var LargerThanHundredError = errors.New("n should be not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	//if n < 0 || n > 100 {
	//	return nil, errors.New("n should be in [0,100]")
	//}

	if n < 2 {
		return nil, LessThanTwoError
	}

	if n > 100 {
		return nil, LargerThanHundredError
	}

	fibList := []int{1, 1}
	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

// 最佳实践
// 避免嵌套，快速失败
func GetFibonacci2(str string) {
	var (
		i    int
		err  error
		list []int
	)

	if i, err = strconv.Atoi(str); err != nil {
		fmt.Println("str cast to int error.", err)
		return
	}

	if list, err = GetFibonacci(i); err != nil {
		fmt.Println("Error.", err)
		return
	}

	fmt.Println(list)
}

func TestFibonacci(t *testing.T) {
	if v, err := GetFibonacci(-10); err != nil {
		if err == LessThanTwoError {
			fmt.Println("It is less")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}
