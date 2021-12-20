/*
 * @author:kiky
 * @date: 2021/12/2 7:36 下午
**/

package fib

import (
	"fmt"
	"testing"
)

/**
 * @Description: 产生斐波那契数列
 * @param t
 */
func TestFibList(t *testing.T)  {
	a := 1
	b := 1
	t.Log(a)
	for i:=0;i<5;i++ {
		t.Log(b)
		a,b = b,b+a
	}
	fmt.Println()
}
