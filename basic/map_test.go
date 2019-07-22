package basic

import "testing"

func TestInitMap(t *testing.T) {
	m1 := map[int]int{1: 1, 2: 4, 3: 9}
	t.Log(m1[2])
	t.Logf("len m1=%d", len(m1))

	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("len m2=%d", len(m2))

	m3 := make(map[int]int, 10)
	t.Logf("len m3=%d", len(m3))
}

func TestAccessNotExistingKey(t *testing.T) {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])

	m1[3] = 0
	if v, ok := m1[3]; ok {
		t.Logf("key 3 value is %d", v)
	} else {
		t.Log("ke 3 is not exist.")
	}
}

func TestTravelMap(t *testing.T) {
	m1 := map[int]int{1:1, 2:4, 3:9}
	for k,v := range m1 {
		t.Log(k, v)
	}
}

func TestMapFunValue(t *testing.T) {
	m:=map[int]func(op int)int{}
	m[1]= func(op int) int {
		return op
	}
	m[2] = func(op int) int {
		return op * op
	}
	m[3] = func(op int) int {
		return op * op * op
	}

	t.Log(m[1](3), m[2](3), m[3](3))
}

func TestMapForSet(t *testing.T)  {
	mySet := map[int]bool{}

	mySet[1]= true
	n:=1
	if mySet[n]  {
		t.Logf("%d is exist.", n)
	} else {
		t.Logf("%d is not exist.", n)
	}
}