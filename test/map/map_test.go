package _map

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m1 := map[int]int{1: 2, 2: 3}
	if v, ok := m1[2]; ok {
		t.Log(v)
	} else {
		fmt.Print("no value of this key")
	}
	for k, v := range m1 {
		t.Log(k, v)
	}
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	if mySet[1] {
		t.Logf("%d is existing", 1)
	} else {

	}
}
