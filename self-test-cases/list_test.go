package test

import (
	"fmt"
	"strings"
	"testing"
)

func TestList(t *testing.T) {
	a := []string{"a", "b", "c", "DD"}
	fmt.Printf("list: %v\n", strings.Join(a, "_"))
	b := []string{}
	fmt.Printf("list: %v\n", strings.Join(b, "_"))
}
