/*
 * @author:kiky
 * @date: 2022/4/25 7:24 下午
sync.Pool对象获取
1、尝试从私有对象获取
2、私有对象不存在，尝试从当前processor的共享池获取
3、如果当前processor共享池也是空的 ，那么就尝试去其他processor的共享池获取
4、如果所有子池都是空的，最后就用用户指定的new函数产生一个新的对象返回


sync.Pool对象放回
1、如果私有对象不存在则保存为私有对象
2、如果私有对象存在，放入当前processor共享池中

sync.Pool对象的生命周期
1、GC会清除sync.Pool缓存的对象
2、对象的缓存有效期为下一次GC之前

**/

package obj_pool

import "testing"

func TestObjPool(t *testing.T)  {
/*	pool := NewObjPool(10)
	for  {
		
	}*/
}