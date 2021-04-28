package main

import (
	"fmt"
	"testing"
)

func TestNilSlice(t *testing.T) {
	var s []int
	if s == nil {
		fmt.Printf("s:%v\n\n", s)
		fmt.Printf("s len:%v\n\n", len(s))
		fmt.Printf("s cap:%v\n\n", cap(s))
	}
}
