package test

import (
	"fmt"
	"testing"
)

func TestBinInt(t *testing.T) {
	a := 0b10
	b := 2
	fmt.Printf("a:%v  b:%v, Equal: %v\n", a, b, a == b)
}
