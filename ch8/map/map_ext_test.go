/*
 * @author:kiky
 * @date: 2021/12/16 7:49 下午
**/

package map_ext

import (
	"testing"
)

//map 实现工厂
func TestMaoWithFunValue(t *testing.T)  {
	m1:=map[int]func(op int)int{}
	m1[1] = func(op int) int{return op}
	m1[2] = func(op int) int{return op*op}
	m1[3] = func(op int) int{return op*op*op}
	t.Log(m1[1](2),m1[2](2),m1[3](2))

}


// map实现set  map[type]bool
func TestMapForSet(t *testing.T) {
	myset := map[int]bool{}
	myset[1] = true
	n:=3
	if myset[n] {
		t.Logf("%d n exit",n)
	}else {
		t.Logf("%d n no exit",n)
	}
	myset[3] = true
	t.Log(len(myset))
	delete(myset,1)

	n=1
	if myset[n] {
		t.Logf("%d n exit",n)
	}else {
		t.Logf("%d n no exit",n)
	}
}
