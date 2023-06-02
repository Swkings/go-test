package fsm

import (
	"fmt"
	"testing"
	"time"

	"github.com/zeromicro/go-zero/core/threading"
)

var (
	Start = FSMState("start")
	End   = FSMState("end")
)

var (
	Process = FSMEvent("process")
	Reset   = FSMEvent("reset")
)

type Payload struct {
	Data string
}

var GlobalHandler = Handler[Payload]{
	Handler: func(f *FSM[Payload]) error {
		// fmt.Printf("transfer event %v->%v\n", f.currentState, f.processingEvent)
		for t := 1; t <= 2; t++ {
			time.Sleep(1 * time.Second)
			fmt.Printf("transfer event %v->%v, sleep %vth s\n", f.currentState, f.processingEvent, t)
		}

		return nil
	},
}

func TestFSM(t *testing.T) {
	fsm := NewFSM("testFSM", Start, &Payload{
		Data: "test",
	})
	fsm.SetTransitions(TransitionItem[Payload]{
		Event:     Process,
		FromState: Start,
		ToState:   End,
		Handler:   GlobalHandler,
	},
	).SetTransitions(
		TransitionItem[Payload]{Reset, End, Start, GlobalHandler},
		TransitionItem[Payload]{Reset, Start, Start, GlobalHandler},
	)

	fmt.Printf("Payload: %v\n", fsm.Payload.Data)

	_, err := fsm.ProcessEvent(Process)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("currentState: %v\n", fsm.GetCurrentState())

	_, err = fsm.ProcessEvent(Process)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("currentState: %v\n", fsm.GetCurrentState())

	_, err = fsm.ProcessEvent(Reset)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("currentState: %v\n", fsm.GetCurrentState())

	_, err = fsm.ProcessEvent(Reset)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("currentState: %v\n", fsm.GetCurrentState())

	_, err = fsm.ProcessEvent(Process)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("currentState: %v\n", fsm.GetCurrentState())

	_, err = fsm.ProcessEvent(Process)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("currentState: %v\n", fsm.GetCurrentState())

	fsm.VisualizeGraphviz("./fig/test.png")
}

var (
	StateIdle                  = FSMState("StateIdle")
	StateAssigned              = FSMState("StateAssigned")
	StateProcessing            = FSMState("StateProcessing")
	StateCompleted             = FSMState("StateCompleted")
	StateAbort                 = FSMState("StateAbort")
	StateMaintenance           = FSMState("StateMaintenance")
	StateAssignedMaintenance   = FSMState("StateAssignedMaintenance")
	StateProcessingMaintenance = FSMState("StateProcessingMaintenance")
	StateCompletedMaintenance  = FSMState("StateCompletedMaintenance")
	StateAbortMaintenance      = FSMState("StateAbortMaintenance")
)

var (
	EventAddTask     = FSMEvent("EventAddTask")
	EventStartCmd    = FSMEvent("EventStartCmd")
	EventDoneCmd     = FSMEvent("EventDoneCmd")
	EventAbortCmd    = FSMEvent("EventAbortCmd")
	EventMaintenance = FSMEvent("EventMaintenance")
	EventOperation   = FSMEvent("EventOperation")
	EventRepair      = FSMEvent("EventRepair")
	EventClearTask   = FSMEvent("EventClearTask")
)

