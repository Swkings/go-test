package test

import (
	"fmt"
	"testing"
	"time"
)

func TestGetNow(t *testing.T) {
	now := time.Now().Unix()
	fmt.Printf("now: %v\n", now)
}
