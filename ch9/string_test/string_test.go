/*
 * @author:kiky
 * @date: 2021/12/20 7:25 下午
 * string 是数据类型，不是引用或者指针类型
 * string 是只读的byte slice ,len 计算它所包含的byte数
 * string 的byte 数组可以存放任何数据
**/

package string_test

import (
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T)  {
	var s string
	t.Log(s)

	s = "hello"
	t.Log(len(s))

	s = "\xE4\xB8\xA5"
	t.Log(s)
	t.Log(len(s))

	c :=  []rune(s) // Unicode  以 rune 为单位  类似 byte
	t.Log(len(c))

	t.Logf("中 unicode  %x",c[0])
	t.Logf("中 UTF8  %x",s)
}

func TestStringFn(t *testing.T)  {
	s:= "A,B,C"
	parts := strings.Split(s,",")
	for _,part :=  range parts{
		t.Log(part)
	}
	t.Log(strings.Join(parts,"-"))
}

func TestConv(t *testing.T)  {
	s:=strconv.Itoa(10)
	t.Log("str"+s)
	if i,err := strconv.Atoi("10"); err == nil { //函数返回值有两个，因此先赋值判断再+
		t.Log(10+i)
	}
}