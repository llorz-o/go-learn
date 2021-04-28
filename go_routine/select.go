package go_routine

import (
	"fmt"
)

func t1(ch chan string) {
	//time.Sleep(time.Second * 5)
	ch <- "test1"
}

func t2(ch chan string) {
	//time.Sleep(time.Second * 2)
	ch <- "test2"
}

func FuncSelect() {
	output1 := make(chan string,128)
	output2 := make(chan string,128)
	go t1(output1)
	go t2(output2)

	select {
	case s1 := <-output1:
		fmt.Println("s1=", s1)
	case s2 := <-output2:
		fmt.Println("s2=", s2)
	}
}
