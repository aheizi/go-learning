package basic

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringToRune(t *testing.T)  {
	s := "中华人民共和国"
	for _, c:=range s {
		t.Logf("%[1]c %[1]x", c)
	}
}

func TestStringFun(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _,part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestConvert(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	if i, err:=strconv.Atoi(s); err == nil {
		t.Log(10 + i)
	}

}