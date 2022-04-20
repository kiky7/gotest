/*
 * @author:kiky
 * @date: 2022/4/12 22:23 下午
 * 仅需任意任务完成，利用go CPS 并发控制机制实现
**/

package until_anyone_reply

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func runTask(id int) string  {
	time.Sleep(100 * time.Microsecond)  //如果sleep的时间太短，每次出现的结果都是最后一个
	return fmt.Sprintf("结果来自于：%d",id)
}

func AllResponse() string {
	numOfRunner := 10
	ch := make(chan string,numOfRunner) //防止协程泄露
	for i:=0;i<numOfRunner;i++ {
		go func(i int) {
			ret := runTask(i)
			ch <- ret
		}(i)
	}
	finalRet := ""
	for j:=0;j<numOfRunner;j++ {
		finalRet += <-ch + "\n"
	}
	return finalRet
}

func TestAllResponse(t *testing.T)  {
	t.Log("before:",runtime.NumGoroutine())//输出当前系统协程数
	t.Log(AllResponse())
	time.Sleep(time.Second * 1)
	t.Log("after:",runtime.NumGoroutine())
}