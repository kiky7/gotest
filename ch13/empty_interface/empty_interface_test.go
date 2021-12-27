/*
 * @author:kiky
 * @date: 2021/12/27 7:29 下午
 * 空接口   断言
 * 倾向于 使用小的 接口定义，很多接口只包含 一个方法，接口 小负担 轻
 * 较大的接口定义，可以由多个小接口 定义组合而成
 * 只依赖于 必要功能的 最小 接口
**/

package empty_interface

import (
	"fmt"
	"testing"
)

//类型判断
func DoSomething(p interface{})  {
/*	if i,ok := p.(int);ok {
		fmt.Println("int：",i)
	}
	if i,ok := p.(string);ok {
		fmt.Println("string：",i)
	}
	fmt.Println("unknow type")*/

	//
	switch v:=p.(type) {
		case int:
			fmt.Println("int：",v)
		case string:
			fmt.Println("string：",v)
		default:
			fmt.Println("unknow type")
	}

}

func TestEmptyInterfaceAssertion(t *testing.T)  {
	DoSomething(1)
	DoSomething("22")
}