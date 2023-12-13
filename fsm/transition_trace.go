package fsm

import "time"

type TransitionTrace struct {
	FromState        FSMState
	Event            FSMEvent
	ToState          FSMState
	TransitionResult bool
	FinalState       FSMState
	BeginTime        int64 // nanosecond timestamp
	EndTime          int64 // nanosecond timestamp
}

func NewTransitionTrace(fromState FSMState, event FSMEvent) *TransitionTrace {
	return &TransitionTrace{
		FromState: fromState,
		Event:     event,
	}
}

func (tt *TransitionTrace) Start() *TransitionTrace {
	tt.BeginTime = time.Now().UnixNano()

	return tt
}

func (tt *TransitionTrace) End() {
	tt.EndTime = time.Now().UnixNano()
}

func (tt *TransitionTrace) ReleaseControl() *TransitionTrace {
	return tt
}

func (tt *TransitionTrace) SetToState(toState FSMState) *TransitionTrace {
	tt.ToState = toState

	return tt
}

func (tt *TransitionTrace) SetFinalState(finalState FSMState) *TransitionTrace {
	tt.FinalState = finalState

	return tt
}

func (tt *TransitionTrace) TransitionSuccess() *TransitionTrace {
	tt.TransitionResult = true

	return tt
}

func (tt *TransitionTrace) TransitionFail() *TransitionTrace {
	tt.TransitionResult = false

	return tt
}

func (tt *TransitionTrace) TrimPtr() TransitionTrace {
	return *tt
}
