package obj_pool

import (
	"errors"
	"time"
)

type ReusableObj struct { //一个可重用对象结构体
	id int
}

type ObjPool struct { //一个对象池结构体
	bufChan chan *ReusableObj //用于缓存可重用对象指针的信道
}

//创建对象池，并指定大小，返回对象池指针
func NewObjPool(numOfObj int) *ObjPool {
	ObjPool := ObjPool{}                                //创建一个对象池对象
	ObjPool.bufChan = make(chan *ReusableObj, numOfObj) //创建可重用对象指针的信道
	//在信道中预置对象，可以预置一些比较难以创建的对象在里面
	for i := 0; i < numOfObj; i++ {
		ReusableObj := &ReusableObj{}
		ReusableObj.id = i + 1
		ObjPool.bufChan <- ReusableObj
	}
	return &ObjPool
}

//如果想放置任何对象，将*ReusableObj换成空接口obj interface{}，但是每次取对象时需要进行判断是什么对象，所以并不推荐。
//使用不同池缓存不同对象
func (p *ObjPool) GetObj(timeout time.Duration) (*ReusableObj, error) {
	select {
	case ret := <-p.bufChan: //正确取到了对象
		return ret, nil
	case <-time.After(timeout): //超时仍没有取到对象，slow response比quick failure更可怕
		return nil, errors.New("time out")
	}
}

//把对象放回池中
func (p *ObjPool) ReleaseObj(obj *ReusableObj) error {
	select {
	case p.bufChan <- obj: //正确放回池中
		return nil
	default: //池已满，无法放回，返回溢出错误
		return errors.New("overflow")
	}
}
