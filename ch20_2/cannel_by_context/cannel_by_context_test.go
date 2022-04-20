/*
 * @author:kiky
 * @date: 2022/4/12 20:52 下午
 * Context与任务取消
根 Context ：通过context.Background()创建
子 Context ：通过context.WithCancel(parentContext)创建
	ctx,cannel := context.WithCancel(context.Background())
当Context被取消时，基于他的子context都会被取消
接收取消通知 <-ctx.Done()
**/

package cannel_by_context

import (
	"context"
	"fmt"
	"testing"
	"time"
)


func isCancelled(ctx context.Context) bool  {
	select {
	case <- ctx.Done(): //接收取消通知
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T)  {
	//cannel方法可取消该任务及其子任务
	ctx,cannel := context.WithCancel(context.Background())//创建根节点、子节点
	for i:=0;i<5;i++ {
		go func(i int,cxt context.Context) {
			for {//循环检查
				if isCancelled(ctx) { //收到消息，任务被取消
					break
				}
				time.Sleep(time.Microsecond * 5)
			}
			fmt.Println(i,"Cancelled")
		}(i,ctx)
	}
	cannel()//取消任务
	time.Sleep(time.Second * 1)
}
