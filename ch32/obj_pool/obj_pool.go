/*
 * @author:kiky
 * @date: 2022/4/25 7:24 下午
 * buf channel 对象池
**/

package obj_pool

import (
	"errors"
	"time"
)

type ResuableObj struct {

}

type ObjPool struct {
	bufChan chan *ResuableObj //用于缓存可重用对象
}

//创建对象池
func NewObjPool(numOfObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ResuableObj,numOfObj)

	for i:=0; i<numOfObj; i++ {
		objPool.bufChan  <- &ResuableObj{}
	}
	return &objPool
}

//获取对象
func (p *ObjPool)GetObj(timeout time.Duration) (*ResuableObj,error) {
	select {
	case ret:= <-p.bufChan:
		return ret,nil
	case <-time.After(timeout): //超时控制
		return nil,errors.New("time out")
	}
}

//不建议使用--不限制对象类型
/*func (p *ObjPool)GetObj2(timeout time.Duration) (obj interface{},error) {
	select {
	case ret:= <-p.bufChan:
		return ret,nil
	case <-time.After(timeout): //超时控制
		return nil,errors.New("time out")
	}
}*/


//释放对象，放入对象池
func (p *ObjPool) ReleaseObj(obj *ResuableObj) error  {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}






