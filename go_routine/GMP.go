package go_routine

import (
	"os"
	"runtime/trace"
)

func FuncGMP() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	defer trace.Stop()

	println("hello world!")

	// go run this function then get the trace.out file
	// go tool trace trace.out
}
