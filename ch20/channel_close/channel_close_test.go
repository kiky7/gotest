/*
 * @author:kiky
 * @date: 2022/3/20 4:46 下午
 * 所有的channel接收者都会在channel关闭时，立刻从阻塞等待中返回上述ok值为false。
 * 这个广播机制常被利用，进行向多个订阅者同时发送信号。如：退出信号。
**/

package channel_close

import (
	"fmt"
	"sync"
	"testing"
)

//生产者
func dataProducer(ch chan int,wg *sync.WaitGroup)  {
	go func() {
		for i:=0;i<10;i++ { //数量不定，可以用特定值代表结束，但多个receiver则有冲突
			ch <- i
		}
		wg.Done()
	}()
}

//接收者
func dataReceiver(ch chan int,wg *sync.WaitGroup)  {
	go func() {
		for i:=0;i<10;i++ { //数量不定
			data := <-ch
			fmt.Println(data)
		}
		wg.Done()
	}()
}

//生产者
func dataProducerT(ch chan int,wg *sync.WaitGroup)  {
	go func() {
		for i:=0;i<10;i++ { //数量不定，可以用特定值代表结束，但多个receiver则有冲突
			ch <- i
		}
		close(ch)//关闭通道
		//ch<-11 //向关闭的chanel发送数据，会导致panic
		wg.Done()
	}()
}

//接收者
func dataReceiverT(ch chan int,wg *sync.WaitGroup)  {
	go func() {
		for { //数量不定
			if data,ok := <-ch ; ok{//true 表示正常接收，false表示通道关闭
				fmt.Println(data)
			}else {
				break
			}
		}
		wg.Done()
	}()
}


func TestCloseChannel(t *testing.T)  {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducer(ch,&wg)
	wg.Add(1)
	dataReceiver(ch,&wg)
	wg.Wait()
}

func TestCloseChannelT(t *testing.T)  {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	dataProducerT(ch,&wg)
	wg.Add(1)
	dataReceiverT(ch,&wg) //可以有多个receiver
	wg.Add(1)
	dataReceiverT(ch,&wg)
	wg.Wait()
}