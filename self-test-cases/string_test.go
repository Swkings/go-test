package test

import (
	"fmt"
	"testing"
)

type Stu struct {
	name string
	age  int32
	ha   string
}

func (s *Stu) String() string {
	return s.name + s.ha
}

func TestString(t *testing.T) {
	s := &Stu{
		name: "a",
		age:  10,
		ha:   "b",
	}
	fmt.Printf("testString: %+v\n", s)
}
