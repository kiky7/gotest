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
	time.Sleep(time.Second * 5)
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
	//retCh := make(chan string)  //普通channel、会阻塞
	retCh := make(chan string,1) //buf channel,增加容量参数  channel取数不会阻塞
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
	otherTask() //
	fmt.Println(<-retCh) //channel取数据
	time.Sleep(time.Second * 1)
}


//26--无法像教程内容，达到超时执行第二个case的效果
func TestSelect(t *testing.T)  {
/*	ret := <-AsyncService();
	t.Log(ret)*/
	//timer := time.After(time.Millisecond * 10) // timer
	select {
	case ret := <-AsyncService() :
		t.Log(ret)
	case <-time.After(time.Second * 1) :
		t.Error("time out")
	}
}