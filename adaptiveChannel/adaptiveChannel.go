package adaptiveChannel

import "sync/atomic"

const (
	InitialAdaptiveChannelSize = 1 << 8
)

type adaptiveChannel[T any] struct {
	bufferElementNum int64
	InChan           chan<- T
	OutChan          <-chan T
	ringBufferLink   *RingBufferLink[T]
}

func (a *adaptiveChannel[T]) GetAdaptiveChannelLen() int64 {
	return int64(len(a.InChan)+len(a.OutChan)) + a.GetBufferElementNum()
}

func (a *adaptiveChannel[T]) GetBufferElementNum() int64 {
	return atomic.LoadInt64(&a.bufferElementNum)
}

func (a *adaptiveChannel[T]) GetBufferNodeNum() int {
	return a.ringBufferLink.GetNodeNum()
}

func (a *adaptiveChannel[T]) GetBufferNodeElementListLen() int {
	return a.ringBufferLink.nodeElementListLen
}

func (a *adaptiveChannel[T]) Write(element T) {
	a.ringBufferLink.WriteElement(element)
	atomic.AddInt64(&a.bufferElementNum, 1)
}

func (a *adaptiveChannel[T]) Pop() {
	a.ringBufferLink.PopElement()
	atomic.AddInt64(&a.bufferElementNum, -1)
}

func (a *adaptiveChannel[T]) Pop2OutChan(outChan chan T) {
	outChan <- a.ringBufferLink.PopElement()
	atomic.AddInt64(&a.bufferElementNum, -1)
}

func (a *adaptiveChannel[T]) ResetBufferElementNum() {
	atomic.StoreInt64(&a.bufferElementNum, 0)
}

// NewAdaptiveChannel
//	- sizeArgs ...int:
//		1. if len(sizeArgs)==0:
//			- inChanSize=outChanSize=InitialAdaptiveChannelSize=1<<8
//			- bufferListSize=InitialElementListSize=1<<5
//		2. if len(sizeArgs)==1:
//			- inChanSize=outChanSize=sizeArgs[0]
//			- bufferListSize=InitialElementListSize=1<<5
//		3. if len(sizeArgs)==2:
//			- inChanSize=sizeArgs[0]
//			- outChanSize=sizeArgs[1]
//			- bufferListSize=InitialElementListSize=1<<5
//		4. if len(sizeArgs)>2:
//			- inChanSize=sizeArgs[0]
//			- outChanSize=sizeArgs[1]
//			- bufferListSize=sizeArgs[2]
func NewAdaptiveChannel[T any](sizeArgs ...int) *adaptiveChannel[T] {
	var (
		inChan         chan T
		outChan        chan T
		inChanSize     int
		outChanSize    int
		bufferListSize int = 0
	)
	if len(sizeArgs) == 0 {
		inChanSize = InitialAdaptiveChannelSize
		outChanSize = InitialAdaptiveChannelSize
	} else if len(sizeArgs) == 1 {
		inChanSize = sizeArgs[0]
		outChanSize = sizeArgs[0]
	} else if len(sizeArgs) == 2 {
		inChanSize = sizeArgs[0]
		outChanSize = sizeArgs[1]
	} else {
		inChanSize = sizeArgs[0]
		outChanSize = sizeArgs[1]
		bufferListSize = sizeArgs[2]
	}
	inChan = make(chan T, inChanSize)
	outChan = make(chan T, outChanSize)
	ringBuffer := NewRingBufferLink[T](bufferListSize)

	ac := &adaptiveChannel[T]{
		InChan:         inChan,
		OutChan:        outChan,
		ringBufferLink: ringBuffer,
	}

	go dealAdaptiveChannel(inChan, outChan, ac)

	return ac
}

func dealAdaptiveChannel[T any](inChan chan T, outChan chan T, ac *adaptiveChannel[T]) {
	defer close(outChan)
loop:
	for {
		element, ok := <-inChan
		if !ok {
			break
		}

		if ac.GetBufferElementNum() > 0 {
			ac.Write(element)
		} else {
			select {
			case outChan <- element:
				continue
			default:
			}

			ac.Write(element)
		}

		for !ac.ringBufferLink.IsEmpty() {
			select {
			case elem, ok := <-inChan:
				if !ok {
					break loop
				}

				ac.Write(elem)
			case outChan <- ac.ringBufferLink.GetPeekElement():
				ac.Pop()
				if ac.ringBufferLink.IsEmpty() {
					ac.ringBufferLink.Reset()
					ac.ResetBufferElementNum()
				}
			}

		}
	}

	for !ac.ringBufferLink.IsEmpty() {
		ac.Pop2OutChan(outChan)
	}
	ac.ringBufferLink.Reset()
	ac.ResetBufferElementNum()
}
