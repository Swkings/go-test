package handlers

import (
	"test/fsm"
	"test/door/types"
)

func GetHandler_EventUnlock_StateOpen() fsm.Handler[types.Payload] {
	return fsm.Handler[types.Payload]{
		BeforeHandler: _BeforeHandler_EventUnlock_StateOpen,
		Handler:       _Handler_EventUnlock_StateOpen,
		AfterHandler:  _AfterHandler_EventUnlock_StateOpen,
	}
}

func _BeforeHandler_EventUnlock_StateOpen(f *fsm.FSM[types.Payload]) error {
	switch f.Name {
    case types.FSMNameSmartDoor:
		return _BeforeHandlerForSmartDoor_EventUnlock_StateOpen(f)
	default:
		return nil
	}
}

func _Handler_EventUnlock_StateOpen(f *fsm.FSM[types.Payload]) error {
	switch f.Name {
    case types.FSMNameSmartDoor:
		return _HandlerForSmartDoor_EventUnlock_StateOpen(f)
	default:
		return nil
	}

}

func _AfterHandler_EventUnlock_StateOpen(f *fsm.FSM[types.Payload]) error {
	switch f.Name {
    case types.FSMNameSmartDoor:
		return _AfterHandlerForSmartDoor_EventUnlock_StateOpen(f)
	default:
		return nil
	}
}


func _BeforeHandlerForSmartDoor_EventUnlock_StateOpen(f *fsm.FSM[types.Payload]) error {
// TODO: your code

	return nil
}

func _HandlerForSmartDoor_EventUnlock_StateOpen(f *fsm.FSM[types.Payload]) error {
	// TODO: your code

	return nil
}

func _AfterHandlerForSmartDoor_EventUnlock_StateOpen(f *fsm.FSM[types.Payload]) error {
	// TODO: your code

	return nil
}
