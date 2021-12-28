/*
 * @author:kiky
 * @date: 2021/12/28 7:29 下午
 *
 * panic 用于不可以恢复的错误
 * panic 退出前会执行defer指定的内容
 * os.Exit 退出时不会调用defer指定的函数
 * os.Exit 退出时不输出当前调用栈信息
 *
**/

package panic_recover_test

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T)  {
/*	defer func() {
		fmt.Println("finally!")
	}()*/

	defer func() {
		if err:= recover();err !=nil {  //recover 恢复panic
			fmt.Println("recover from..",err)
		}
	}()

	fmt.Println("开始")
	panic(errors.New("something wrong!"))
	//os.Exit(-1)
}

//当心使用recover  形成僵尸服务进程，导致health check失效
//crash 使用程序退出重启来恢复