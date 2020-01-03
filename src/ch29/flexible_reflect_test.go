package ch29

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestDeepEqual(t *testing.T) {
	a := map[int]string{1: "one", 2: "two", 3: "three"}
	b := map[int]string{1: "one", 2: "two", 3: "three"}
	//t.Log(a==b)
	t.Log(reflect.DeepEqual(a, b))

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	s3 := []int{2, 3, 1}
	t.Log("s1 == s2?", reflect.DeepEqual(s1, s2))
	t.Log("s1 == s3?", reflect.DeepEqual(s1, s3))
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

func fillBySettings(st interface{}, settings map[string]interface{}) error {

	// func (v Value) Elem() Value
	// Elem returns the value that the interface v contains or that the pointer v points to.
	// It panics if v's Kind is not Interface or Ptr.
	// It returns the zero Value if v is nil.

	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		return errors.New("the first param should be a pointer to the struct type")
	}

	// Elem()获取指针指向的值，且这个方法只有指针才能调用
	if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
		return errors.New("the first param should be a pointer to the struct type")
	}

	if settings == nil {
		return errors.New("settings is nil")
	}

	var (
		field reflect.StructField
		ok    bool
	)

	for k, v := range settings {
		//遍历st和settings中有没有相同的key，如果没有，继续，直到找到有相同的key或者结束
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		fmt.Println("settings的值：", k, v)
		fmt.Println("匹配结果：", field, ok)
		//如果相同的key，type也相同
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)                 //返回指针
			vstr = vstr.Elem()                          //返回值
			vstr.FieldByName(k).Set(reflect.ValueOf(v)) //对st中的key，更新其在settings中的value
		}
	}
	//fmt.Println("==", reflect.ValueOf(st))                 //返回的是指针 &{ Mike 30}
	//fmt.Println("--", (reflect.ValueOf(st)).Elem())        //返回的是值value { Mike 30}
	//fmt.Println("**", (reflect.ValueOf(st)).Elem().Type()) //返回的是类型type ch29.Employee
	return nil
}

func TestFillNameAndAge(t *testing.T) {
	settings := map[string]interface{}{"Name": "Mike", "Age": 30, "QAQ": "QWQ"}
	e := Employee{}
	if err := fillBySettings(&e, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(e, &e) //{ Mike 30} &{ Mike 30}
	//t.Log(reflect.TypeOf(e), reflect.TypeOf(&e)) //ch29.Employee *ch29.Employee
	c := new(Customer)
	if err := fillBySettings(c, settings); err != nil {
		t.Fatal(err)
	}
	t.Log(c, &c, *c) //&{ Mike 30} 0xc00000e038 { Mike 30}
	//t.Log(reflect.TypeOf(c), reflect.TypeOf(&c), reflect.TypeOf(*c)) //*ch29.Customer **ch29.Customer ch29.Customer
}
