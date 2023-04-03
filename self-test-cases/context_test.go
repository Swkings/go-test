package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func F1(ctx context.Context, ch chan int) {
	for {
		<-ctx.Done()
		fmt.Printf("F1 Done\n")
		ch <- 1
		break
	}
}

func F2(ctx context.Context) {
	for {
		<-ctx.Done()
		fmt.Printf("F2 Done\n")
		break
	}
}

func TestContext(t *testing.T) {
	ch := make(chan int, 1)
	ctx, cancel := context.WithCancel(context.Background())

	go F1(ctx, ch)
	cancel()
	<-ch
	time.Sleep(2 * time.Second)
	F2(ctx)
	time.Sleep(5 * time.Second)
	F2(ctx)
}
