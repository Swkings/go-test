package test

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := map[string]int{
		"a": 1,
	}
	fmt.Printf("map size: %v\n", len(m))
}
