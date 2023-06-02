package handlers

import (
	"test/fsm"
	"test/door/types"
)

func GetHandler_EventLock_StateOpen() fsm.Handler[types.Payload] {
	return fsm.Handler[types.Payload]{
		BeforeHandler: _BeforeHandler_EventLock_StateOpen,
		Handler:       _Handler_EventLock_StateOpen,
		AfterHandler:  _AfterHandler_EventLock_StateOpen,
	}
}

func _BeforeHandler_EventLock_StateOpen(f *fsm.FSM[types.Payload]) error {
	switch f.Name {
    case types.FSMNameDoor:
		return _BeforeHandlerForDoor_EventLock_StateOpen(f)
    case types.FSMNameSmartDoor:
		return _BeforeHandlerForSmartDoor_EventLock_StateOpen(f)
	default:
		return nil
	}
}

func _Handler_EventLock_StateOpen(f *fsm.FSM[types.Payload]) error {
	switch f.Name {
    case types.FSMNameDoor:
		return _HandlerForDoor_EventLock_StateOpen(f)
    case types.FSMNameSmartDoor:
		return _HandlerForSmartDoor_EventLock_StateOpen(f)
	default:
		return nil
	}

}

func _AfterHandler_EventLock_StateOpen(f *fsm.FSM[types.Payload]) error {
	switch f.Name {
    case types.FSMNameDoor:
		return _AfterHandlerForDoor_EventLock_StateOpen(f)
    case types.FSMNameSmartDoor:
		return _AfterHandlerForSmartDoor_EventLock_StateOpen(f)
	default:
		return nil
	}
}


func _BeforeHandlerForDoor_EventLock_StateOpen(f *fsm.FSM[types.Payload]) error {
// TODO: your code

	return nil
}

func _HandlerForDoor_EventLock_StateOpen(f *fsm.FSM[types.Payload]) error {
	// TODO: your code

	return nil
}

func _AfterHandlerForDoor_EventLock_StateOpen(f *fsm.FSM[types.Payload]) error {
	// TODO: your code

	return nil
}
func _BeforeHandlerForSmartDoor_EventLock_StateOpen(f *fsm.FSM[types.Payload]) error {
// TODO: your code

	return nil
}

func _HandlerForSmartDoor_EventLock_StateOpen(f *fsm.FSM[types.Payload]) error {
	// TODO: your code

	return nil
}

func _AfterHandlerForSmartDoor_EventLock_StateOpen(f *fsm.FSM[types.Payload]) error {
	// TODO: your code

	return nil
}
