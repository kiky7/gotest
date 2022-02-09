/*
 * @author:kiky
 * @date: 2022/1/23 10:18 下午
 * 共享内存并发机制--锁
**/

package share_mem_test

import (
	"sync"
	"testing"
	"time"
)

/**
 * @Description: 非线程安全计数
 * @param t
 */
func TestCounter(t *testing.T)  {
	couter :=0
	for i:=0;i<5000;i++ {
		go func() {
			couter++ //非线程安全
		}()
	}
	time.Sleep(1 * time.Second)
	t.Logf("counter = %d",couter)
}

/**
 * @Description: 线程安全计数--加锁
 * @param t
 */
func TestCounterThreadSafe(t *testing.T)  {
	var mut sync.Mutex //锁
	couter :=0
	for i:=0;i<5000;i++ {
		go func() {
			defer func() {  //函数执行后再执行
				mut.Unlock()
			}()
			mut.Lock()
			couter++ //非线程安全
		}()
	}
	time.Sleep(1 * time.Second) //等待其他协程执行完,无法预知其他协程运行时间，强制等待1S
	t.Logf("counter = %d",couter)
}

/**
 * @Description:线程安全计数--更好的方案、节约耗时
 * @param t
 */
func TestCounterWaitGroup(t *testing.T)  {
	var mut sync.Mutex //锁
	var wg sync.WaitGroup
	couter :=0
	for i:=0;i<5000;i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
			}()
			mut.Lock()
			couter++
			wg.Done()
		}()
	}
	wg.Wait()//阻塞，等待所有协程结束
	t.Logf("counter = %d",couter)
}


//CSP并发控制
/*
channel两种机制
1、发送、接收消息双方阻塞
2、buf channel 有容量，发送消息的容量够的可以放，接收消息只要有消息，都可以获取执行
*/