package core

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	in := &In{}
	for i := 0; i < 10; i++ {
		in.Add(i)
	}

	fmt.Println(in.Get())
}
