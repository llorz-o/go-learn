package go_routine

import (
	"sync"
	"time"
)

var wg2 sync.WaitGroup

func show(i int) {
	println("my id is:", i)
	wg2.Done() // 完成后减少计数
}

func FuncAsync() {
	wg2.Add(10) // 计数
	for i := 0; i < 10; i++ {
		//wg2.Add(1)
		go show(i)
	}
	println("wait...")
	wg2.Wait() // 等待并发任务执行完毕
}

var icons map[string]string

var loadIconsOnce sync.Once

func loadIcons() {

	//icons = map[string]string{
	//	"left": "left.png",
	//	"up": "up.png",
	//	"right": "right.png",
	//	"down": "down.png",
	//}

	// 现代编译器和cpu在保证每个 goroutine 满足串行一致性的基础上进行自由的重排访问内存的顺序
	// 可能会被重排为以下
	icons = make(map[string]string)
	time.Sleep(time.Second) // 使用 Sleep 模拟图片文件的加载时间
	icons["left"] = "left.png"
}

func Icon(name string, id int) {
	if icons == nil {
		loadIcons()
	}
	// 可使用 loadIconsOnce.Do(loadIcons) 保证并发安全以及仅执行一次
	//loadIconsOnce.Do(loadIcons)
	println(icons[name], id)
	wg.Done()
}

func FuncOnce() {
	// 在并发模式下，经过重拍的函数执行时初始化 icons 不为nil，其他goroutine访问时可能会发生错误
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go Icon("left", i)
	}
	wg.Wait()
}


// go中的内置map也是非并发安全的,并发访问时使用
var m = sync.Map{}

func FuncSyncMap()  {
	m.Store("key","value")
	value,_ := m.Load("key")
	println("sync map get value by key:key",value)
}