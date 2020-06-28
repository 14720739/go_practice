package go_interface

import (
	"fmt"
	"testing"
)

// interface Go语言里面设计最精妙的应该算interface，它让面对对象，内容组织实现非常的方便。简单地说，interface是一组method的组合，我们通过interface来定义对象的一组行为。

// interface类型。 interface类型定义了一组方法，如果某个对象实现了某个接口的所有方法，则此对象就实现了此接口。

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  //匿名字段Human
	school string
	loan   float32
}

type Employee struct {
	Human   //匿名字段Human
	company string
	money   float32
}

// Human对象实现Sayhi方法
func (h *Human) SayHi() {
	fmt.Println("Hello, Everybody")
}

// Human对象实现Sing方法
func (h *Human) Sing(lyrics string) {
	fmt.Println("La la, la la, la la la la...", lyrics)
}

// Human对象实现Guzzle方法
func (h *Human) Guzzle(beerStein string) {
	fmt.Println("Guzzle Guzzle Guzzle...", beerStein)
}

//Employee重载Human的Sayhi方法
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. call me on %s \n", e.name, e.company, e.phone)
}

//Student实现BorrowMoney方法
func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount //(again and again and ...)
}

// Employee 实现SpendSalary方法
func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount
}

func (singer *Employee) Sing(lyrics string) {
	fmt.Printf("Dou dou dou dou dou....,%s", lyrics)
}

// 定义interface
type Men interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type YoungChap interface {
	SayHi()
	Sing(lyrics string)
	Guzzle(beerStein string)
}

type ElderLyGent interface {
	SayHi()
	Sing(song string)
	SpendSalary(amount float32)
}

/*
interface可以被任意的对象实现。上面的Men interface被Human、Student和Employee实现。同理，一个对象可以实现任意多个interface，例如上面的Student实现了Men和YoungChap两个interface。
最后，任意的类型都实现了空interface(我们这样定义：interface{})，也就是包含0个method的interface。
*/

func Test_interface_1(t *testing.T) {
	mike := Student{Human{"Mike", 25, "111-111-111"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 22, "135-9089-1902"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "111-222-333-111"}, "Golang Inc", 101000}
	Tom := Employee{Human{"Tom", 29, "119-0918-7231"}, "Intel Co. Ltd", 9081}

	var i Men
	i = &mike
	i.SayHi()
	i = &paul
	i.SayHi()
	i = &sam
	i.SayHi()
	i = &Tom
	i.SayHi()
	/*
	   运行结果分别是：
	      Hello, Everybody
	      Hello, Everybody
	      Hi, I am Sam, I work at Golang Inc. call me on 111-222-333-111
	      Hi, I am Tom, I work at Intel Co. Ltd. call me on 119-0918-7231
	   这里表现了多态性，同一个接口变量，赋给不同类型的变量，表现出不同的特征。赋给employ类型的变量，则SayHi为重写的Emploee类型的动作
	*/
	//定义一个接口Men类型的slice，容量为3

	x := make([]Men, 3)
	fmt.Println(len(x), " ", cap(x))
	x[0], x[1], x[2] = &mike, &paul, &sam
	for _, value := range x {
		value.SayHi()
	}
	/* Interface 实际上就是一组抽象方法的集合，它必须由其他非interface类型实现，而不能自我实现。*/
}
