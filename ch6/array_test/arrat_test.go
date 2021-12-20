/*
 * @author:kiky
 * @date: 2021/12/14 7:03 下午
**/

package array_test

import "testing"

func TestArrayInit(t *testing.T)  {
	var arr [3] int
	arr1 := [4]int{1,2,3,4}
	arr2 := [...]int{1,2,4,5}
	arr[1] = 5
	t.Log(arr[1],arr[2])
	t.Log(arr1,arr2)
}

func TestArrayTravel(t *testing.T)  {
	arr2 := [...]int{1,2,4,5}
	for idx,e := range arr2{  //_,e
		t.Log(idx,e)
	}

	arr3 := [3][3]int{{1,4,5},{1,2,4}}
	for idx,e := range arr3{  //_,e
		t.Log(idx,e)
	}
}

func TestArraySection(t *testing.T)  {
	arr2 := [...]int{1,2,4,5}
	arr2_s := arr2[:3]
	t.Log(arr2_s)
}