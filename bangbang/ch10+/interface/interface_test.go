package go_interface

import (
	"testing"
)

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
	var b = Vertex{
		X: 3,
		Y: 8,
	}
	t.Log((&b).Abs())
	//由于 *Vertex 实现了Abs(), 所以他实现了接口a的所有方法, 可以赋值给接口a
	a = &Vertex{3, 5}
	//c = Vertex{3,4} // 运行错误, Vertex 上没有实现Abs()
	t.Log(a.Abs())
}

// Reader 定义一个接口
type Reader interface {
	Read(b []byte) (n int, err error)
}

type Writer interface {
	Write(b []byte) (n int, err error)
}

//隐式接口，相当于合并了其他接口的定义, Reader 和Writer接口
type ReadWrite interface {
	Reader
	Writer
}

func Test_imp_interface(t *testing.T) {
	//	var w ReadWrite
	var abc []byte
	abc = []byte("abc")
	abc = []byte("efg")
	abc = append(abc, 'h')
	abc[1] = 'o'
	t.Logf("length of abc is %d", len(abc))

	//  w = os.Stdout
	//  t.Log(w)
	// log.Println(w.Read(abc))
}
