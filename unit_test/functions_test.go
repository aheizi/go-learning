package aaa

import (
	"testing"
)

func TestSquare(t *testing.T) {
	inputs := [...]int{1,2,3,4}
	expected := [...]int{1, 4, 9, 16}

	for i := 0; i< len(inputs); i++{
		ret := square(inputs[i])
		if ret != expected[i] {
			t.Errorf("input is %d, output is %d, expected is %d.", inputs[i], ret, expected[i])

		}
	}
}




