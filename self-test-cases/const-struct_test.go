package test

import (
	"fmt"
	"testing"
)

func TestConstStruct(t *testing.T) {
	fmt.Println(PrettyMapStruct(VehicleControlErr, true))
	VehicleControlErr = err{1, "task err"}
	fmt.Println(PrettyMapStruct(VehicleControlErr, true))
}
