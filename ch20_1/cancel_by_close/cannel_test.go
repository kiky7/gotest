/*
 * @author:kiky
 * @date: 2022/4/12 19:18 上午
 * 任务取消
**/

package cancel_by_close

import (
	"fmt"
	"testing"
	"time"
)

//任务是否已被取消
//实现原理：
//检查是否从 channel 收到一个消息，如果收到一个消息，我们就返回 true，代表任务已经被取消了
//当没有收到消息，channel 会被阻塞，多路选择机制就会走到 default 分支上去。
func isCancelled(cancelChan chan struct{}) bool  {
	select {
	case <- cancelChan: //收到消息
		return true
	default:
		return false
	}
}

func cancel_1(cancelChan chan struct{})  {
	for i:=0;i<2;i++ {
		cancelChan <- struct{}{} //往channel里发消息
	}
}

func cancel_2(cancelChan chan struct{})  {
	close(cancelChan) //关闭、阻塞、自带广播，所有协程都会收到消息
}

//利用 CSP, 多路选择机制和 channel 的关闭与广播实现任务取消功能
func TestCancel(t *testing.T)  {
	cancelChan := make(chan struct{},0)
	for i:=0;i<5;i++ {
		go func(i int,cancelCh chan struct{}) {
			for {//循环检查
				if isCancelled(cancelChan) { //收到消息，任务被取消
					break
				}
				time.Sleep(time.Microsecond * 5)
			}
			fmt.Println(i,"Cancelled")
		}(i,cancelChan)
	}
	cancel_2(cancelChan)
	//cancel_1(cancelChan)
	time.Sleep(time.Second * 1)
}