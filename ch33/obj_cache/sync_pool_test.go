/*
 * @author:kiky
 * @date: 2022/4/25 7:45 下午
 * 对象池测试
**/

package obj_cache

import (
	"fmt"
	"gotest/ch32/obj_pool"
	"testing"
	"time"
)

func TestObjPool(t *testing.T)  {
	pool := obj_pool.NewObjPool(10)
	if err := pool.ReleaseObj(&obj_pool.ResuableObj{});err != nil { //尝试放置超出线程池大小，失败
		t.Error(err)
	}
	for i:=0; i<11;i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil{
			t.Error(err)
		}else {
			fmt.Printf("%T\n",v)
			if err := pool.ReleaseObj(v);err != nil {
				t.Error(err)
			}
		}
	}
	fmt.Println("Done")
}