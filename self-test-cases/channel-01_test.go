package test

import (
	"fmt"
	"testing"
	"time"
)

// 判断管道有没有存满
func TestChannel01(t *testing.T) {
	// 创建管道
	output1 := make(chan string, 5)
	// 子协程写数据
	go write(output1)
	// 取数据
	for s := range output1 {
		fmt.Println("res:", s)
		time.Sleep(time.Second)
	}
}

func write(ch chan string) {
out:
	for {
		select {
		// 写数据
		case ch <- "hello":
			fmt.Println("write hello")
		default:
			fmt.Println("channel full")
			close(ch)
			break out
		}
		time.Sleep(time.Millisecond * 500)
	}
}
