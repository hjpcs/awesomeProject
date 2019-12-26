package panic_recover

import (
	"errors"
	"fmt"
	"testing"
)

func TestPanicVxExit(t *testing.T) {
	//defer func() {
	//	fmt.Println("finally")
	//}()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recovered from", err)
			//log.Error("recovered panic", err)
		}
	}()
	fmt.Println("start")
	panic(errors.New("something wrong"))
	//os.Exit(-1)
}

//recover可能会导致僵尸进程，health check失效，不如let it crash，由daemon进程重启服务
