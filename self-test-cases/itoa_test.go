package test

import (
	"fmt"
	"strconv"
	"testing"
)

func TestItoa(t *testing.T) {
	var i int32 = -1
	fmt.Printf("str: %v\n", strconv.Itoa(int(i)))
}
