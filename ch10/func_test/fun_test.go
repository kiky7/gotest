/*
 * @author:kiky
 * @date: 2021/12/21 9:40 上午
 * 函数
 * 可以有多个返回值
 * 所有参数都是值传递，slice map channel会有传引用的错觉（实际是复制的结构包含了指针导致修改了源数据）
 * 函数可以作为变量的值、可以作为参数和返回值
 *
**/

package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValue() (int,int) {
	return rand.Intn(10),rand.Intn(20)
}

func TestFn(t *testing.T)  {
	a,b := returnMultiValue()
	t.Log(a,b)
}

//计算函数时长
func timeSpent(inner func(op int)int ) func(op int)int {
	return func(n int) int {  //内层函数便于传参
		start := time.Now()
		ret := inner(n) //执行函数
		fmt.Println("time spent:",time.Since(start).Seconds()) //耗时
		return  ret
	}
}

//另一个求耗时的思路
func timeSpentT(inner func(op int)int, n  int ) float64 {
		start := time.Now()
		inner(n) //执行函数
		return time.Since(start).Seconds()
}

func slowFuc(op int)int  {
	time.Sleep(time.Second * 1)
	return op
}

func TestTimeSpent(t *testing.T)  {
	tssF := timeSpent(slowFuc)
	t.Log(tssF(10))

	tssFT := timeSpentT(slowFuc,10)
	t.Log(tssFT)
}


//累加
func Sum(op ...int) int {
	re := 0
	for _,v := range op{
		re +=  v
	}
	return re
}


//可变长参数
func TestVarParam(t *testing.T)  {
	t.Log(Sum(1,2,3,4))
	t.Log(Sum(1,2,3,4,5))
}


func Clear()  {
	fmt.Println("清空资源")
}

//defer延迟执行、释放资源、是否锁
func TestDefer(t *testing.T)  {
	defer Clear()
	fmt.Println("开始")
	panic("err")//panic程序异常中断
	//fmt.Println("结束")
}








