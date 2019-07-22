package basic

import "testing"

func TestIfMultiSec(t *testing.T)  {
	if v,err := someFun(); err==nil{
		t.Log("1==1")
	} else {
		t.Log(v)
	}
}
func someFun() (int, error) {
	return 0, nil
}

func TestSwitchMultiCase(t *testing.T)  {
	for i:=0; i< 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("unknown")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T){
	for i := 0; i< 5; i++ {
		switch {
		case i%2==0:
			t.Log("Even")
		case i%2==1:
			t.Log("Odd")
		default:
			t.Log("unknown")
		}
	}
}
