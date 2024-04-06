package crdt

import "time"

type LWWRegister[T any] struct {
	name      string
	value     T
	timestamp int64
}

func NewLWWRegister[T any](name string, value T) *LWWRegister[T] {
	reg := new(LWWRegister[T])
	reg.name = name
	reg.value = value
	reg.timestamp = time.Now().UnixNano()
	return reg
}

func (reg *LWWRegister[T]) Value() T {
	return reg.value
}

func (reg *LWWRegister[T]) Assign(value T) {
	reg.timestamp = time.Now().UnixNano()
	reg.value = value
}

func (reg *LWWRegister[T]) Merge(that *LWWRegister[T]) {
	if reg.timestamp < that.timestamp {
		reg.value = that.value
		reg.timestamp = that.timestamp
	}
}
