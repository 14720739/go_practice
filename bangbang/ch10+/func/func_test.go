package method_func

import (
	"fmt"
	"testing"
)

/*
func function_name([parameter list]) [ return_types } {
函数体
}

参数： 无参，一个参数，多个参数，可变参数，均可以
返回值： 无返回值，一个返回值，多个返回值均可
*/
func no_parameter_func() {
	fmt.Println("Okay")
}

func Test_no_parameter_func(t *testing.T) {
	no_parameter_func()
}

// 多参数
func max(a1, a2 int) int {
	if a1 >= a2 {
		return a1
	} else {
		return a2
	}

}

func Test_multiple_parameter(t *testing.T) {
	t.Log(max(3, 5))
}

// go语言中，可以直接指定函数返回值的变量名，一旦指定，则无需再具体指定return的变量名
func rectProperties(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = (length + width) * 2
	return
}

func Test_rectProperties(t *testing.T) {
	area, peri := rectProperties(10, 5)
	t.Logf("Area is :%f, Perimeter is:%f", area, peri)
}

//空白符： _在Go中被用作空白符，可以用作表示任何类型的任何值

func Test_blankCh(t *testing.T) {
	// _接收的周长返回值，将被丢弃
	area, _ := rectProperties(22, 12)
	t.Logf("area=%f", area)
}

//可变参数，如果函数的最后一个参数被记作...T，这是函数可以接收任意个T类型的参数作为最后一个参数
//这个功能类似于python中的， *args, **kwargs
/*
func find(num int, nums ...int)

func 函数名(固定参数列表, v … T)(返回参数列表){
    函数体
}
    可变参数一般被放置在函数列表的末尾，前面是固定参数列表，当没有固定参数时，所有变量就将是可变参数。
    v 为可变参数变量，类型为 []T，也就是拥有多个 T 元素的 T 类型切片，v 和 T 之间由...即3个点组成。
    T 为可变参数的类型，当 T 为 interface{} 时，传入的可以是任意类型。
————————————————
版权声明：本文为CSDN博主「huycheaven」的原创文章，遵循CC 4.0 BY-SA版权协议，转载请附上原文出处链接及本声明。
原文链接：https://blog.csdn.net/huyuchaoheaven/java/article/details/89888478
*/

func find(num int, nums ...int) {
	fmt.Printf("type of nums is %T \n", nums)
	found := false
	for i, v := range nums {
		for v == num {
			fmt.Println(num, "found at index", i, "in", nums)
			found = true
			break
		}
	}
	if !found {
		fmt.Println(num, "not found in", nums)
	}
	fmt.Printf("\n")
	return
}

func Test_find(t *testing.T) {
	find(100, 1, 3, 5, 7, 8, 11, 98, 65, 100, 109, 67)
	nums := []int{101, 100, 109}
	//参数nums...，代表着切片nums中的所有值，一个证书系列
	find(100, nums...)
	//find(100,nums),这种写法是错误的，因为nums代表的是[]int，是切片的值，则会报错：cannot use nums (type []int) as type int in argument to find
}

//切片参数：默认为引用传递,可以认为有...就是指针，例如welcome...，就是一个指向切片welcome的指针
func changes(s ...string) {
	s[0] = "Go"
	s = append(s, "playground")
	fmt.Println(s)
}

func Test_changes(t *testing.T) {
	welcome := []string{"hello", "world"}
	t.Log("welcome...", welcome)
	changes(welcome...)
	fmt.Println(welcome)
}

//方法，方法其实就是一个函数，在func这个关键字和方法名中间键入了一个特殊的接收器类型。接收器可以是结构体类型，或者是非结构体类型。接收器可以在方法的内部访问。
type Student struct {
	Name   string
	Mobile string
	age    int
	Sex    int
}

