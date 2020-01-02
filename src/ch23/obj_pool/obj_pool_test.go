package obj_pool

import (
	"fmt"
	"testing"
	"time"
	"unsafe"
)

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	//初始化已预置满对象，尝试放置超出池大小的对象
	//if err := pool.ReleaseObj(&ReusableObj{}); err != nil {
	//	t.Error(err) //返回overflow
	//}
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err) //注释掉下方if语句，10个对象取完后，没有对象可取，返回time out
		} else {
			//如果可重用对象结构体没有定义任何成员，其所有对象地址都是一样的
			//加了一个id int子成员后，对象地址依次加8
			fmt.Printf("第%d个对象为%T，其地址为%d，其id为%d\n", i+1, v, unsafe.Pointer(v), v.id)
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}
	}
	fmt.Println("done")
}
