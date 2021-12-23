/*
 * @author:kiky
 * @date: 2021/12/23 7:12 下午
 * 自定义类型
**/

package customer_type

import (
	"fmt"
	"testing"
	"time"
)

//定义函数类型
type InConv  func(op int)int

//计算函数时长
func timeSpent(inner InConv ) InConv {
	return func(n int) int {  //内层函数便于传参
		start := time.Now()
		ret := inner(n) //执行函数
		fmt.Println("time spent:",time.Since(start).Seconds()) //耗时
		return  ret
	}
}


func slowFuc(op int)int  {
	time.Sleep(time.Second * 1)
	return op
}

func TestTimeSpent(t *testing.T)  {
	tssF := timeSpent(slowFuc)
	t.Log(tssF(10))
}