package crdt

import "time"

// LWWRegister is a CRDT register that holds the most recent value that was
// assigned to it.
// The integer `time.Now().UnixNano()` is used to identify a total ordering of
// assignments and is assumed to be universal across all replicas of this
// LWWRegister.
type LWWRegister[T any] struct {
	name      string
	value     T
	timestamp int64
}

// NewLWWRegister constructs a LWWRegister with an initial value of value.
// The name of the register is assumed to be unique across the cluster this
// register exists in.
func NewLWWRegister[T any](name string, value T) *LWWRegister[T] {
	reg := new(LWWRegister[T])
	reg.name = name
	reg.value = value
	reg.timestamp = time.Now().UnixNano()
	return reg
}

// Value returns the current value stored in the register.
func (reg *LWWRegister[T]) Value() T {
	return reg.value
}

// Assign sets the value of this register to the parameter and updates the
// timestamp to `time.Now().UnixNano()`.
func (reg *LWWRegister[T]) Assign(value T) {
	reg.timestamp = time.Now().UnixNano()
	reg.value = value
}

// Merge will assign the value and timetamp in that to reg if and only if
// reg.name == that.name AND reg.timestamp < that.timestamp.
func (reg *LWWRegister[T]) Merge(that *LWWRegister[T]) {
	if reg.name == that.name && reg.timestamp < that.timestamp {
		reg.value = that.value
		reg.timestamp = that.timestamp
	}
}
