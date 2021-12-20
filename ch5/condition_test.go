/*
 * @author:kiky
 * @date: 2021/12/14 2:19 下午
**/

package ch5

import "testing"

//condition结果必须是布尔值
//
func TestCondition(t *testing.T)  {
	if a:= 1==2;a {
		t.Log("1==1")
	}else {
		t.Log(a)
	}

/*	if v,err := somefuc();err ==nil {
		t.Log("1==1")
	}else {
		t.Log("1==1")
	}*/
}


func TestSwitchCase(t *testing.T)  {
	for i:=0;i<5;i++ {
		switch i {
		case 0,2:
			t.Log("偶数")
		case 1,3:
			t.Log("奇数")
		default:
			t.Log("其他数")
		}
	}
}

func TestSwitchCaseCondition(t *testing.T)  {
	for i:=0;i<5;i++ {
		switch {
		case i%2==0:
			t.Log("偶数")
		case i%2==1:
			t.Log("奇数")
		default:
			t.Log("其他数")
		}
	}
}