package go_routine

func FuncTimer() {
	//1。 timer基本使用
	//timer1 := time.NewTimer(2 * time.Second)
	//t1 := time.Now()
	//fmt.Printf("t1:%v \n", t1)
	//// 单向通道将阻塞线程2秒
	//t2 := <-timer1.C
	//fmt.Printf("t2:%v \n", t2)

	//2。 timer 只相应一次
	//timer2 := time.NewTimer(time.Second)
	//for {
	//	<- timer2.C
	//	fmt.Printf("时间到")
	//}

	////3。 timer 实现延时功能
	//// 1.
	//time.Sleep(time.Second)
	////	2.
	//timer3 := time.NewTimer(time.Second)
	//<- timer3.C
	//// 3.
	//<- time.After(2 * time.Second)

	// 4. close 定时器
	//timer4 := time.NewTimer(2 * time.Second)
	//
	//go func(timer *time.Timer) {
	//	<-timer.C
	//	println("time up")
	//}(timer4)
	//
	//t := timer4.Stop()
	//if t {
	//	println("timer4 is closed!")
	//}

	// 5. 重置定时器
	//timer5 := time.NewTimer(3 * time.Second)
	//timer5.Reset(time.Second)
	//fmt.Println(time.Now())
	//fmt.Println(<- timer5.C)
	//
	//for {}

	// 6. 时间刻度
	//ticker := time.NewTicker(time.Second)
	//i := 0
	//go func(tick *time.Ticker) {
	//	for {
	//		i++
	//		fmt.Println(<-tick.C)
	//		if i == 5 {
	//			tick.Stop()
	//		}
	//	}
	//}(ticker)
	//for {}

}
