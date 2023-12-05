package test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func F1(ctx context.Context, ch chan int) {
	fmt.Printf("Run F1\n")
	for {
		<-ctx.Done()
		fmt.Printf("F1 Recv Ctx Done\n")
		fmt.Printf("F1 Send Ctx 1\n")
		ch <- 1
		fmt.Printf("F1 Break\n")
		break
	}
}

func F2(ctx context.Context) {
	fmt.Printf("Run F2\n")
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
	fmt.Printf("Cancel Ctx\n")
	recvNum := <-ch
	fmt.Printf("Recv Num: %v\n", recvNum)
	time.Sleep(2 * time.Second)
	F2(ctx)
	time.Sleep(5 * time.Second)
	F2(ctx)
}
