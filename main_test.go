package gotest

import (
	"fmt"
	test "test/self-test-cases"
	"testing"
)

// type errI interface {
// 	GetCode() int32
// 	GetMsg() string
// 	String() string
// 	private()
// }
// type err struct {
// 	Code int32
// 	Msg  string
// }

// func (e err) GetCode() int32 {
// 	return e.Code
// }

// func (e err) GetMsg() string {
// 	return e.Msg
// }

// func (e err) String() string {
// 	return fmt.Sprintf("%v: %v", e.Code, e.Msg)
// }

// func (e err) private() {
// }

func TestConstStruct(t *testing.T) {
	fmt.Println(test.PrettyMapStruct(test.VehicleControlErr, true))
	// test.VehicleControlErr = err{1, "task err"}
	fmt.Println(test.ConstA.GetValue())
	fmt.Println(test.ConstB.GetValue())
	fmt.Println(test.ConstC.GetValue())
	fmt.Println(test.ConstC)
	a := []int32{1, 2, 3}
	b := struct {
		A []int32
		B string
	}{
		A: a,
		B: "test",
	}
	fmt.Println(test.PrettyMapStruct(b, false))
}
