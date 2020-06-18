package string_test

import (
	"strconv"
	"strings"
	"testing"
)

// strings.Split用于对字符串按指定符号分隔，
// strings.Join用于对字符串按指定符号合并，func Join(t string, sep byte) string
// func Split(t string, set byte) string
func TestStringFn(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	for _, part := range parts {
		t.Log(part)
	}
	t.Log(strings.Join(parts, "-"))
}

func TestConv(t *testing.T) {
	s := strconv.Itoa(10)
	t.Log("str" + s)
	if i, err := strconv.Atoi("10"); err == nil {
		t.Log(10 + i)
	}
}
