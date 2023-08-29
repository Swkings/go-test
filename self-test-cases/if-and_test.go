package test

import (
	"fmt"
	"testing"
)

func if1(boolean bool) bool {
	fmt.Printf("IF1: %v\n", boolean)
	return boolean
}

func if2(boolean bool) bool {
	fmt.Printf("IF2: %v\n", boolean)
	return boolean
}

func TestIfAnd(t *testing.T) {
	if if1(false) && if2(false) {
		fmt.Println("IF3")
	}
}
