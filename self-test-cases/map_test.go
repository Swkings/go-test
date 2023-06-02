package test

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m := map[string]interface{}{
		"a": 1,
		"b": "<",
	}
	fmt.Printf("map size: %v\n", len(m))
	fmt.Printf("item: %v, %v\n", m["2"], ">")
	fmt.Printf("map: %v\n", PrettyMapStruct(m, true))

	b := map[string]int64{}
	fmt.Printf("v: %v\n", b["a"])

	c := map[string]interface{}{
		"a": 1,
		"b": "<",
		"c": "<",
		"d": "<",
	}
	fmt.Printf("c: %v\n", c)
	for id := range c {
		if id == "a" || id == "d" {
			delete(c, id)
		}
	}
	fmt.Printf("c: %v\n", c)
}
