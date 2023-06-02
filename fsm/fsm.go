package fsm

import (
	"fmt"
	"sync"
	"time"
)

type (
	FSMName  string
	FSMState string
	FSMEvent string
)

type EventStateKey struct {
	Event     FSMEvent
	FromState FSMState
}

func (esk *EventStateKey) String() string {
	return string(esk.Event) + "_" + string(esk.FromState)
}

var (
	EventEmpty                = FSMEvent("")
	StateEmpty                = FSMState("")
	ErrorEventStateKeyInvalid = func(eventStateKey EventStateKey) error {
		return fmt.Errorf("event state key %v not exist", eventStateKey.String())
	}
	ErrorEventStateKeyRepeated = func(eventStateKey EventStateKey) error {
		return fmt.Errorf("event state key %v is repeated", eventStateKey.String())
	}
)

type TransitionItem[T any] struct {
	Event     FSMEvent
	FromState FSMState
	ToState   FSMState
	Handler   Handler[T]
}

type TransitionTrace struct {
	FromState        FSMState
	Event            FSMEvent
	ToState          FSMState
	TransitionResult bool
	FinalState       FSMState
	BeginTime        int64 // nanosecond timestamp
	EndTime          int64 // nanosecond timestamp
}

func NewTransitionItem[T any](event FSMEvent, fromState FSMState, toState FSMState, handler Handler[T]) TransitionItem[T] {
	item := TransitionItem[T]{
		Event:     event,
		FromState: fromState,
		ToState:   toState,
		Handler:   handler,
	}

	return item
}

type (
	Actor[T any] func(*FSM[T]) error

	Handler[T any] struct {
		// BeforeHandler check state etc.
		//	- can not change the current state
		// 	- if check fail, should exit all handler function
		BeforeHandler Actor[T]

		// Handler real logic etc.
		// 	- finish state change logic
		// 	- if logic abort, should recover the data of ahead exec Handler, then exit
		Handler Actor[T]

		// AfterHandler save state, clear data etc.
		AfterHandler Actor[T]
	}

	FSM[T any] struct {
		Name            FSMName
		fromState       FSMState
		processingEvent FSMEvent
		currentState    FSMState
		// map[EventFromStateKey]ToState, FromState --Event-> ToState
		transitions map[EventStateKey]FSMState
		// map[EventFromStateKey]Handler, FromState --Handler(Event)-> ToState
		handlers  map[EventStateKey]Handler[T]
		Payload   *T
		stateLock sync.RWMutex
		eventLock sync.Mutex
	}
)

func NewFSM[T any](fsmName FSMName, initState FSMState, payloads ...*T) *FSM[T] {
	fsm := &FSM[T]{
		Name:            fsmName,
		fromState:       StateEmpty,
		currentState:    initState,
		processingEvent: EventEmpty,
		transitions:     make(map[EventStateKey]FSMState),
		handlers:        make(map[EventStateKey]Handler[T]),
		Payload: func() *T {
			if len(payloads) > 0 {
				return payloads[0]
			}
			return nil
		}(),
	}

	return fsm
}

func (fsm *FSM[T]) SetTransitions(transitionItems ...TransitionItem[T]) *FSM[T] {
	for _, transitionItem := range transitionItems {
		eventStateKey := EventStateKey{
			Event:     transitionItem.Event,
			FromState: transitionItem.FromState,
		}
		if fsm.ExistEventStateKey(eventStateKey) {
			panic(ErrorEventStateKeyRepeated)
		}
		fsm.transitions[eventStateKey] = transitionItem.ToState
		fsm.handlers[eventStateKey] = transitionItem.Handler
	}

	return fsm
}

func (fsm *FSM[T]) SetPayload(transitionItems ...TransitionItem[T]) *FSM[T] {
	for _, transitionItem := range transitionItems {
		eventStateKey := EventStateKey{
			Event:     transitionItem.Event,
			FromState: transitionItem.FromState,
		}
		if fsm.ExistEventStateKey(eventStateKey) {
			panic(ErrorEventStateKeyRepeated(eventStateKey))
		}
		fsm.transitions[eventStateKey] = transitionItem.ToState
		fsm.handlers[eventStateKey] = transitionItem.Handler
	}

	return fsm
}

func (fsm *FSM[T]) SetCurrentState(state FSMState) {
	fsm.stateLock.Lock()
	defer fsm.stateLock.Unlock()

	fsm.currentState = state
}

func (fsm *FSM[T]) GetCurrentState() FSMState {
	fsm.stateLock.RLock()
	defer fsm.stateLock.RUnlock()

	return fsm.currentState
}

func (fsm *FSM[T]) ExistEventStateKey(eventStateKey EventStateKey) bool {
	_, ok := fsm.transitions[eventStateKey]

	return ok
}

func (fsm *FSM[T]) ProcessEvent(event FSMEvent) (TransitionTrace, error) {
	fsm.eventLock.Lock()
	defer fsm.eventLock.Unlock()

	fsm.stateLock.RLock()
	defer fsm.stateLock.RUnlock()

	fsm.processingEvent = EventEmpty

	transTrace := TransitionTrace{
		FromState:        fsm.currentState,
		Event:            event,
		BeginTime:        time.Now().UnixNano(),
		TransitionResult: false,
	}

	eventStateKey := EventStateKey{
		Event:     event,
		FromState: fsm.currentState,
	}

	if !fsm.ExistEventStateKey(eventStateKey) {
		transTrace.FinalState = fsm.currentState
		transTrace.EndTime = time.Now().UnixNano()
		return transTrace, ErrorEventStateKeyInvalid(eventStateKey)
	}

	toState, handler := fsm.transitions[eventStateKey], fsm.handlers[eventStateKey]
	transTrace.ToState = toState

	if handler.BeforeHandler != nil {
		err := handler.BeforeHandler(fsm)
		if err != nil {
			transTrace.FinalState = fsm.currentState
			transTrace.EndTime = time.Now().UnixNano()
			return transTrace, err
		}
	}

	if handler.Handler != nil {
		err := handler.Handler(fsm)
		if err != nil {
			transTrace.FinalState = fsm.currentState
			transTrace.EndTime = time.Now().UnixNano()
			return transTrace, err
		}
	}

	if handler.AfterHandler != nil {
		err := handler.AfterHandler(fsm)
		if err != nil {
			transTrace.EndTime = time.Now().UnixNano()
			return transTrace, err
		}
	}

	transTrace.FinalState = toState
	transTrace.TransitionResult = true
	transTrace.EndTime = time.Now().UnixNano()

	fsm.fromState = fsm.currentState
	fsm.processingEvent = EventEmpty

	fsm.currentState = toState

	return transTrace, nil
}
