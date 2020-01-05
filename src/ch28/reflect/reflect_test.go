package reflect

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

//reflect.TypeOf返回类型(reflect.Type)
//reflect.ValueOf返回值(reflect.Value)
//可以从reflect.Value获得类型
//通过kind来判断类型

func TestTypeAndValue(t *testing.T) {
	var f float64 = 10
	t.Log(reflect.TypeOf(f), reflect.ValueOf(f))
	t.Log(reflect.ValueOf(f).Type())
}

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Integer")
	default:
		fmt.Println("Unknown", t)
	}
}

func TestBasicType(t *testing.T) {
	var f float64 = 12
	CheckType(f)
	CheckType(&f)
}

type Employee struct {
	EmployeeID string
	Name       string `format:"normal"`
	Age        int
}

func (e *Employee) UpdateAge(newVal int) {
	e.Age = newVal
}

type Customer struct {
	CookieID string
	Name     string
	Age      int
}

//利用反射编写灵活的代码
//按名字访问结构的成员
//reflect.ValueOf(*e).FieldByName("Name")
//按名字访问结构的方法
//reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(1)})

func TestInvokeByName(t *testing.T) {
	e := &Employee{"1", "Mike", 30}
	t.Log(reflect.TypeOf(e), unsafe.Pointer(e), e)
	t.Log(reflect.TypeOf(*e))
	//按名字获取成员
	t.Log(reflect.ValueOf(*e))
	t.Log(reflect.ValueOf(*e).FieldByName("Name").Type())
	t.Log(reflect.TypeOf(*e).FieldByName("Name"))
	t.Logf("Name: value(%[1]v), type(%[1]T)", reflect.ValueOf(*e).FieldByName("Name"))
	if nameField, ok := reflect.TypeOf(*e).FieldByName("Name"); !ok {
		t.Error("failed to get 'Name' field")
	} else {
		t.Log("Tag:format", nameField.Tag.Get("format"))
	}
	reflect.ValueOf(e).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(5)})
	t.Log("Update Age:", e)
}
