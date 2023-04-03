package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/zeromicro/go-zero/core/threading"
)

var a = func(a, b int32) func() {
	return func() {
		for i := 1; i <= 10; i++ {
			fmt.Println(a + b)
			time.Sleep(time.Second * 1)
			c := int32(0)
			a = b / c
		}
	}
}

var b = func(a, b int32) func() {
	return func() {
		for i := 1; i <= 10; i++ {
			fmt.Println(a + b)
			time.Sleep(time.Second * 2)
		}
	}
}

func TestThreading(t *testing.T) {

	threading.GoSafe(a(1, 2))
	threading.GoSafe(b(10, 20))
	// go a(1, 2)()
	// go b(10, 20)()

	time.Sleep(time.Second * 10)
}
