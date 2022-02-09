/*
 * @author:kiky
 * @date: 2022/2/4 5:58 下午
**/

package csp

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Microsecond * 50)
	return "done"
}

func otherTask()  {
	fmt.Println("working on something else")
	time.Sleep(time.Microsecond * 100)
	fmt.Println("Task is done")
}

func TestService(t *testing.T)  {
	fmt.Println(service())
	otherTask()
}

/**
 * @Description:
 * @return chan
 */
func AsyncService() chan string {
	//retCh := make(chan string)  //普通channel
	retCh := make(chan string,1) //buf channel  channel取数不会阻塞
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret  //往channel存数据
		fmt.Println("Service exited.")
	}()
	return retCh
}

func TestAsyncService(t *testing.T)  {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh) //channel取数据
	time.Sleep(time.Second * 1)
}