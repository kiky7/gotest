/*
 * @author:kiky
 * @date: 2022/4/12 20:22 下午
go语言并发任务
1、仅执行一次
Eg:单例模式（懒汉式，线程安全）
**/

package once_test

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {

}

var singleInstance *Singleton
var once sync.Once

//单例
func GetSingletonObj() *Singleton {
	once.Do(func() {
		fmt.Println("创建对象")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestGetSingletonObj(t *testing.T)  {
	var wg sync.WaitGroup //等待所有协程都运行完毕
	for i:=0;i<5;i++ {
		wg.Add(1)
		go func() {
			obj := GetSingletonObj() //多个协程创建对象，但实际只创建一次
			fmt.Printf("%d\n",unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}