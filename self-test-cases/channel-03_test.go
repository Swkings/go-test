package test

import (
	"fmt"
	"testing"
	"time"
)

func ExitChan(c *chan int) {

	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("Sleep %d\n", i)
	}
	fmt.Println("ExitChan")
	*c <- 0
}

func TestChanExit(t *testing.T) {
	ch := make(chan int, 2)
	go ExitChan(&ch)

	a := <-ch // 通道关闭后再取值会panic
	fmt.Printf("ch: %#v\n", ch)
	close(ch) // 关闭通道
	fmt.Println(a)

	fmt.Println("ok")
}
