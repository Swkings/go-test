// Code generated by fsm-gen. DO NOT EDIT.
// source: scheduler.fsm

package types

import "test/fsm"

var (
	FSMNameScheduler = fsm.FSMName("Scheduler")
)

var (
	StateIdle                  = fsm.FSMState("StateIdle")
	StateAssigned              = fsm.FSMState("StateAssigned")
	StateProcessing            = fsm.FSMState("StateProcessing")
	StateCompleted             = fsm.FSMState("StateCompleted")
	StateAbort                 = fsm.FSMState("StateAbort")
	StateMaintenance           = fsm.FSMState("StateMaintenance")
	StateAssignedMaintenance   = fsm.FSMState("StateAssignedMaintenance")
	StateProcessingMaintenance = fsm.FSMState("StateProcessingMaintenance")
	StateCompletedMaintenance  = fsm.FSMState("StateCompletedMaintenance")
	StateAbortMaintenance      = fsm.FSMState("StateAbortMaintenance")
)

var (
	EventAddTask     = fsm.FSMEvent("EventAddTask")
	EventStartCmd    = fsm.FSMEvent("EventStartCmd")
	EventDoneCmd     = fsm.FSMEvent("EventDoneCmd")
	EventAbortCmd    = fsm.FSMEvent("EventAbortCmd")
	EventMaintenance = fsm.FSMEvent("EventMaintenance")
	EventOperation   = fsm.FSMEvent("EventOperation")
	EventRepair      = fsm.FSMEvent("EventRepair")
	EventClearTask   = fsm.FSMEvent("EventClearTask")
)
