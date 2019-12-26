package _interface

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() Code {
	return "go hello world!"
}

type JavaProgrammer struct {
}

func (j *JavaProgrammer) WriteHelloWorld() Code {
	return "java hello world!"
}

func TestClient(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
	p = new(JavaProgrammer)
	t.Log(p.WriteHelloWorld())

	//interface需要是指针类型
	//g := new(GoProgrammer)
	g := &GoProgrammer{}
	j := new(JavaProgrammer)
	writeFirstProgram(g)
	writeFirstProgram(j)
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v\n", p, p.WriteHelloWorld())
}

//duck type
//接口为非入侵性，实现不依赖于接口定义
//所以接口的定义可以包含在接口使用者包内
