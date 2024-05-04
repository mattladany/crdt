package crdt

// MVRegister is a CRDT register that holds a registerValue for each node that
// exists within the cluster.
// TODO the values of each register here could be a LWWRegister to avoid
// duplicating logic.
type MVRegister[T any] struct {
	name      string
	node      string
	registers map[string]*LWWRegister[T]
}

// NewMVRegister constructs a MVRegister with a value for this node set to initialValue.
// It is assumed the name of this specific MVRegister uniquely identifies this
// register throughout the cluster.
func NewMVRegister[T any](name string, node string, initialValue T) *MVRegister[T] {
	mvRegister := new(MVRegister[T])
	mvRegister.name = name
	mvRegister.node = node
	mvRegister.registers = make(map[string]*LWWRegister[T])

	mvRegister.registers[node] = NewLWWRegister(name, initialValue)

	return mvRegister
}

// Assign sets the value in the parameter to the value of the register for
// this node.
func (reg *MVRegister[T]) Assign(value T) {
	reg.registers[reg.node] = NewLWWRegister(reg.name, value)
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
// This is an idempotent operation and is a no-op if reg.name != that.name.
func (reg *MVRegister[T]) Merge(that *MVRegister[T]) {
	if reg.name != that.name {
		return
	}
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
