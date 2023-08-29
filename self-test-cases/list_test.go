package test

import (
	"fmt"
	"strings"
	"testing"
)

type StuForList struct {
	Name string
	Age  int
	Like []string
}

func TestList(t *testing.T) {
	a := []string{"a", "b", "c", "DD"}
	fmt.Printf("list: %v\n", strings.Join(a, "_"))
	b := []string{}
	fmt.Printf("list: %v\n", strings.Join(b, "_"))

	fmt.Printf("list: %v\n", strings.Join(a, "\t"))
	fmt.Printf("list: %v\n", string([]int32{1, 2, 3}))

	var s1 []StuForList
	s1 = append(s1, StuForList{"A", 1, []string{"A"}})
	fmt.Printf("list: %+v\n", s1)
	var s2 []*StuForList
	s2 = append(s2, &StuForList{"A", 1, []string{"A"}})
	fmt.Printf("list: %#v\n", s2)
}
