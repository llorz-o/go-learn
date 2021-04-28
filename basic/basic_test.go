package basic

import (
	"fmt"
	"testing"
)

func ExampleRun() {
	fmt.Println("Example Run")
}

func TestDivide(t *testing.T) {
	t.Name()
	Divide("test divide", func() {
		var re = 0
		for i := 0; i < 1000; i++ {
			re += i
		}
		t.Run("run fn",func(t *testing.T) {
			t.Log("done")
		})
	})
}

func TestDivideGroup(t *testing.T) {
	// 使用测试组与子测试
	//tests := map[string]func() {
	//	"t1": func() {
	//
	//	},
	//}
	tests := []struct{
		name string
		f func()
	} {
		{
			name: "t1",
			f: func() {
				println("run fn t1")
			},
		},
		{
			name: "t2",
			f: func() {
				println("run fn t2")
			},
		},
	}

	for _,ti := range tests {
		// 子测试
		t.Run(ti.name, func(t *testing.T) {
			ti.f()
		})
	}
}

func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Split("枯藤老树昏鸦", "老")
	}
}