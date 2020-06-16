package my_map

import (
	"testing"
)

// var 变量名 map[key数据类型] value数据类型
func TestInitMap(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapForSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}
	mySet[3] = true
	t.Log(len(mySet))
	// delete(map名称,key)，指定key从map中删除
	delete(mySet, 1)
	n = 1
	if mySet[n] {
		t.Logf("%d is existing", n)
	} else {
		t.Logf("%d is not existing", n)
	}

}

func TestStrucMap(t *testing.T) {
	type personInfo struct {
		ID      string
		Name    string
		Address string
		Credit  int
	}
	score := map[string]float32{"C": 100, "C++": 98, "Python": 99, "Java": 100}
	infodb := map[string]personInfo{"zhangsan": personInfo{"20201010", "Tommy", "Shanghai Pudong Jinke Road", 192}, "lisi": personInfo{"20201011", "Jackie", "Shanghai Putuo", 111}}
	t.Log(score)
	t.Log(infodb)
	// value为map的value, ok则是返回值，是否有对应的值，ok为true，找到，否则没找到
	value, ok := infodb["zhangsan"]
	if ok {
		t.Log(value)
	} else {
		t.Log("Error")
	}
}

func TestMakeMap(t *testing.T) {
	// make创建一个map，格式make(map[key的类型] value的类型)
	score := make(map[string]int32, 10)
	score["Chinese"] = 100
	score["English"] = 99
	score["Math"] = 100
	t.Log(score)
}

/*
  map是无序的，每次打印出来的map都会不一样，它不能通过index获取，而必须通过key获取。
    map的长度是不固定的，也就是和slice一样，也是一种引用类型。
    内置的len函数同样适用于map，返回map拥有的key的数量。
    map的值可以很方便的修改，通过重新赋值即可。
*/
