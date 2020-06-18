package exception

import (
	"fmt"
	"os"
	"testing"
)

/*
defer
defer后边会接一个函数，但该函数不会立刻被执行，
而是等到包含它的程序返回时(包含它的函数执行了return语句、运行到函数结尾自动返回、对应的goroutine panic）defer函数才会被执行。
通常用于资源释放、打印日志、异常捕获等
*/

func Test_basic_exception(t *testing.T) {
	f, err := os.Open("/tmp/aa.txt")
	if err != nil {
		t.Log(err)
	}
	/**
	 * 这里defer要写在err判断的后边而不是os.Open后边
	 * 如果资源没有获取成功，就没有必要对资源执行释放操作
	 * 如果err不为nil而执行资源执行释放操作，有可能导致panic
	 */
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error opening file")
		} else {
			f.Close()
		}
	}()
}

// 如果有多个defer函数，调用顺序类似于栈，越后面的defer函数越先被执行(后进先出)
func Test_defer_order(t *testing.T) {
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	defer fmt.Println(4)
}

func Test_defer_call_1(t *testing.T) {
	result := 1
	defer func() {
		result++
	}()
	t.Log(result)
	return
}

func Test_defer_call_2(t *testing.T) {
	k := 5
	defer func() {
		k = k + 5
	}()
	t.Log(k)
}

func Test_defer_call_3(t *testing.T) {
	r := 9
	defer func(r int) {
		r += 5
	}(r)
	t.Log(r)
	return
}

func deer(i int) int {
	defer func() {
		i += 10
	}()
	return i
}

func Test_defer_call_4(t *testing.T) {
	t.Log(deer(10))
}

func cook() int {
	result := 1
	defer func() {
		result += 1
	}()
	return result
}

func Test_defer_5(t *testing.T) {
	t.Log(cook())
}

//defer, panic与recover
// defer如果再panic的后面，则无法执行
func Test_defer_panic_1(t *testing.T) {
	// defer的内容看不到，因为到了panic，就被视为出错，终止了
	//	panic("a")
	defer func() {
		fmt.Println("What do you do ")
	}()
}

//defer如果在panic的前面，则可以执行
func Test_defer_panic_2(t *testing.T) {
	// defer的内容可以看到，因为panic在后面
	defer func() {
		fmt.Println("Can you see me?")
	}()
	//	panic("Error")
}

// G是调用函数，F是被调用函数，如果F中出现panic时，F函数会立刻终止，不会执行F函数㕯panic后面的内容，但不会立刻return，而是调用F的defer
//如果F的defer中有recover捕获，则F在执行玩玩defer后正常返回，调用函数F的函数G继续正常运行
// d注意：defer定义的类似于析构函数，最后一定要执行的，无论出错与否。这里的最后包含两层意思： 1. 函数正常功能完成 2. 函数出现panic异常后
// panic报异常--recover捕获 recover()得到的就是panic报出的异常
func Test_defer_panic_3(t *testing.T) {
	G()
}

func G() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:3")
		}
	}()
	F()
	fmt.Println("Go on")
	fmt.Println("Why should I continue?")
}

func F() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("捕获异常", err)
		}
		fmt.Println("3")
	}()
	panic("3 do")
}

//如果F的defer中无recover捕获，则将panic抛到G中，G函数会立刻终止，不会执行G函数后面的内容，但不会立刻return，而是调用G的defer，依次类推
//也即相当于：如果内层函数没有异常捕获，则这个异常会被抛到外层函数，外城函数如果没有捕获，会继续抛向再外层，依次类推
func Test_painc_defer_4(t *testing.T) {
	G1()
}

func G1() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Raise Reception from G, outer function, get from recover")
		}
		fmt.Println("Normal function in defer")
	}()
	F1()
	fmt.Println("Continue after panic in outer function")
}

func F1() {
	defer func() {
		fmt.Println("defer function without recover()")
	}()
	panic("Error occurred in F1")
}

/*
defer function without recover()
Raise Reception from G, outer function, get from recover
Normal function in defer
*/

//如果一直没有recover，抛出的panic到当前goroutine最上层函数时，程序直接异常终止
func Test_panic_defer_5(t *testing.T) {
	G2()
}

func G2() {
	defer func() {
		fmt.Println("defer in outer function G2")
	}()
	F2()
	fmt.Println("Normal steps in outer function G2")
}

func F2() {
	defer func() {
		fmt.Println("defer in inner function F2")
	}()
	//	panic("Error Occured here when running F2")
	fmt.Println("Normal steps in inner function F2")
}

/*
发生错误的时候，先调内层的defer，如果没有recover，再调外层的defer，如果还没有recover，y依次类推，如果最顶层也没有，就报错，异常退出
defer in inner function F2
defer in outer function G2
--- FAIL: Test_panic_defer_5 (0.00s)
panic: Error Occured here when running F2 [recovered]
	panic: Error Occured here when running F2
*/

/* recover都是在当前的goroutine里进行捕获的，
这就是说，对于创建goroutine的外层函数，如果goroutine内部发生panic并且内部没有用recover，
外层函数是无法用recover来捕获的，这样会造成程序崩溃 */

func Test_defer_panic_6(t *testing.T) {
	G3()
}

func G3() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error captured", err)
		}
	}()
	// 产生了一个新的goroutine调用函数，则异常不会抛到这个goroutine中
	go F3()
	fmt.Println("Normal func")
}

func F3() {
	defer func() {
		fmt.Println("No recover func here")
	}()
	//	panic("Error F3")
	fmt.Println("F3 function")
}

/*
No recover func here
panic: Error F3
*/

/*
recover返回的是interface{}类型而不是go中的 error 类型，如果外层函数需要调用err.Error()，会编译错误，也可能会在执行时panic
*/

func Test_error_type(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			//			fmt.Println("Error Capturred", err.Error())
			fmt.Println("Error captured", err)
		}
	}()
	panic("error")
}

func Test_defer_panic_recover(t *testing.T) {
	fmt.Println("It's finished")
}
