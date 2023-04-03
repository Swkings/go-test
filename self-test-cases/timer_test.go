package test

import (
	"fmt"
	"testing"
	"time"
)

func sleepTime(t int64) {
	for i := int64(1); i <= t; i++ {
		fmt.Printf("sleep: %v\n", i)
		time.Sleep(time.Second * time.Duration(1))
	}

}

func TestTimer(t *testing.T) {
	ch := make(chan string)
	stopStayTime := time.Duration(5) * time.Second
	timer := time.NewTimer(stopStayTime)
	timerBeginTime := time.Now()
	exitSignal := false
	go func() {
		sleepTime(3)
		ch <- "stop timer"
	}()

	for {
		select {
		case <-timer.C:
			exitSignal = true
			fmt.Printf("timeout: %v\n", time.Since(timerBeginTime))
		case event := <-ch:
			fmt.Printf("recv chan event: %v\n", event)
			timeRemainder := stopStayTime - time.Since(timerBeginTime)
			timer.Stop()
			sleepTime(10)
			timer.Reset(timeRemainder)
			fmt.Printf("time recover: %v\n", timeRemainder)
		}
		if exitSignal {
			break
		}
	}
}
