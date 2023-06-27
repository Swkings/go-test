package test

import (
	"fmt"
	"testing"
)

func Num(num int) {
	fmt.Println(num)
}

func DeferFunc() {
	// if true {
	// 	return
	// }
	defer Num(1)
	defer Num(2)
	defer Num(3)
	defer Num(4)
	defer Num(5)
	fmt.Println(0)
}

func DeferOuter() {
	defer Num(6)
	defer Num(7)
	DeferFunc()
}

func TestDefer(t *testing.T) {
	DeferOuter() // output: 0, 5, 4, 3, 2, 1, 7, 6
}
