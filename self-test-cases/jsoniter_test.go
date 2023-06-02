package test

import (
	"fmt"
	"os"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

func TestJsoniter(t *testing.T) {
	type Like struct {
		One string
		Two string
	}
	type student struct {
		Name string
		Age  int64
		Like Like
	}
	stu := student{
		Name: "AA",
		Age:  10,
		Like: Like{
			One: "1",
			Two: "2",
		},
	}
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	encode, _ := json.Marshal(&stu)
	os.Stdout.Write(encode)
	var sCopy student
	json.Unmarshal(encode, &sCopy)
	fmt.Printf("\ncopy data: %+v\n", sCopy)
}
