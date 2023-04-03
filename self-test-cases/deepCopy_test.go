package test

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"testing"
)

// func DeepCopyByJson(src map[string]string) (map[string]string, error) {
// 	var dst = make(map[string]string)
// 	b, err := json.Marshal(src)
// 	if err != nil {
// 		return nil, err
// 	}

// 	err = json.Unmarshal(b, dst)
// 	return dst, err
// }

func DeepCopyByGob(dst, src interface{}) error {
	var buffer bytes.Buffer
	if err := gob.NewEncoder(&buffer).Encode(src); err != nil {
		return err
	}

	return gob.NewDecoder(&buffer).Decode(dst)
}

func DeepCopyByJson[T any](src T) (T, error) {
	var dst T
	b, err := json.Marshal(src)
	if err != nil {
		return dst, err
	}

	err = json.Unmarshal(b, dst)
	return dst, err
}

func TestRWMap(t *testing.T) {
	a := map[string]string{
		"a": "a",
		"b": "b",
	}
	fmt.Printf("a = %v\n", a)
	b, _ := DeepCopyByJson(a)
	for k, _ := range b {
		delete(a, k)
	}
	fmt.Printf("a = %v\n", a)
	fmt.Printf("b = %v\n", b)

	c := []string{"a", "b", "c"}
	fmt.Printf("c = %v\n", c)
	d, _ := DeepCopyByJson(c)
	c = append(c, "d")
	fmt.Printf("c = %v\n", c)
	fmt.Printf("d = %v\n", d)
}

type AStruct struct {
	Name string
	Age  int
}

func changeAStruct(aStruct *AStruct) {
	aStruct.Name += aStruct.Name
	aStruct.Age += aStruct.Age
}

func TestRef(t *testing.T) {

	AList := []*AStruct{
		{
			Name: "A",
			Age:  1,
		},
		{
			Name: "B",
			Age:  2,
		},
		{
			Name: "C",
			Age:  3,
		},
	}

	fmt.Printf("AList Before: %v\n", PrettyMapStruct(AList, true))

	for _, aStruct := range AList {
		// ind := AList[i]
		changeAStruct(aStruct)
	}

	fmt.Printf("AList After: %v\n", PrettyMapStruct(AList, true))
}
