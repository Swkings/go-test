package test

import (
	"fmt"
	"testing"
)

func TestForList(t *testing.T) {
	// list := []int64{1, 2, 3}
	var list []int64 = []int64{1, 2, 3}
	for i := 0; i < len(list); i++ {
		fmt.Printf("%v: %v, len: %v\n", i, list[i], len(list))
		if i == 1 {
			list = append(list, 4)
		}
	}
	fmt.Printf("%v\n", list)
}
