package fsmtest

import (
	"fmt"
	"test/fsm"
	"test/fsmtest/handlers"
	"test/fsmtest/types"
	"testing"
	"time"

	"github.com/zeromicro/go-zero/core/threading"
)

func TestSchedulerFSMConcurrency(t *testing.T) {
	f := fsm.NewFSM(types.FSMNameScheduler, types.StateIdle, &types.Payload{
		Data: "scheduler",
	})

	f.SetTransitions(
		fsm.NewTransitionItem(types.EventAddTask, types.StateIdle, types.StateAssigned, handlers.GetHandler_EventAddTask_StateIdle()),
	).SetTransitions(
		fsm.NewTransitionItem(types.EventStartCmd, types.StateAssigned, types.StateProcessing, handlers.GetHandler_EventAddTask_StateIdle()),
		fsm.NewTransitionItem(types.EventStartCmd, types.StateAssignedMaintenance, types.StateProcessingMaintenance, handlers.GetHandler_EventAddTask_StateIdle()),
	).SetTransitions(
		fsm.NewTransitionItem(types.EventDoneCmd, types.StateProcessing, types.StateCompleted, handlers.GetHandler_EventAddTask_StateIdle()),
		fsm.NewTransitionItem(types.EventDoneCmd, types.StateProcessingMaintenance, types.StateCompletedMaintenance, handlers.GetHandler_EventAddTask_StateIdle()),
	).SetTransitions(
		fsm.NewTransitionItem(types.EventAbortCmd, types.StateAssigned, types.StateAbort, handlers.GetHandler_EventAddTask_StateIdle()),
		fsm.NewTransitionItem(types.EventAbortCmd, types.StateProcessing, types.StateAbort, handlers.GetHandler_EventAddTask_StateIdle()),
		fsm.NewTransitionItem(types.EventAbortCmd, types.StateAssignedMaintenance, types.StateAbortMaintenance, handlers.GetHandler_EventAddTask_StateIdle()),
		fsm.NewTransitionItem(types.EventAbortCmd, types.StateProcessingMaintenance, types.StateAbortMaintenance, handlers.GetHandler_EventAddTask_StateIdle()),
	).SetTransitions(
		fsm.NewTransitionItem(types.EventOperation, types.StateIdle, types.StateMaintenance, handlers.GetHandler_EventAddTask_StateIdle()),
		fsm.NewTransitionItem(types.EventOperation, types.StateAssigned, types.StateAssignedMaintenance, handlers.GetHandler_EventAddTask_StateIdle()),
		fsm.NewTransitionItem(types.EventOperation, types.StateProcessing, types.StateProcessingMaintenance, handlers.GetHandler_EventAddTask_StateIdle()),
	).SetTransitions(
		fsm.NewTransitionItem(types.EventMaintenance, types.StateCompletedMaintenance, types.StateMaintenance, handlers.GetHandler_EventAddTask_StateIdle()),
		fsm.NewTransitionItem(types.EventMaintenance, types.StateAbortMaintenance, types.StateMaintenance, handlers.GetHandler_EventAddTask_StateIdle()),
	).SetTransitions(
		fsm.NewTransitionItem(types.EventClearTask, types.StateCompleted, types.StateIdle, handlers.GetHandler_EventAddTask_StateIdle()),
		fsm.NewTransitionItem(types.EventClearTask, types.StateAbort, types.StateIdle, handlers.GetHandler_EventAddTask_StateIdle()),
	).SetTransitions(
		fsm.NewTransitionItem(types.EventRepair, types.StateMaintenance, types.StateIdle, handlers.GetHandler_EventAddTask_StateIdle()),
	)

	var err error
	threading.GoSafe(func() {
		_, err = f.ProcessEvent(types.EventAddTask)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	})

	threading.GoSafe(func() {
		_, err = f.ProcessEvent(types.EventDoneCmd)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	})

	threading.GoSafe(func() {
		_, err = f.ProcessEvent(types.EventStartCmd)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	})

	threading.GoSafe(func() {
		_, err = f.ProcessEvent(types.EventOperation)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	})

	threading.GoSafe(func() {
		_, err = f.ProcessEvent(types.EventDoneCmd)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	})

	threading.GoSafe(func() {
		_, err = f.ProcessEvent(types.EventMaintenance)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	})

	threading.GoSafe(func() {
		_, err = f.ProcessEvent(types.EventRepair)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	})

	<-time.After(20 * time.Second)
}
