package client

import (
	"ch15/series"
	"testing"
)

func TestPackage(t *testing.T) {
	t.Log(series.GetFibonacciSeries(5))
	t.Log(series.Square(5))
}

//方法名必须是大写开头，否则无法引用
