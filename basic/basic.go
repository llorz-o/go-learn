package basic

import (
	"errors"
	"fmt"
	"reflect"
)

func add(x int, args ...int) int {
	var re int
	for v := range args {
		re += v
	}
	return re + x
}

// 不定类型与参数
func notConfirmArgs(args ...interface{}) int {
	var re int
	switch args[0] {
	case reflect.Int:
		re = 1
	default:
		re = 0
	}
	return re
}

// 显式声明返回参数
func t1(x int) (y int) {
	// 当前层级不可声明 y，将复用返回值声明
	y = x
	// 第二层级可以显式声明，也必须显式返回
	//{
	//	var y = 1
	//	return  y
	//}
	return // 将会隐式返回
}

func deferFn(x, y int) (z int) {
	// defer在return前执行
	// 多个defer前进后出
	defer func() {
		z += 100
		println("defer func running", z)
	}() // defer 调用函数时可在结尾转入参数，保证参数被copy
	z = x + y
	println("x + y = ", z)
	return z + 300 // x + y => defer z += 100 => z = z + 300,最终return时会将计算结果赋值给 z
}

func testPanic() {

	//1.利用recover处理panic指令，defer 必须放在 panic 之前定义，
	//  另外 recover 只有在 defer 调用的函数中才有效。否则当panic时，
	//  recover无法捕获到panic，无法防止panic扩散。
	//2.recover 处理异常后，逻辑并不会恢复到 panic 那个点去，函数跑到 defer 之后的那个点。
	//3.多个 defer 会形成 defer 栈，后定义的 defer 语句会被最先调用。

	defer func() {
		println("after recover panic")
	}()

	defer func() {
		if err := recover(); err != nil {
			println(err.(string))
		}
	}()

	panic("test panic")

	println("after panic")
}

func testConstructError() {
	var errZero = errors.New("zero err")

	var castErr = func() (int, error) {
		return 0, errZero
	}

	defer func() {
		fmt.Println(recover())
	}()

	switch v, err := castErr(); err {
	case nil:
		println(v)
	case errZero:
		panic(err)

	}

}

func emulationTry(f func(), catch func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			catch(err)
		}
	}()

	f()

}

func FunctionAndPanic() {
	Divide("函数", func() {
		println(deferFn(1, 2))

		// 匿名函数可在函数体内声明
		var x = 1
		var anonymousFunc = func(x int) {
			println("anonymous func:", x)
		}
		anonymousFunc(x)
	})

	Divide("异常处理", func() {
		testPanic()

		testConstructError()

		emulationTry(func() {
			panic("cast error emulation try catch")
		}, func(err interface{}) {
			fmt.Printf("catch err:%v \n", err)
		})
	})
}
