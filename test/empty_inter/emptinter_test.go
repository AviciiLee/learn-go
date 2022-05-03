package empty_inter

import (
	"fmt"
	"testing"
)

func DoSomething(p interface{}) {
	//if i, ok := p.(int); ok {
	//	fmt.Println(i)
	//}
	switch v := p.(type) {
	case int:
		fmt.Println("interge", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("unknow type", v)
	}
}

func TestDoSomething(t *testing.T) {
	DoSomething(10)
	DoSomething("20")
}
