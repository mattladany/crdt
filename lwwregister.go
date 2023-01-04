package crdt

type LWWRegister[T any] struct {
	clock int
	value T
}

func NewLWWRegister[T any](value T) *LWWRegister[T] {
	reg := new(LWWRegister[T])
	reg.clock = 0
	reg.value = value
	return reg
}

func (reg *LWWRegister[T]) Get() T {
	return reg.value
}

func (reg *LWWRegister[T]) Set(value T) {
	reg.clock++
	reg.value = value
}

func (reg *LWWRegister[T]) Merge(that *LWWRegister[T]) {
	if reg.clock < that.clock {
		reg.clock = that.clock
		reg.value = that.value
	}
}
