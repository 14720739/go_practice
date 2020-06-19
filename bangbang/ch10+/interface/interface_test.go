package go_interface

import "testing"

/*
接口声明了方法集合

如果某个结构体实现了接口内的所有方法, 就可以把该结构体赋值给接口

接口可以帮助我们实现类似面向对象的"类的继承"
*/

// Abser接口定义了方法集合
type Abser interface {
	Abs() float64
}

// Vertex - 结构体
type Vertex struct {
	X, Y float64
}

// Abs -
func (v *Vertex) Abs() float64 {
	return v.X*v.X + v.Y*v.Y
}

func Test_interface(t *testing.T) {
	var a Abser
	//
	a = &Vertex{3, 5}
	//	b = Vertex{3,4}
	t.Log(a.Abs())
}
