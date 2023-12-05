package test

type OrderedMap[KeyType comparable, ValueType any] struct {
	_Keys []KeyType
	_Map  map[KeyType]ValueType
}

func NewOrderedMap[KeyType comparable, ValueType any]() OrderedMap[KeyType, ValueType] {
	return OrderedMap[KeyType, ValueType]{
		_Keys: []KeyType{},
		_Map:  map[KeyType]ValueType{},
	}
}

func (om *OrderedMap[KeyType, ValueType]) Empty() bool {
	return len(om._Keys) == 0 || len(om._Map) == 0
}

func (om *OrderedMap[KeyType, ValueType]) Len() int {
	return len(om._Keys)
}

func (om *OrderedMap[KeyType, ValueType]) Keys() []KeyType {
	return om._Keys
}

func (om *OrderedMap[KeyType, ValueType]) Exist(key KeyType) bool {
	_, ok := om._Map[key]
	return ok
}

func (om *OrderedMap[KeyType, ValueType]) Value(key KeyType) ValueType {
	return om._Map[key]
}

func (om *OrderedMap[KeyType, ValueType]) Add(key KeyType, value ValueType) {
	om._Map[key] = value
}

func (om *OrderedMap[KeyType, ValueType]) Delete(key KeyType) {
	delete(om._Map, key)
}
