package handlers

import (
	"test/fsm"
	"test/door/types"
)

func GetHandler_EventUnlock_StateClose() fsm.Handler[types.Payload] {
	return fsm.Handler[types.Payload]{
		BeforeHandler: _BeforeHandler_EventUnlock_StateClose,
		Handler:       _Handler_EventUnlock_StateClose,
		AfterHandler:  _AfterHandler_EventUnlock_StateClose,
	}
}

func _BeforeHandler_EventUnlock_StateClose(f *fsm.FSM[types.Payload]) error {
	switch f.Name {
    case types.FSMNameDoor:
		return _BeforeHandlerForDoor_EventUnlock_StateClose(f)
    case types.FSMNameSmartDoor:
		return _BeforeHandlerForSmartDoor_EventUnlock_StateClose(f)
	default:
		return nil
	}
}

func _Handler_EventUnlock_StateClose(f *fsm.FSM[types.Payload]) error {
	switch f.Name {
    case types.FSMNameDoor:
		return _HandlerForDoor_EventUnlock_StateClose(f)
    case types.FSMNameSmartDoor:
		return _HandlerForSmartDoor_EventUnlock_StateClose(f)
	default:
		return nil
	}

}

func _AfterHandler_EventUnlock_StateClose(f *fsm.FSM[types.Payload]) error {
	switch f.Name {
    case types.FSMNameDoor:
		return _AfterHandlerForDoor_EventUnlock_StateClose(f)
    case types.FSMNameSmartDoor:
		return _AfterHandlerForSmartDoor_EventUnlock_StateClose(f)
	default:
		return nil
	}
}


func _BeforeHandlerForDoor_EventUnlock_StateClose(f *fsm.FSM[types.Payload]) error {
// TODO: your code

	return nil
}

func _HandlerForDoor_EventUnlock_StateClose(f *fsm.FSM[types.Payload]) error {
	// TODO: your code

	return nil
}

func _AfterHandlerForDoor_EventUnlock_StateClose(f *fsm.FSM[types.Payload]) error {
	// TODO: your code

	return nil
}
func _BeforeHandlerForSmartDoor_EventUnlock_StateClose(f *fsm.FSM[types.Payload]) error {
// TODO: your code

	return nil
}

func _HandlerForSmartDoor_EventUnlock_StateClose(f *fsm.FSM[types.Payload]) error {
	// TODO: your code

	return nil
}

func _AfterHandlerForSmartDoor_EventUnlock_StateClose(f *fsm.FSM[types.Payload]) error {
	// TODO: your code

	return nil
}
