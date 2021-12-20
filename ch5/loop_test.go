/*
 * @author:kiky
 * @date: 2021/12/14 2:15 下午
**/

package ch5

import "testing"

func TestWhileLoop(t *testing.T)  {
	n := 0
	for n<5 {
		t.Log(n)
		n++
	}
}

