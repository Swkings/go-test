package test

import (
	"fmt"
	"testing"
)

func TestRange(t *testing.T) {
	a := []int32{1, 2, 3}
	for range a {
		fmt.Printf("%v\n", "test")
	}
}
