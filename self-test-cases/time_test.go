package test

import (
	"fmt"
	"testing"
	"time"
)

func TestGetNow(t *testing.T) {
	now := time.Now().Unix()
	fmt.Printf("now: %v\n", now)

	a := time.Duration(10) * time.Second
	time.Sleep(a)
	fmt.Printf("duration: %v\n", a)
}
