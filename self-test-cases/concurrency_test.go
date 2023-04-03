package test

import (
	"fmt"
	"testing"
)

type Stud struct {
	Name string
	Age  int64
}

func TestConcurrency(t *testing.T) {
	m := map[string]*Stud{
		"11": {Name: "11", Age: 11},
		"22": {Name: "22", Age: 22},
		"33": {Name: "33", Age: 33},
	}

	m11 := m["11"]
	fmt.Printf("m11: %v\n", m11)
	m["11"].Name = "test"
	fmt.Printf("m11: %v\n", m11)
}
