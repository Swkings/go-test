package handlers

import (
	"fmt"
	"test/fsm"
	"test/fsmtest/types"
	"time"
)

func _BeforeHandlerForScheduler(f *fsm.FSM[types.Payload]) error {
	// TODO: your code

	return nil
}

func _HandlerForScheduler(f *fsm.FSM[types.Payload]) error {
	for t := 1; t <= 2; t++ {
		time.Sleep(1 * time.Second)
		fmt.Printf("transfer event %v->%v, sleep %vth s\n", f.GetCurrentState(), fsm.FSMEvent("AddTask"), t)
	}

	return nil
}

func _AfterHandlerForScheduler(f *fsm.FSM[types.Payload]) error {
	// TODO: your code

	return nil
}

func _BeforeHandler(f *fsm.FSM[types.Payload]) error {
	switch f.Name {
	case types.FSMNameScheduler:
		return _BeforeHandlerForScheduler(f)
	default:
		return nil
	}
}

func _Handler(f *fsm.FSM[types.Payload]) error {
	switch f.Name {
	case types.FSMNameScheduler:
		return _HandlerForScheduler(f)
	default:
		return nil
	}

}

func _AfterHandler(f *fsm.FSM[types.Payload]) error {
	switch f.Name {
	case types.FSMNameScheduler:
		return _AfterHandlerForScheduler(f)
	default:
		return nil
	}
}

func GetHandler_EventAddTask_StateIdle() fsm.Handler[types.Payload] {
	return fsm.Handler[types.Payload]{
		BeforeHandler: _BeforeHandler,
		Handler:       _Handler,
		AfterHandler:  _AfterHandler,
	}
}
