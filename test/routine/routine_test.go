package _routine

import (
	"fmt"
	"testing"
)

func TestGoRoutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}
}
