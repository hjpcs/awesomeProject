package groutine

import (
	"fmt"
	"testing"
	"time"
)

func TestGroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		//go方法调用的时候都是值传递，在传递i的同时复制了一份，所以在每个协程里面i的地址是不一样的，没有竞争关系
		go func(i int) {
			fmt.Println(i)
		}(i)
		//输出10个10，因为i这个变量在test所在的协程以及启动的其他协程里面被共享了，存在竞争条件，一般来说这种情况需要用锁的机制来完成
		//go func() {
		//	fmt.Println(i)
		//}()
	}
	time.Sleep(time.Millisecond * 50)
}
