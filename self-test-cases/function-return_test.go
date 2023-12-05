package test

import (
	"fmt"
	"testing"
)

func TestFunctionReturn(t *testing.T) {
	l := []int{1, 2, 3, 4, 5, 6}
	l1, l2 := func() (l1 []int, l2 []int) {
		for _, item := range l {
			if item%2 == 0 {
				l2 = append(l2, item)
			} else {
				l1 = append(l1, item)
			}
		}
		return l1, l2
	}()

	fmt.Printf("l1: %+v\n", l1)
	fmt.Printf("l2: %+v\n", l2)
}
