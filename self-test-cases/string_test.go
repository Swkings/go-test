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

func (s Stu) String() string {
	return s.name + s.ha
}

func TestString(t *testing.T) {
	s := &Stu{
		name: "a",
		age:  10,
		ha:   "b",
	}
	fmt.Printf("testString: %+v\n", s)
	a := Stu{
		name: "a",
		age:  10,
		ha:   "b",
	}
	b := Stu{
		age:  11,
		name: "a",
		ha:   "b",
	}
	fmt.Printf("equal: %v\n", a == b)
	aMap := map[Stu]string{}
	aMap[a] = "a"
	aMap[b] = "b"
	fmt.Printf("aMap: %v, mapLen: %v\n", aMap, len(aMap))
}
