/*
 * @author:kiky
 * @date: 2021/12/14 7:22 下午
 * 容量是否可伸缩、是否可比较
 * 切片可伸缩、不可比较   数组可比较、不可伸缩
**/

package slice_test

import "testing"

//ptr指针   len个数  cap 容量
func TestSliceInit(t *testing.T)  {
	var  s0 [] int
	t.Log(len(s0),cap(s0))
	s0 = append(s0,1) //长度不够会申请新的空间，因此需要赋值
	t.Log(len(s0),cap(s0))

	s1 := []int{1,2,4,5}
	t.Log(len(s1),cap(s1))

	s2 := make([]int,3,5)
	t.Log(len(s2),cap(s2))
	t.Log(s2[0],s2[1],s2[2])

	s2 = append(s2,1)
	t.Log(s2[0],s2[1],s2[2],s2[3])
	t.Log(len(s2),cap(s2))
}

//切片如何实现可边长
func TestSliceGrowing(t *testing.T)  {
	s := []int{}
	for i:=0;i<10;i++{
		s = append(s,i) //长度不够会申请新的空间，因此需要赋值
		t.Log(len(s),cap(s))  //cap不够了，则cap = cap*2
	}
}

//切片共享存储空空间
func TestSliceShareMemory(t *testing.T)  {
	year :=  []int{1,2,3,4,5,6,7,8,9}
	q2 := year[3:6]
	t.Log(q2,len(q2),cap(q2))

	summer := year[5:8]
	t.Log(summer,len(summer),cap(summer))

	summer[0] = 77
	t.Log(q2)
	t.Log(year)
}

/*func TestSliceCompare(t *testing.T)  {
	a:= []int{1,2,3,4}
	b:= []int{1,2,3,4}
	if a==b { //
		t.Log("相等")
	}
}*/