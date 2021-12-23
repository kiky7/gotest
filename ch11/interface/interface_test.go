/*
 * @author:kiky
 * @date: 2021/12/23 6:54 下午
 * 接口依赖
 * 接口的实现不依赖于接口的定义，采用 duck type 方式
 *
 * 接口为非入侵性，实现不依赖于接口定义
 * 所以接口的定义可以包含在接口使用者包内
**/

package _interface

import "testing"

//接口
type Programmer interface {
	WriteHelloWord() string
}

//
type GoProgrammer struct {

}

//实现接口   //duck type 方法、签名一致
func (g *GoProgrammer) WriteHelloWord( ) string{
	return "fmt.Println(\"hello world\")"
}

func TestClient(t *testing.T)  {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWord())
}

//接口变量 var prog coder = & GoProgrammer{}


