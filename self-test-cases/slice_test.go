package test

import (
	"fmt"
	"testing"
)

func TestSliceDelete(t *testing.T) {
	a := []int{1, 2, 3, 4, 5, 6}

	fmt.Printf("Origin: %+v\n", a)
	for i, item := range a {
		if item == 2 || item == 4 || item == 6 {
			if i < len(a)-1 {
				a = append(a[:i], a[i+1:]...)
			} else {
				a = append(a[:i])
			}
		}

		fmt.Printf("%v: %v, left: %+v\n", i, item, a)
	}
	fmt.Printf("Origin Left: %+v\n", a)
}
