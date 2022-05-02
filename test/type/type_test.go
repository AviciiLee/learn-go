package _type

import "testing"

type MyInt int64

//1.没有隐式类型转换
//2.不支持指针运算
//3.string为值类型，默认值是空字符串
func TestInternalType(t *testing.T) {
	var a int32 = 1
	var b int64
	var c MyInt
	b = int64(a)
	c = MyInt(b)
	t.Log(a, b, c)
}
