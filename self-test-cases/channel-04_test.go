package test

import (
	"fmt"
	"testing"
)

var exit chan bool = make(chan bool)
var chin chan int32 = make(chan int32, 1)

func Fun() {
	exitSignal := false
	for {
		select {
		case v := <-chin:
			fmt.Printf("f v: %v\n", v)
			Inner()
			exitSignal = true
		}
		if exitSignal {
			break
		}
	}
	exit <- true
}

func Inner() {
	fmt.Printf("inner v: %v\n", <-chin)
}

func TestChannel04(t *testing.T) {
	go Fun()
	chin <- 1
	chin <- 2
	<-exit
}
