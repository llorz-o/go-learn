package basic

import "fmt"

type Animal struct {
	name string
	age  int
}

// Myint 不仅仅是结构体，任何类型都可以拥有方法,但仅支持对本地包修改方法
type Myint int

func (mi Myint) say() {
	println("added function say")
}

// Say 结构体的方法
func (a Animal) Say() {
	fmt.Printf("我是动物的的说法方法\n")
}

// ChangeName 设置结构体的名字
func (a Animal) ChangeName(name string) {
	a.name = name
}

// SetAge 设置年龄
func (a *Animal) SetAge(age int) {
	a.age = age
}

func Base() {
	Divide("位操作", func() {
		println("1 << n = 1*(2^n)", 1<<10, 2*2*2*2*2*2*2*2*2*2)
		println("1 >> n = 1/(2^n)", 1/4)
	})

	Divide("多行字符串", func() {
		pt := `
http://www.baidu.com
`
		println(pt)
	})

	Divide("遍历字符串", func() {
		// 遍历字符串
		s := "pprof.cn博客"
		for i := 0; i < len(s); i++ { //byte，代表一个 ASCII字符
			fmt.Printf("%v(%c) ", s[i], s[i])
		}
		fmt.Println()
		// var r rune = 'r'
		for _, r := range s { //rune ,代表一个 utf-8字符
			fmt.Printf("%v(%c) ", r, r)
		}
		fmt.Println()
	})

	Divide("修改字符串 与类型转换 T()", func() {
		// 现转换为 rune 或是 byte
		s := "pprof.cn博客"
		runeS1 := []rune(s)
		runeS1[9] = '人'
		println(string(runeS1))
	})

	Divide("数组", func() {
		// 指针数组 [n]*T   数组指针 *[n]T
		var arr0 = [5]int{1, 2, 3} // 未初始化元素值为 0。
		var arr1 = [5]int{1, 2, 3, 4, 5}
		var arr2 = [...]int{1, 2, 3, 4, 5, 6}  // 通过初始化值确定数组长度。
		var arr3 = [5]string{3: "hhe", 4: "4"} // 使用引号初始化元素。
		var arr4 = [...]struct {
			name string
			age  uint8
		}{
			{"jojo", 12},
			{"Ali", 34},
		}

		fmt.Println(arr0)
		fmt.Println(arr1)
		fmt.Println(arr2)
		fmt.Println(arr3)
		fmt.Println(arr4)
	})

	Divide("多维数组", func() {
		// 第 2 纬度不能用 "..."。
		var arr0 = [...][3]int{
			{1, 2, 3},
			{3, 4, 5},
		}
		fmt.Println(arr0)

		println("找出数组【1，3，5，7，8】两个元素只和等于8的下标")
		var a = [...]int8{1, 2, 5, 7, 8, 6}

		for i := 0; i < len(a); i++ {
			for j := i + 1; j < len(a); j++ {
				b := a[i]
				c := a[j]
				if b+c == 8 {
					println("found re:", b, c, "the index:", i, j)
				}
			}
		}
	})

	Divide("切片", func() {
		var slice0 = []int{}
		var slice1 = make([]int, 0)
		var slice2 = make([]int, 0, 0)
		var arr = [...]int{1, 2, 3, 4, 5, 6}
		// [start:end] [start:] [:end] [:]
		var slice3 = arr[1:3]

		fmt.Println("[]int{} :", slice0)
		fmt.Println("make([]int, 0) :", slice1)
		fmt.Println("make([]int, 0, 0) :", slice2)
		fmt.Println("[...]int{1, 2, 3, 4, 5, 6} arr[1:3] :", slice3)

		fmt.Println("多维切片：", [][]int{
			{1, 2, 3},
			{100, 200},
			{1, 22, 33, 44},
		})

		// slice 超出容量后会自动扩容，并且会重新分配底层数组
		// 通常扩容是以之前的两，当容量大于1024时会增长之前容量的1/4

		data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

		s1 := data[:5]
		s2 := data[8:]

		fmt.Println("copy 函数将以长度较小的为基准", copy(s1, s2), s1, s2)
		fmt.Println("copy 函数将以长度较小的为基准", copy(s2, []int{0, 1, 2}), s2)

		s3 := data[2:6:10]

		println(fmt.Printf("字面量切片[start:end:cap] -> %v len: %d cap: %d", s3, len(s3), cap(s3)))

		var sli []int
		fmt.Println("nil切片：", sli)   // nil 切片是一个未申请内存的切片类型
		fmt.Println("空切片：", []int{}) // 空切片是一个有内存但是没有数据的切片

		// 切片扩容一定程度上会影响原切片，因为底层使用同一个数组

		var arr1 = [6]int{0, 1, 2, 3, 4, 5}
		sli1 := arr1[0:2]
		println(fmt.Printf("Pointer:%p slice:%v arr:%v", &sli1, sli1, arr))
		newSlice := append(sli1, 50)
		println(fmt.Printf("Pointer:%p newSlice:%v arr:%v", &newSlice, newSlice, arr))
		newSlice[1] += 10 // 影响了旧切片,实际使用中应当覆盖原切片
		newSlice[2] += 10
		println(fmt.Printf("Pointer:%p newSlice:%v arr:%v", &newSlice, newSlice, arr))
		println(fmt.Printf("Pointer:%p slice:%v arr:%v", &sli1, sli1, arr))

	})

	Divide("map", func() {

		kv := make(map[string]int)

		kv["jojo"] = 1

		v, ok := kv["jojo"]

		fmt.Printf("jojo value: %v is ok: %v", v, ok)

		delete(kv, "jojo")

	})

	Divide("结构体与类型", func() {

		type Myint int       // Myint 将会编译为实际的类型
		type TypeAlias = int // TypeAlias 只会存在于代码中，编译时会变成int型

		type Person struct { // 结构体是用来描述一组值的聚合型数据类型
			name string
			age  int
			city string
		}

		jojo := Person{
			age:  25,
			name: "jojo",
			city: "Manila",
		}

		fmt.Printf("结构体 person：%#v\n", jojo)

		var user struct {
			Name string
			Age  int
		}

		user.Name = "Jo"

		fmt.Printf("匿名结构体：%#v\n", user)

		// Or var user = new(person)
		// user.name = 'name'

		// Or var user = &person{}
		// 取结构体的地址相当于将该结构体实例化
		// 结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

		var a = Animal{}
		a.name = "Dog"
		a.age = 12
		a.Say()
		fmt.Printf("Animal a %#v \n", a)

		// 非指针方法将不会修改到实例，而是会创建并修改实例的副本，会消耗一定的内存与性能
		a.ChangeName("Cat")
		fmt.Printf("Animal a %#v \n", a)

		a.SetAge(24)
		fmt.Printf("Animal a %#v \n", a)

		type Address struct {
			Province string
			City     string
		}

		type User struct {
			Name string
			Address
			*Animal
		}

		u := User{
			Name:   "zhou",
			Animal: &Animal{},
		}
		u.Address.Province = "黑龙江"
		// 当访问结构体成员时会先在结构体中查找该字段，找不到再去匿名结构体中查找。
		// 当出现同名字段时，需要指定当前的匿名结构体
		u.City = "哈尔滨"

		fmt.Printf("匿名嵌套结构体：%#v\n", u)

		u.SetAge(36) // 通过匿名结构体，可以实现面向对象的继承

	})
}