func TestSchedulerFSM(t *testing.T) {
	fsm := NewFSM("testFSM", StateIdle, &Payload{
		Data: "scheduler",
	})

	fsm.SetTransitions(
		TransitionItem[Payload]{EventAddTask, StateIdle, StateAssigned, GlobalHandler},
	).SetTransitions(
		TransitionItem[Payload]{EventStartCmd, StateAssigned, StateProcessing, GlobalHandler},
		TransitionItem[Payload]{EventStartCmd, StateAssignedMaintenance, StateProcessingMaintenance, GlobalHandler},
	).SetTransitions(
		TransitionItem[Payload]{EventDoneCmd, StateProcessing, StateCompleted, GlobalHandler},
		TransitionItem[Payload]{EventDoneCmd, StateProcessingMaintenance, StateCompletedMaintenance, GlobalHandler},
	).SetTransitions(
		TransitionItem[Payload]{EventAbortCmd, StateAssigned, StateAbort, GlobalHandler},
		TransitionItem[Payload]{EventAbortCmd, StateProcessing, StateAbort, GlobalHandler},
		TransitionItem[Payload]{EventAbortCmd, StateAssignedMaintenance, StateAbortMaintenance, GlobalHandler},
		TransitionItem[Payload]{EventAbortCmd, StateProcessingMaintenance, StateAbortMaintenance, GlobalHandler},
	).SetTransitions(
		TransitionItem[Payload]{EventOperation, StateIdle, StateMaintenance, GlobalHandler},
		TransitionItem[Payload]{EventOperation, StateAssigned, StateAssignedMaintenance, GlobalHandler},
		TransitionItem[Payload]{EventOperation, StateProcessing, StateProcessingMaintenance, GlobalHandler},
	).SetTransitions(
		TransitionItem[Payload]{EventMaintenance, StateCompletedMaintenance, StateMaintenance, GlobalHandler},
		TransitionItem[Payload]{EventMaintenance, StateAbortMaintenance, StateMaintenance, GlobalHandler},
	).SetTransitions(
		TransitionItem[Payload]{EventClearTask, StateCompleted, StateIdle, GlobalHandler},
		TransitionItem[Payload]{EventClearTask, StateAbort, StateIdle, GlobalHandler},
	).SetTransitions(
		TransitionItem[Payload]{EventRepair, StateMaintenance, StateIdle, GlobalHandler},
	)

	var err error
	_, err = fsm.ProcessEvent(EventAddTask)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	_, err = fsm.ProcessEvent(EventDoneCmd)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	_, err = fsm.ProcessEvent(EventStartCmd)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	_, err = fsm.ProcessEvent(EventOperation)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	_, err = fsm.ProcessEvent(EventDoneCmd)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	_, err = fsm.ProcessEvent(EventMaintenance)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	_, err = fsm.ProcessEvent(EventRepair)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}

	fsm.VisualizeWithDetails("./fig/fsm-scheduler.png", fsm.DotString(), "png", "dot", "72", "-Gsize=20,20 -Gdpi=200")

	fsm.VisualizeMermaid("./fig/fsm-scheduler.md")
}

func TestSchedulerFSMConcurrency(t *testing.T) {
	fsm := NewFSM("testFSM", StateIdle, &Payload{
		Data: "scheduler",
	})

	fsm.SetTransitions(
		TransitionItem[Payload]{EventAddTask, StateIdle, StateAssigned, GlobalHandler},
	).SetTransitions(
		TransitionItem[Payload]{EventStartCmd, StateAssigned, StateProcessing, GlobalHandler},
		TransitionItem[Payload]{EventStartCmd, StateAssignedMaintenance, StateProcessingMaintenance, GlobalHandler},
	).SetTransitions(
		TransitionItem[Payload]{EventDoneCmd, StateProcessing, StateCompleted, GlobalHandler},
		TransitionItem[Payload]{EventDoneCmd, StateProcessingMaintenance, StateCompletedMaintenance, GlobalHandler},
	).SetTransitions(
		TransitionItem[Payload]{EventAbortCmd, StateAssigned, StateAbort, GlobalHandler},
		TransitionItem[Payload]{EventAbortCmd, StateProcessing, StateAbort, GlobalHandler},
		TransitionItem[Payload]{EventAbortCmd, StateAssignedMaintenance, StateAbortMaintenance, GlobalHandler},
		TransitionItem[Payload]{EventAbortCmd, StateProcessingMaintenance, StateAbortMaintenance, GlobalHandler},
	).SetTransitions(
		TransitionItem[Payload]{EventOperation, StateIdle, StateMaintenance, GlobalHandler},
		TransitionItem[Payload]{EventOperation, StateAssigned, StateAssignedMaintenance, GlobalHandler},
		TransitionItem[Payload]{EventOperation, StateProcessing, StateProcessingMaintenance, GlobalHandler},
	).SetTransitions(
		TransitionItem[Payload]{EventMaintenance, StateCompletedMaintenance, StateMaintenance, GlobalHandler},
		TransitionItem[Payload]{EventMaintenance, StateAbortMaintenance, StateMaintenance, GlobalHandler},
	).SetTransitions(
		TransitionItem[Payload]{EventClearTask, StateCompleted, StateIdle, GlobalHandler},
		TransitionItem[Payload]{EventClearTask, StateAbort, StateIdle, GlobalHandler},
	).SetTransitions(
		TransitionItem[Payload]{EventRepair, StateMaintenance, StateIdle, GlobalHandler},
	)

	var err error
	threading.GoSafe(func() {
		_, err = fsm.ProcessEvent(EventAddTask)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	})

	threading.GoSafe(func() {
		_, err = fsm.ProcessEvent(EventDoneCmd)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	})

	threading.GoSafe(func() {
		_, err = fsm.ProcessEvent(EventStartCmd)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	})

	threading.GoSafe(func() {
		_, err = fsm.ProcessEvent(EventOperation)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	})

	threading.GoSafe(func() {
		_, err = fsm.ProcessEvent(EventDoneCmd)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	})

	threading.GoSafe(func() {
		_, err = fsm.ProcessEvent(EventMaintenance)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	})

	threading.GoSafe(func() {
		_, err = fsm.ProcessEvent(EventRepair)
		if err != nil {
			fmt.Printf("err: %v\n", err)
		}
	})

	<-time.After(20 * time.Second)
}
