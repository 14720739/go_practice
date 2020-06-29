package go_interface

import (
	"fmt"
	"strconv"
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

/*
interface的变量可以持有任意实现该interface类型的对象，这给我们编写函数(包括method)提供了一些额外的思考，我们是不是可以通过定义interface参数，让函数接受各种类型的参数。
*/

/*
举个例子：fmt.Println是我们常用的一个函数，但是你是否注意到它可以接受任意类型的数据。打开fmt的源码文件，你会看到这样一个定义:
复制代码 代码如下:

type Stringer interface {
     String() string
}
*/
type Human1 struct {
	name  string
	age   int
	phone string
}

//通过这个方法Human实现了fmt.Stringer
func (h Human1) String() string {
	return "(" + h.name + "-" + strconv.Itoa(h.age) + "years-" + h.phone + ")"
}

func Test_interface_as_parameter(t *testing.T) {
	Bob := Human1{"Bob", 39, "13709081091"}
	fmt.Println("This human being is", Bob)
}

/*
再回顾一下前面的Box示例，你会发现Color结构也定义了一个method：String。其实这也是实现了fmt.Stringer这个interface，即如果需要某个类型能被fmt包以特殊的格式输出，你就必须实现Stringer这个接口。如果没有实现这个接口，fmt将以默认的方式输出。
//实现同样的功能
fmt.Println("The biggest one is", boxes.BiggestsColor().String())
fmt.Println("The biggest one is", boxes.BiggestsColor())
*/

/*
interface变量存储的类型
我们知道interface的变量里面可以存储任意类型的数值(该类型实现了interface)。那么我们怎么反向知道这个变量里面实际保存了的是哪个类型的对象呢？目前常用的有两种方法：
Comma-ok断言
Go语言里面有一个语法，可以直接判断是否是该类型的变量： value, ok = element.(T)，这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型。
如果element里面确实存储了T类型的数值，那么ok返回true，否则返回false。
让我们通过一个例子来更加深入的理解。

*/
type Element interface{}
type List []Element
type Person struct {
	name string
	age  int
}

//定义了String方法，实现了fmt.Stringer

func (p Person) String() string {
	return "(name:" + p.name + "-age;" + strconv.Itoa(p.age) + "years)"
}

func Test_interface_element(t *testing.T) {
	list := make(List, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Person{"Dennis", 70}
	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("List[%d] is int,its value is %d\n", index, value)
		case string:
			fmt.Printf("List[%d] is string ,its value is %s\n", index, value)
		case Person:
			fmt.Printf("List[%d] is string, its value is %s\n", index, value)
		default:
			fmt.Printf("list[%d] is of a different type\n", index)
		}
	}
}

// 这里有一点需要强调的是：element.(type)语法不能在switch外的任何逻辑里面使用，如果你要在switch外面判断一个类型就使用comma-ok。

//嵌入interface

/*
嵌入interface
Go里面真正吸引人的是它内置的逻辑语法，就像我们在学习Struct时学习的匿名字段，多么的优雅啊，那么相同的逻辑引入到interface里面，那不是更加完美了。如果一个interface1作为interface2的一个嵌入字段，那么interface2隐式的包含了interface1里面的method。
我们可以看到源码包container/heap里面有这样的一个定义

复制代码 代码如下:

type Interface interface {
    sort.Interface //嵌入字段sort.Interface
    Push(x interface{}) //a Push method to push elements into the heap
    Pop() interface{} //a Pop elements that pops elements from the heap
}

我们看到sort.Interface其实就是嵌入字段，把sort.Interface的所有method给隐式的包含进来了。也就是下面三个方法：

复制代码 代码如下:

type Interface interface {
    // Len is the number of elements in the collection.
    Len() int
    // Less returns whether the element with index i should sort
    // before the element with index j.
    Less(i, j int) bool
    // Swap swaps the elements with indexes i and j.
    Swap(i, j int)
}

另一个例子就是io包下面的 io.ReadWriter ，它包含了io包下面的Reader和Writer两个interface：

复制代码 代码如下:

// io.ReadWriter
type ReadWriter interface {
    Reader
    Writer
}

*/
