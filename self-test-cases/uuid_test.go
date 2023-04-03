package test

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func TestUUID(t *testing.T) {
	str := uuid.NewString()
	fmt.Printf("uuid: %v\n", str)
}
