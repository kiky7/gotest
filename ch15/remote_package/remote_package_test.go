/*
 * @author:kiky
 * @date: 2022/1/3 10:35 上午
 * 如何使用第三方的package
https://github.com/easierway/concurrent_map
**/

package remote_package

import (
	cm "github.com/easierway/concurrent_map"
	"testing"
)

func TestConcurrentMap(t *testing.T)  {
	m:= cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("key"),10)
	t.Log(m.Get(cm.StrKey("key")))
}



/*go未解决的依赖问题
1、同一个环境下，不同项目使用同一包的不同版本
2、无法管理对包的特定版本的依赖

vendor目录
获取依赖包，逐级去查找
godep
glide
dep

*/