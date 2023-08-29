package test

import (
	"fmt"
	"testing"
)

func TestTable(t *testing.T) {
	table := [][]int64{
		{1, 2, 3, 4},
		{11, 22, 33, 44},
		{111, 222, 333, 444},
		{1111, 2222, 3333, 4444},
		{111111111111, 2222, 33333, 4444444444444444444},
	}

	title := []string{"AA", "BBBB", "CCCC", "DD"}
	// table = nil
	// title = nil
	fmt.Printf("%v\n", PrettyArrayTable(title, table))
	type AA struct {
		Name string
		Age  int32
		Data struct {
			Field1 string
			Field2 map[string]int64
		}
	}

	structList := []*AA{
		&AA{
			Name: "11",
			Age:  1,
			Data: struct {
				Field1 string
				Field2 map[string]int64
			}{
				Field1: "11",
				Field2: map[string]int64{
					"1": 1,
				},
			},
		},
		nil,
		&AA{
			Name: "22",
			Age:  2,
			Data: struct {
				Field1 string
				Field2 map[string]int64
			}{
				Field1: "22",
				Field2: map[string]int64{
					"2": 2,
				},
			},
		},
		&AA{},
	}

	structList2 := []AA{
		AA{
			Name: "11",
			Age:  1,
			Data: struct {
				Field1 string
				Field2 map[string]int64
			}{
				Field1: "11",
				Field2: map[string]int64{
					"1": 1,
				},
			},
		},
		AA{
			Name: "22",
			Age:  2,
			Data: struct {
				Field1 string
				Field2 map[string]int64
			}{
				Field1: "22",
				Field2: map[string]int64{
					"2": 2,
				},
			},
		},
		AA{},
	}
	fmt.Printf("%v\n", PrettyStructTable(structList))
	fmt.Printf("%v\n", PrettyStructTable(structList2, false))
}

func TestPrettyStructTableTest(t *testing.T) {
	type T struct {
		name string
		age  int32
	}
	type args struct {
		elementList []T
		rightAlign  []bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "t1",
			args: args{elementList: []T{
				{"t1", 11},
				{"t2", 22},
			}, rightAlign: []bool{false}},
			want: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PrettyStructTable(tt.args.elementList, tt.args.rightAlign...); got != tt.want {
				t.Errorf("PrettyStructTable() = \n%v, \nwant \n%v", got, tt.want)
			}
		})
	}
}
