package customer_type

import (
	"fmt"
	"testing"
	"time"
)

type IntConv func(op int) int

// （inner IntConv）是输入参数的类型，以函数为输入参数类型；IntConv是返回值的类型，以函数作为返回值的类型

func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func TestFn(t *testing.T) {
	tsSF := timeSpent(slowFun)
	t.Log(tsSF(10))
}
