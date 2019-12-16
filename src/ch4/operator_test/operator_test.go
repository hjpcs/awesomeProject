package operator_test

import "testing"

const (
	Readable   = 1 << iota //0001 1
	Writable               //0010 2
	Executable             //0100 4
)

func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 3, 4, 5}
	//c := [...]int{1, 2, 3, 4, 5}
	d := [...]int{1, 2, 3, 4}
	t.Log(a == b)
	//t.Log(a == c)
	t.Log(a == d)
}

func TestBitClear(t *testing.T) {
	//&^:按位清零，右边为1，左边置0，右边为0，左边不变
	a := 7 //0111
	//a := 1 //0001
	t.Log(a, Readable, Writable, Executable)
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	a = a &^ Readable //0111	0001	0110
	t.Log(a) //6
	a = a &^ Executable //0110	0100	0010
	t.Log(a) //2
	t.Log(a&Readable == Readable, a&Writable == Writable, a&Executable == Executable)
	t.Log(a, Readable, Writable, Executable)
}
