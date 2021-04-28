package go_routine

import (
	"fmt"
	"runtime"
)

func FuncGosched() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")
	// 主协程
	for i := 0; i < 2; i++ {
		// 让出CPU时间片，等待重新安排任务
		// 优先完成协程
		runtime.Gosched()
		fmt.Println("hello")
	}
}

func FuncGoexit() {
	go func() {
		defer fmt.Println("A.defer") // 2
		func() {
			defer fmt.Println("B.defer") // 1
			// 结束协程,不执行后面的任何代码
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for {
	}
}

func couter(out chan <- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan <- int,in <- chan int) {
	// 使用 range 接收channel可以自动处理channel的关闭，并退出range
	// 或是接收通道的 flag 值: i,ok := channel,判断ok即可知道通道是否关闭
	for i := range in {
		out <- i * i
	}
	close(out)
}

func printer(in <- chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

//1.对一个关闭的通道再发送值就会导致panic。
//2.对一个关闭的通道进行接收会一直获取值直到通道为空。
//3.对一个关闭的并且没有值的通道执行接收操作会得到对应类型的零值。
//4.关闭一个已经关闭的通道会导致panic。
func SingleArrowChannel()  {

	// 没有缓存区的通道在发送时必须有确定的接收者，该通道亦被称为同步通道
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 单向通道
	go couter(ch1)
	go squarer(ch2,ch1)
	printer(ch2)
}