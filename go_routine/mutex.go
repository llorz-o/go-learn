package go_routine

import (
	"fmt"
	"sync"
	"time"
)

var x int
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock()
		x = x + 1
		lock.Unlock()
	}
	wg.Done()
}

func FuncMutex()  {
	wg.Add(2)
	go add()
	go add()
	wg.Wait()
	fmt.Println(x)
}


var rwLock sync.RWMutex
var a int
func write()  {
	rwLock.Lock()
	a = a + 1
	time.Sleep(10 * time.Millisecond)
	rwLock.Unlock()
	wg.Done()
}

func read(){
	rwLock.RLock()
	time.Sleep(time.Millisecond)
	rwLock.RUnlock()
	wg.Done()
}

func FuncRWMutex()  {
	//读写锁分为两种：读锁和写锁。当一个goroutine获取读锁之后，其他的goroutine如果是获取读锁会继续获得锁，
	//如果是获取写锁就会等待；当一个goroutine获取写锁之后，其他的goroutine无论是获取读锁还是写锁都会等待。
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start)) // ~ 112ms
}