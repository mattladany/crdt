package crdt

type MVRegister[T any] struct {
	clock  *VectorClock
	values []T
}

func NewMVRegister[T any](name string, value T, clock *VectorClock) *MVRegister[T] {
	reg := new(MVRegister[T])
	reg.clock = clock
	reg.values = []T{value}
	return reg
}

func (reg *MVRegister[T]) Get() []T {
	return reg.values
}

func (reg *MVRegister[T]) Set(value T) {
	reg.clock.inc()
	reg.values = []T{value}
}

func (reg *MVRegister[T]) Merge(that *MVRegister[T]) {
	if reg.clock.merge(that.clock) {
		// reg.values
	}
}
