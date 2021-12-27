/*
 * @author:kiky
 * @date: 2021/12/27 7:45 下午
**/

package error

import (
	"errors"
	"fmt"
	"testing"
)

var LessThanTwoError = errors.New("n大于2")
var LargeThenHundedError = errors.New("n小于100")

func GetFibonacci(n int) ([]int,error)  {

	if n<2  {
		return nil,LessThanTwoError
	}
	if n>100 {
		return nil,LargeThenHundedError
	}

	fibList := []int{1,1}

	for i := 2;i<n;i++ {
		fibList = append(fibList,fibList[i-2]+fibList[i-1])
	}
	return fibList,nil
}

//及早失败，避免嵌套！
func TestGetFibonacci(t *testing.T)  {
	if v,err := GetFibonacci(1);err !=nil {
		if err== LessThanTwoError {
			fmt.Println("太小了！")
		}
		t.Error(err)
	}else {
		t.Log(v)
	}
}