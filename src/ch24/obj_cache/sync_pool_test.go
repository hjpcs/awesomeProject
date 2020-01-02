package obj_cache

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func TestSyncPool(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create a new object")
			return 100
		},
	}

	v := pool.Get().(int)
	fmt.Println(v)
	pool.Put(3)
	runtime.GC() //GC 会清除sync.pool中缓存的对象
	v1, info := pool.Get().(int)
	fmt.Println(v1, info)
	//不做put操作，processor的私有对象会被get操作拿走，再次get会创建一个新对象
	//v2, info := pool.Get().(int)
	//fmt.Println(v2, info)
}

func TestSyncPoolInMultiGoroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create a new object")
			return 10
		},
	}

	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}

//适合于通过复用，降低复杂对象的创建和GC代价
//协程安全，会有锁的开销
//生命周期受GC影响，不适合做连接池等，需自己管理生命周期的资源的池化
