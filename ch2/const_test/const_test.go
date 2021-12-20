/*
 * @author:kiky
 * @date: 2021/12/2 7:58 下午
**/

package const_test

import (
	"testing"
)

const(
	monday = iota +1  ///iota 从0开始
	tuesday
	wednesday
)

const(
	re = 1 << iota  //1  0001  1左移0位
	wr //2    0010  1左移1位
	ex  //4   0100  1左移2位
)

func TestConstTry(t *testing.T)  {
	t.Log(monday,tuesday)
}

//状态位标识、位运算
func TestConstRWE(t *testing.T)  {
	a := 1 //0001
	t.Log(re,wr,ex)
	t.Log(a&re == re, a&wr == wr , a&ex == ex)
}
