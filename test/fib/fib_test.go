package fib

import (
	"testing"
)

func TestFibList(t *testing.T) {
	const (
		Open = 1 << iota
		Close
		Pending
	)
	const (
		Monday = iota + 1
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
		Sunday
	)
	t.Log(Open, Close, Pending)
	t.Log(Monday, Tuesday, Wednesday, Thursday, Friday, Saturday, Sunday)
	a, b := 1, 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		t.Log(b)
		b, a = a+b, b
	}
}
