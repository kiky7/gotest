/*
 * @author:kiky
 * @date: 2021/12/16 7:34 下午
**/

package map_test

import (
	"testing"
)

func TestInitMap(t *testing.T)  {
	m1 := map[int]int{1:1,2:4,3:9}
	t.Log(m1[2])
	t.Logf("m1=%d",len(m1))
	m2 := map[int]int{}
	m2[4] = 16
	t.Logf("m2=%d",len(m2))
	m3 := make(map[int]int,10)
	t.Logf("m3=%d",len(m3))
}

//不存在或者赋值为0，输出都是0，如何区分？ if判断 v，ok
func TestAccNek(t *testing.T)  {
	m1 := map[int]int{}
	t.Log(m1[1])
	m1[2] = 0
	t.Log(m1[2])

	if v, ok := m1[2]; ok {
		t.Log(v)
	}else {
		t.Log("k  n  e")
	}
}


func TestTravelMap(t *testing.T)  {
	m1 := map[int]int{1:1,2:4,3:9}
	for idx,v := range m1{
		t.Log(idx,v)
	}
}

