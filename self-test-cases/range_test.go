package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/zeromicro/go-zero/core/threading"
)

func TestRange(t *testing.T) {
	a := []int32{1, 2, 3}
	for range a {
		fmt.Printf("%v\n", "test")
	}
}

func TestRangeConcurrency(t *testing.T) {
	type Stu struct {
		Name string
		Age  int32
	}

	var students []*Stu = []*Stu{
		{
			Name: "A1",
			Age:  1,
		},
		{
			Name: "A2",
			Age:  2,
		},
		{
			Name: "A3",
			Age:  3,
		},
	}

	var printStu = func(stu *Stu) {
		time.Sleep(1 * time.Second)
		fmt.Printf("stu: %+v\n", stu)
	}

	for _, item := range students {
		threading.GoSafe(func() {
			printStu(item)
		})
	}
	time.Sleep(5 * time.Second)
}

func TestRangeConcurrency2(t *testing.T) {
	type Stu struct {
		Name string
		Age  int32
	}

	var students []*Stu = []*Stu{
		{
			Name: "A1",
			Age:  1,
		},
		{
			Name: "A2",
			Age:  2,
		},
		{
			Name: "A3",
			Age:  3,
		},
	}

	var printStu = func(stu *Stu) {
		time.Sleep(1 * time.Second)
		fmt.Printf("stu: %+v\n", stu)
	}

	for _, item := range students {
		threading.GoSafe(func() {
			printStu(item)
		})
	}
	time.Sleep(5 * time.Second)
}
