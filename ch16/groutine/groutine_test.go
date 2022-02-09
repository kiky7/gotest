/*
 * @author:kiky
 * @date: 2022/1/23 10:08 下午
 * 协程机制
**/

package groutine__test

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T)  {
	for i :=0;i<10;i++ {
		go func(i int) { //值传递、值复制
			fmt.Println(i)
		}(i)

		//go func() {
		//	fmt.Println(i)
		//}()
	}
	time.Sleep(time.Microsecond * 50)
}


