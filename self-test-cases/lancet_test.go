package test

import (
	"fmt"
	"testing"

	"github.com/duke-git/lancet/v2/algorithm"
	"github.com/duke-git/lancet/v2/strutil"
)

func TestLancet(t *testing.T) {
	s := "abdc"
	a := strutil.Reverse(s)
	fmt.Println(a)
	cache := algorithm.NewLRUCache[int, int](5)
	fmt.Printf("%#v\n", cache)
}
