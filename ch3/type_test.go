/*
 * @author:kiky
 * @date: 2021/12/14 11:18 上午
 * 不支持隐式类型转换
**/

package ch3

import "testing"

type Myint int64

func TestImplicit(t *testing.T)  {
	var a int32 = 1
	var b int64
	b = int64(a)
	t.Log(a,b)

	var c Myint
	c = Myint(b)
	t.Log(a,b,c)
}

func TestPoint(t *testing.T)  {
	a := 1
	aptr := &a
	t.Log(a,aptr)
	t.Logf("%T %T",a,aptr)
}

func TestString(t *testing.T)  {
	var s string
	t.Log("*"+s +"*")
	if s == "" {
		t.Log(len(s))
	}
}