package crdt

// PNCounter is a CRDT counter that can both increase and decrease in value.
type PNCounter struct {
	name                 string
	node                 string
	positives, negatives *GCounter
}

// NewPNCounter constructs a new PNCounter with the associated name and an
// initial value of 0.
// It is assumed the name of this specific PNCounter uniquely identifies this
// counter throughout the cluster.
func NewPNCounter(name string, node string) *PNCounter {
	counter := new(PNCounter)
	counter.name = name
	counter.node = node
	counter.positives = NewGCounter(name+"_positives", node)
	counter.negatives = NewGCounter(name+"_negatives", node)
	return counter
}

// Value returns the current value of this counter.
func (counter *PNCounter) Value() int {
	return counter.positives.Value() - counter.negatives.Value()
}

// Increment increases the value of this counter by 1.
func (counter *PNCounter) Increment() {
	counter.positives.values[counter.node]++
}

// Decrement decreases the value of this counter by 1.
func (counter *PNCounter) Decrement() {
	counter.negatives.values[counter.node]++
}

// Merge performs an idempotent operation of which each GCounter in counter is
// merged with its counterpart in that.
// This is an idempotent operation and is a no-op if counter.name != that.name.
func (counter *PNCounter) Merge(that *PNCounter) {
	if counter.name != that.name {
		return
	}
	counter.positives.Merge(that.positives)
	counter.negatives.Merge(that.negatives)
}
