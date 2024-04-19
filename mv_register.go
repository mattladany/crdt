package crdt

import "time"

// MVRegister is a CRDT register that holds a registerValue for each node that
// exists within the cluster.
// TODO the values of each register here could be a LWWRegister to avoid
// duplicating logic.
type MVRegister[T any] struct {
	node      string
	registers map[string]registerValue[T]
}

// registerValue represents the register of a particular node within the
// cluster.
type registerValue[T any] struct {
	value     T
	timestamp int64
}

// NewMVRegister constructs a MVRegister with a current state of the known
// cluster's register values.
func NewMVRegister[T any](node string, nodes map[string]registerValue[T]) *MVRegister[T] {
	mvRegister := new(MVRegister[T])
	mvRegister.node = node
	mvRegister.registers = nodes
	return mvRegister
}

// Assign sets the value in the parameter to the value of the register for
// this node.
func (reg *MVRegister[T]) Assign(value T) {
	singleReg := new(registerValue[T])
	singleReg.value = value
	singleReg.timestamp = time.Now().UnixNano()
	reg.registers[reg.node] = *singleReg
}

// Value returns a map of node -> value for each node in the cluster.
func (reg *MVRegister[T]) Value() map[string]T {
	values := make(map[string]T)
	for node, register := range reg.registers {
		values[node] = register.value
	}
	return values
}

// Merge sets the value of each node's register using a last-writer-wins
// implementation.
func (reg *MVRegister[T]) Merge(that *MVRegister[T]) {
	for node, subReg := range that.registers {
		if _, exists := reg.registers[node]; exists {
			if reg.registers[node].timestamp < subReg.timestamp {
				reg.registers[node] = subReg
			}
		} else {
			reg.registers[node] = subReg
		}
	}
}
