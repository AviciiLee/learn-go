package _error

import (
	"errors"
	"simple-http-server/demo"
	"testing"
)

func TestError(t *testing.T) {
	defer func() {
		t.Log("defer")
		if err := recover(); err != nil {
			t.Log(err)

		}
	}()
	demo.Demo()
	panic(errors.New("something wrong"))
}
