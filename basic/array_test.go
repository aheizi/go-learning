package basic

import "testing"

/**
	测试数组遍历
 */
func TestArrayTravel(t *testing.T) {
	arr3 := [...]int{1,2,3,4,5}
	for i := 0; i < len(arr3); i++ {
		t.Log(arr3[i])
	}

	for idx,e := range arr3 {
		t.Log(idx, e)
	}

	for _,e := range arr3 {
		t.Log(e)
	}
}

/**
	测试数据选取
 */
func TestArraySection(t *testing.T) {
	arr3 := [...]int{1, 2,3,4,5}
	arr3_sec := arr3[3:]
	t.Log(arr3_sec)
}