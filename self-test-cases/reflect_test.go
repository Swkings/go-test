package test

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {

	type Person struct {
		Name string
		Age  int
	}

	arr := []interface{}{
		&Person{Name: "Tom", Age: 20},
		&Person{},
		&Person{Name: "Jerry", Age: 18},
	}

	for i, elem := range arr {
		elemType := reflect.TypeOf(elem)
		if elemType.Kind() != reflect.Ptr {
			fmt.Printf("arr[%d] is empty struct\n", i)
			continue
		}

		zeroElem := reflect.New(elemType.Elem()).Elem().Interface()
		if reflect.DeepEqual(zeroElem, reflect.ValueOf(elem).Elem().Interface()) {
			fmt.Printf("arr[%d] is empty struct\n", i)
		} else {
			fmt.Printf("arr[%d] is not empty\n", i)
		}
	}

}
