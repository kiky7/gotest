/*
 * @author:kiky
 * @date: 2021/12/27 7:16 下午
 * 多态
**/

package polymorphsim

import (
	"fmt"
	"testing"
)

type Code string

//接口
type Programmer interface {
	WriteHelloWord() Code
}

func WriteFirstProgram(p  Programmer)  {
	fmt.Printf("%T %v\n",p,p.WriteHelloWord())
}

//go
type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWord( ) Code{
	return "fmt.Println(\"hello world\")"
}

//java
type JavaProgrammer struct {
}

func (g *JavaProgrammer) WriteHelloWord( ) Code{
	return "System.out.Println(\"hello world\")"
}

//多态测试
func TestPolymorphism(t *testing.T)  {

	goProg :=  new(GoProgrammer)  //&GoProgrammer{}  必须是指针类型
	javaProg := new(JavaProgrammer)
	WriteFirstProgram(goProg)
	WriteFirstProgram(javaProg)
}


