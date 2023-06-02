//go:build ignore

package main

import (
	"fmt"
	"test/door"
	"test/door/types"
)

func main() {
	d := door.NewFSMDoor()
	trace, err := d.ProcessEvent(types.EventLock)
	fmt.Printf("trace: %+v, err: %v\n", trace, err)
	trace, err = d.ProcessEvent(types.EventUnlock)
	fmt.Printf("trace: %+v, err: %v\n", trace, err)
}
