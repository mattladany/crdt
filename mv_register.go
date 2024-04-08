package crdt

import "time"

type registerValue[T any] struct {
	value     T
	timestamp int64
}

type MVRegister[T any] struct {
	node      string
	registers map[string]registerValue[T]
}

func NewMVRegister[T any](node string, nodes map[string]registerValue[T]) *MVRegister[T] {
	mvRegister := new(MVRegister[T])
	mvRegister.node = node
	mvRegister.registers = nodes
	return mvRegister
}

func (reg *MVRegister[T]) Assign(value T) {
	singleReg := new(registerValue[T])
	singleReg.value = value
	singleReg.timestamp = time.Now().UnixNano()
	reg.registers[reg.node] = *singleReg
}

func (reg *MVRegister[T]) Value() map[string]T {
	values := make(map[string]T)
	for node, register := range reg.registers {
		values[node] = register.value
	}
	return values
}

func (reg *MVRegister[T]) Merge(that *MVRegister[T]) {
	for node, subReg := range that.registers {
		if reg.registers[node].timestamp < subReg.timestamp {
			reg.registers[node] = subReg
		}
	}
}
