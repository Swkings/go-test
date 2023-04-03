package test

import (
	"fmt"
	"testing"
	"time"
)

var ch = make(chan int)

func TC() {
	select {
	case <-ch:
		fmt.Println("case1")
	case ch <- 1:
		fmt.Println("case2")
	}
	fmt.Println("TEST")
}
func TestSelect(t *testing.T) {
	go TC()
	time.Sleep(time.Second * 2)
	ch <- 1
	time.Sleep(time.Second * 10)
}
