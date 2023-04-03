package adaptiveChannel

import (
	"errors"
)

const (
	InitialNodeSize        = 1 << 1
	InitialElementListSize = 1 << 5
)

type node[T any] struct {
	ElementList     []T
	isFull          bool
	preNode         *node[T]
	nextNode        *node[T]
	readElementPtr  int
	writeElementPtr int
}

func newNode[T any](listSize int) *node[T] {
	return &node[T]{
		ElementList: make([]T, listSize),
	}
}

type RingBufferLink[T any] struct {
	nodeNum            int
	nodeElementListLen int
	readNode           *node[T]
	writeNode          *node[T]
}

func NewRingBufferLink[T any](listSize ...int) *RingBufferLink[T] {
	var elementListSize = InitialElementListSize
	if len(listSize) > 0 && listSize[0] > 0 {
		elementListSize = listSize[0]
	}
	rootNode, lastNode := newNode[T](elementListSize), newNode[T](elementListSize)

	rootNode.nextNode = lastNode
	lastNode.preNode = rootNode

	rootNode.preNode = lastNode
	lastNode.nextNode = rootNode

	return &RingBufferLink[T]{
		nodeNum:            InitialNodeSize,
		nodeElementListLen: elementListSize,
		readNode:           rootNode,
		writeNode:          rootNode,
	}
}

func (r *RingBufferLink[T]) GetNodeNum() int {
	return r.nodeNum
}

func (r *RingBufferLink[T]) IsEmpty() bool {
	if r.readNode == r.writeNode && r.readNode.readElementPtr == r.readNode.writeElementPtr && !r.readNode.isFull {
		return true
	}

	return false
}

func (r *RingBufferLink[T]) GetPeekElement() T {
	if r.IsEmpty() {
		panic("ring buffer link is empty")
	}

	return r.readNode.ElementList[r.readNode.readElementPtr]
}

func (r *RingBufferLink[T]) ReadElement() (T, error) {
	var res T
	if r.IsEmpty() {
		return res, errors.New("ring buffer link is empty")
	}

	value := r.readNode.ElementList[r.readNode.readElementPtr]
	r.readNode.readElementPtr++

	if r.readNode.readElementPtr == InitialElementListSize {
		r.readNode.readElementPtr = 0
		r.readNode.isFull = false
		r.readNode = r.readNode.nextNode
	}

	return value, nil
}

func (r *RingBufferLink[T]) PopElement() T {
	value, err := r.ReadElement()

	if err != nil {
		panic(err.Error())
	}

	return value
}

func (r *RingBufferLink[T]) WriteElement(element T) {
	r.writeNode.ElementList[r.writeNode.writeElementPtr] = element
	r.writeNode.writeElementPtr++

	if r.writeNode.writeElementPtr == InitialElementListSize {
		r.writeNode.writeElementPtr = 0
		r.writeNode.isFull = true
		r.writeNode = r.writeNode.nextNode
	}

	if r.writeNode.isFull {
		r.adaptiveNodeNum()
	}
}

// adaptiveNodeNum automatic expansion
func (r *RingBufferLink[T]) adaptiveNodeNum() {
	node := newNode[T](r.nodeElementListLen)

	// preNode <--> pre/next <--> writeNode <--> pre/next <--> preNode
	// preNode <--> pre/next node <--> pre/next <--> writeNode <--> pre/next <--> preNode
	preNode := r.writeNode.preNode

	preNode.nextNode = node
	node.preNode = preNode

	node.nextNode = r.writeNode
	r.writeNode.preNode = node

	r.writeNode = r.writeNode.preNode

	r.nodeNum++
}

func (r *RingBufferLink[T]) Capacity() int {
	return r.nodeNum * r.nodeElementListLen
}

// Reset set RingBufferLink to origin status
func (r *RingBufferLink[T]) Reset() {
	root, lastNode := r.readNode, r.readNode.nextNode

	root.readElementPtr, root.writeElementPtr = 0, 0
	lastNode.readElementPtr, lastNode.writeElementPtr = 0, 0

	r.nodeNum = InitialNodeSize
	lastNode.nextNode = root
}
