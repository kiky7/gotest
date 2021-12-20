/*
 * @author:kiky
 * @date: 2021/12/14 12:04 下午
 * 切片、map、管道 不能比较
**/

package ch4

import "testing"

//数组比较：长度、元素顺序必须 一致
func TestCompareArray(t *testing.T)  {
	a := [...]int{1,2,3,4}
	b := [...]int{1,3,2,5}
	//c := [...]int{1,2,3,4,5} //长度不同不能比较
	d := [...]int{1,2,3,4}

	t.Log( a==b)
	//t.Log( a==c)
	t.Log( a==d)

}




const(
	re = 1 << iota  //1  0001  1左移0位
	wr //2    0010  1左移1位
	ex  //4   0100  1左移2位
)

//&^ 按位置清零   右边是1则结果为0  右边为0，则结果=左边
func TestConstRWE(t *testing.T)  {
	a := 7 //0111
	a = a &^ re  // 0110
	t.Log(a)
	a = a &^ ex // 0010
	t.Log(a)
	t.Log(a&re == re, a&wr == wr , a&ex == ex)
}


