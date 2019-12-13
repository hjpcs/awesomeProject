package main //包，表明代码所在的模块（包）

import (
	"fmt"
	"os"
) //引入代码依赖

//功能实现
func main() {
	if len(os.Args) > 1 {
		fmt.Println(len(os.Args))
		fmt.Println("Hello World", os.Args[1])
	}
}

//应用程序入口
//1.必须是main包：package main
//2.必须是main方法：func main()
//3.文件名不一定是main.go

//退出返回值
//Go中main函数不支持任何返回值
//通过os.Exit来返回状态

//获取命令行参数
//main函数不支持传入参数
//在程序中直接通过os.Args获取命令行参数