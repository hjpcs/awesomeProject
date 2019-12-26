package _struct

import (
	"fmt"
	"testing"
	"unsafe"
)

type Employee struct {
	Id   string
	Name string
	Age  int
}

func TestCreateEmployeeObj(t *testing.T) {
	e := Employee{
		Id:   "0",
		Name: "Bob",
		Age:  20,
	}
	e1 := Employee{Name: "Mike", Age: 30}
	e2 := new(Employee) //返回指针
	e2.Id = "2"
	e2.Age = 22
	e2.Name = "Rose"
	t.Log(e)
	t.Log(e1)
	t.Log(e1.Id)
	t.Log(e2)
	t.Logf("e is %T", e)
	t.Logf("e is %T", &e)
	t.Logf("e2 is %T", e2)
}

// 推荐使用这种写法，不会产生内存复制
func (e *Employee) String() string {
	fmt.Printf("address is %x\n",unsafe.Pointer(&e.Name))
	return fmt.Sprintf("Id:%s / Name:%s / Age:%d", e.Id, e.Name, e.Age)
}

// 这种写法会产生内存复制
//func (e Employee) String() string {
//	fmt.Printf("address is %x\n",unsafe.Pointer(&e.Name))
//	return fmt.Sprintf("Id:%s - Name:%s - Age:%d", e.Id, e.Name, e.Age)
//}

func TestStructOperations(t *testing.T) {
	e := Employee{"0", "Bob", 20} //返回实例不会默认带上修改后的String方法
	fmt.Printf("e address is %x\n",unsafe.Pointer(&e.Name))
	ee := new(Employee) //返回指针会默认带上修改后的String方法
	ee.Id= "0"
	ee.Name="Jack"
	ee.Age=20
	fmt.Printf("ee address is %x\n",unsafe.Pointer(&ee.Name))
	t.Log(e)
	t.Log(e.String())
	t.Log(ee)
	t.Log(ee.String())
}