// (s Student)是个接收器，定义了函数可以接收类型为Student的结构体变量，并使用其中的值
// 接收器位于func关键字以及函数名之间
// 引用的时候，使用类似： 变量.方法名 ， 变量需要是接收器指定类型的变量
func (s Student) ToString() {
	fmt.Printf("Name is %s, sex is %d \n", s.Name, s.Sex)
}

func ToString(s Student) {
	fmt.Printf("Name is %s sex is %d \n", s.Name, s.Sex)
}
func Test_method_as_para(t *testing.T) {
	var student = Student{
		Name:   "李",
		Mobile: "13800138000",
		Sex:    1,
	}
	student.ToString()
	ToString(student)
}

// 相同的接收器和相同的方法名只能定义一个
// 不同的接收器可以定义在相同的方法名上,例如：
type Rectangle struct {
	length int
	width  int
}

type Circle struct {
	radius float64
}

type Triangle struct {
	base   float64
	height float64
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (c Circle) Area() float64 {
	return 3.14 * c.radius * c.radius
}

func (t Triangle) Area() float64 {
	return t.base * t.height / 2
}

func Area(r Rectangle) int {
	return r.length * r.width
}
func Test_RCT(t *testing.T) {
	var rectangle1 = Rectangle{
		length: 20,
		width:  15,
	}
	var circle1 = Circle{
		radius: 21,
	}
	var triangle1 = Triangle{
		base:   11.2,
		height: 23.1,
	}
	t.Logf("Circle1=%f", circle1.Area())
	t.Logf("rectangle1=%d", rectangle1.Area())
	t.Logf("triangle1=%f", triangle1.Area())
}

// 指针接收器与值接收器
func (s Student) ChnageName(name string) {
	s.Name = name
	fmt.Printf("Inner value:%s\n", s.Name)
}

func (s *Student) ChangeName2(name string) {
	s.Name = name
}

func Test_changename(t *testing.T) {
	var student = Student{
		Name:   "Alex",
		Mobile: "13876789890",
		age:    36,
		Sex:    1,
	}
	student.ChnageName("People")
	fmt.Println(student.Name)
	(&student).ChangeName2("Joan")
	fmt.Println(student.Name)
	//调用ChangeNam2的时候，student会自动转为&student，结果相同
	student.ChangeName2("Thomas")
	fmt.Println(student.Name)
}

// 什么时候用指针接收器，什么时候使用值接收器？
// 改变调用者的字段值就用指针接收器，否则用值接收器

// 在方法中使用值接收器，与在函数中使用值参数
// 当一个函数有一个值参数，它只能接收一个值参数
// 当一个方法有一个值接收器，它可以接收值接收器和指针接收器
func Test_func_call(t *testing.T) {
	r := Rectangle{
		length: 100,
		width:  50,
	}
	t.Log(Area(r))
	t.Log(r.Area())

	p := &r
	// Go语言把 p.area() 解释为 (*p).area()
	t.Log(p.Area())
}

//在方法中使用指针接收器与在函数中使用指针参数
func perimeter(r *Rectangle) {
	fmt.Println("perrimter funciton output:", 2*(r.length+r.width))
}

func (r *Rectangle) perimeter() {
	fmt.Println("perimeter method output:", 2*(r.width+r.width))
}

func Test_permeter(t *testing.T) {
	r := Rectangle{
		length: 100,
		width:  30,
	}
	p := &r
	perimeter(p)
	p.perimeter()
	//perimeter(r) 错误，参数必须为指针
	r.perimeter() //始终值来调用指针接收器
}

// 非结构体上的方法
//  r如果在一个类型上定义一个方法，则方法的接收器类型定义与这个方法的定义应该在同一个包中
// 下面就是一个错误的定义，因为类型int的定义与方法的定义不在一个包中
/*
func (a int) add(b int) {

}
*/
type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

func Test_self_define(t *testing.T) {
	num1 := myInt(10)
	num2 := myInt(20)
	sum := num1.add(num2)
	t.Log(sum)
}
