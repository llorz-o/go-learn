package go_routine

import (
	"sync"
	"sync/atomic"
	"time"
)

// 使用加锁机制保证并发安全会导致性能损耗
// 通常可以使用基本数据的原子操作来保证操作的原子性
// int32 int64 uint32 uint64 uintptr Pointer

var z int64
var l sync.Mutex

func addNumberZ() {
	z++
	wg.Done()
}

func mutexAdd() {
	l.Lock()
	z++
	l.Unlock()
	wg.Done()
}

func atomicAdd() {
	atomic.AddInt64(&z, 1)
	wg.Done()
}

func FuncAtomic() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		//go addNumberZ() // 9312 3269629
		//go mutexAdd()  // 10000 3859360
		go atomicAdd() // 10000   3156584
	}

	wg.Wait()
	end := time.Now()
	println(z, end.Sub(start))
}
