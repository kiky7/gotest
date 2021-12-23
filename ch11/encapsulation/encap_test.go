/*
 * @author:kiky
 * @date: 2021/12/22 7:04 下午
 * go是面向对象语言吗？没有继承、有 数据&行为的封装
**/

package encapsulation

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id string
	Name string
	Age int
}

func TestCreateEmployeeObj(t *testing.T)  {
	e := Employee{"0","李四",20}
	e1 := Employee{Name: "张三",Age: 30}
	e2 := new(Employee) //返回指针
	//t.Log(e2)
	e2.Id = "2"
	e2.Age = 25
	e2.Name = "Rose"
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)  //输出与视频案例有差异 ，多了& 符
	t.Logf("e is %T",e)  //%T 类型
	t.Logf("e2 is %T",e2)
}

//第一种定义方式在实例对应方法被调用时，实例的成员会进行复制
func (e Employee) String() string {
	fmt.Printf("地址是： %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d",e.Id,e.Name,e.Age)
}

//通常情况下为了避免内存拷贝我们使用第二种定义方式
func (e *Employee) StringT() string {
	fmt.Printf("地址是： %x", unsafe.Pointer(&e.Name))
	return fmt.Sprintf("ID:%s-Name:%s-Age:%d",e.Id,e.Name,e.Age)
}


func TestStructOperations(t  *testing.T)  {
	e  := Employee{"0","baby",20} //直接调用实例的方法
	//e  := &Employee{"0","baby",20}  //指向实例的指针、也可以直接调用实例的方法
	fmt.Printf("地址是： %x\n", unsafe.Pointer(&e.Name))  //输出不同，存在地址复制
	t.Log(e.String())

	//ee  := &Employee{"0","baby2",20} //不管是指针访问还是实例访问都可以，区别在哪？
	ee  := Employee{"0","baby2",20}
	fmt.Printf("地址是： %x\n", unsafe.Pointer(&ee.Name)) //输出相同
	t.Log(ee.StringT())
}