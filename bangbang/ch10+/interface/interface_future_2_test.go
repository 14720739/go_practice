package go_interface

import (
	"fmt"
	"testing"
)

// 空interface
// 空interface的定义为interface{},不包含任何方法。因此，所有的类型都实现了空interface。
// 空interface对于描述起不到任何作用（因为它不包含任何的Method），但是空interface在我们需要存储任意类型的数值的时候相当有用，因为它可以存储任意类型的数值，有点类似于C语言的void*类型

func Test_empty_interface(t *testing.T) {
	var a interface{}
	var b int = 5
	c := "Hello Kitty"
	//b为整型数值，c为字符串，但是都可以赋给空接口a，去实现空接口a来
	a = b
	fmt.Println("a=", a)
	a = c
	fmt.Println("a=", a)

}
