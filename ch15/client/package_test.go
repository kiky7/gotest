/*
 * @author:kiky
 * @date: 2022/1/3 10:22 上午
 * packeage
 * 1、基本复用模块单元
 * 以首字母大写来 表明可被包外代码访问
 * 2、代码的package可以和所在的目录不一致
 * 3、同一目录里的go代码的package要保持一致
**/

package client

import (
	"gotest/ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSeries(5))
	series.Sequare(2) //大写开头才能被包外访问
}



/*
init方法
在main被执行之前，所有依赖的package的init方法都会被执行
不同包的init函数按照包导入的依赖关系决定执行顺序
每个包可以有多个init函数
包的每个源文件也可以用有多个init函数，这点比较特殊
*/


/*
如何使用第三方的package
*/

