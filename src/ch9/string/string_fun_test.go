package string

import (
	"strconv"
	"strings"
	"testing"
)

func TestStringFn(t *testing.T) {
	s := "A,B,C"
	t.Log(s)
	t.Log(len(s))
	parts := strings.Split(s, ",")
	t.Log(parts)
	t.Log(len(parts))
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
	t.Log(len(strings.Join(parts, "-")))

}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
