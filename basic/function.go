package basic

import "fmt"

type Test struct {
	name string
}

func (t Test) m1() {
	fmt.Printf("type T name:%v t poiter:%p \n", t.name, &t)
}

func (t *Test) m2() {
	t.name = "reset name"
	fmt.Printf("type *T name:%v t poiter:%p \n", t.name, t)
}

type Ter interface {
	m2()
}

func FunctionExp() {
	Divide("方法集", func() {

		t := Test{"Jojo"}
		t.m1()
		t.m2()

		t2 := &t
		t2.m1()
		t2.m2()

		m2 := t.m2
		M2 := (*Test).m2

		m2()   // 隐式传递 receiver，当方法当receiver为非指针类型，receiver将会被复制
		M2(&t) // 显式传递 receiver

		var x Ter
		//var ta = Test{}
		//x = ta // interface Ter 的 x 不能接收 值类型的 Test 实例
		var tb = &Test{}
		x = tb
		fmt.Println(x)

	})

	Divide("类型断言", func() {
		var x interface{}
		x = "12"
		if v, ok := x.(string); ok {
			println("断言类型为string成功", v)
		}

		switch v := x.(type) {
		case int:
			println("match int",v)
		default:
			fmt.Printf("not match %v \n",v)
		}
	})
}
