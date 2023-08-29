package test

import "fmt"

type errI interface {
	GetCode() int32
	GetMsg() string
	String() string
	private()
}
type err struct {
	Code int32
	Msg  string
}

func (e err) GetCode() int32 {
	return e.Code
}

func (e err) GetMsg() string {
	return e.Msg
}

func (e err) String() string {
	return fmt.Sprintf("%v: %v", e.Code, e.Msg)
}

func (e err) private() {
}

var (
	VehicleControlErr errI = err{0, "vehicle err"}
)

type constI[T any] interface {
	GetValue() T
	private()
}

type constV[T any] struct {
	Value T
}

func (c constV[T]) GetValue() T {
	return c.Value
}

func (c constV[T]) private() {
}

type A struct {
	Name string
	Age  int32
}

func (a A) String() string {
	return fmt.Sprintf("Name: %v, Age: %v", a.Name, a.Age)
}

var (
	ConstA constI[string] = constV[string]{"111"}
	ConstB constI[int32]  = constV[int32]{111}
	ConstC constI[A]      = constV[A]{
		A{"111", 111},
	}
)
