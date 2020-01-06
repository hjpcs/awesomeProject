package unsafe_programming

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
	"unsafe"
)

func TestUnsafe(t *testing.T) {
	i := 10
	f := *(*float64)(unsafe.Pointer(&i))
	ff := *(*int)(unsafe.Pointer(&i))
	t.Log(unsafe.Pointer(&i))
	t.Log(f, ff)
}

//the cases is suitable for unsafe
type MyInt int

//合理的类型转换
func TestConvert(t *testing.T) {
	a := []int{1, 2, 3, 4}
	b := *(*[]MyInt)(unsafe.Pointer(&a))
	t.Log(b)
}

//原子类型操作
func TestAtomic(t *testing.T) {
	var shareBufPtr unsafe.Pointer
	//往slice里写20个数
	writeDataFn := func() {
		var data []int
		for i := 0; i < 20; i++ {
			data = append(data, i)
		}
		//写完后用一个原子操作指到共享指针上
		atomic.StorePointer(&shareBufPtr, unsafe.Pointer(&data))
	}
	//加载共享指针的内容并打印
	readDataFn := func(i int) {
		data := atomic.LoadPointer(&shareBufPtr)
		fmt.Println("协程", i, data, *(*[]int)(data), &shareBufPtr)
	}
	var wg sync.WaitGroup
	writeDataFn()
	t.Log(shareBufPtr, &shareBufPtr)
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go func() {
			for i := 0; i < 3; i++ {
				writeDataFn()
				time.Sleep(time.Millisecond * 100)
			}
			wg.Done()
		}()
		wg.Add(1)
		go func() {
			for i := 0; i < 3; i++ {
				readDataFn(i)
				time.Sleep(time.Millisecond * 100)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
