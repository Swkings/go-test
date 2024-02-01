//go:build ignore

package main

import (
	"test/door_fsm/door"
	"test/door_fsm/door/types"
	"test/fsm"
)

func main() {
	var (
		door1                                      = "door1"
		door2                                      = "door2"
		doorMap map[string]*fsm.FSM[types.Payload] = map[string]*fsm.FSM[types.Payload]{
			door1: door.NewFSMDoor(),
			door2: door.NewFSMDoor(),
		}
	)
	doorMap[door1].ProcessEvent(types.EventLock)
	doorMap[door2].ProcessEvent(types.EventLock)
}
