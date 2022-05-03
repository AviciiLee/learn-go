package _string

import (
	"fmt"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	s := "学习go语言"
	//for _, v := range s {
	//	log.Println(string(v))
	//}
	parts := strings.Split(s, "")
	fmt.Println(parts)
	for _, v := range parts {
		t.Log(v)
	}
}

func TestDefer(t *testing.T) {
	defer func() {
		t.Log("clear resources")
	}()
	t.Log("started")
	panic("Fatal Error")
}
