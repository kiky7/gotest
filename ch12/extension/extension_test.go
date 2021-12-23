/*
 * @author:kiky
 * @date: 2021/12/23 7:28 下午
**/

package extension

import (
	"fmt"
	"testing"
)

//父类
type Pet struct {

}
//父类方法1
func (p *Pet) Speak() {
	fmt.Print("...")
}
//父类方法2
func (p *Pet) SpeakTo(host string)  {
	p.Speak()
	fmt.Println(" ",host)
}

type Dog struct {
	//p *Pet
	Pet //继承的感觉，子类不用写方法，可以直接使用; 子类重载无效，无法使用子类，因此没有继承
}

//子类方法1
func (d *Dog) Speak() {
	//d.p.Speak()
	fmt.Print("Wang!")
}
//子类方法2
/*func (d *Dog) SpeakTo(host string)  {
	//d.p.SpeakTo(host)
	d.Speak()
	fmt.Println(" ",host)
}*/

func TestDog(t  *testing.T)  {
	dog := new(Dog)
	//var dog Pet := new(Dog)  // XX 无法支持【子类交换原则LSP】
	dog.SpeakTo("chao")
}


//匿名嵌套类型