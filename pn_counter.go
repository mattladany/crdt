package crdt

// PNCounter is a CvRDT counter that can both increase and decrease in value.
type PNCounter struct {
	name                 string
	node                 string
	positives, negatives *GCounter
}

// NewPNCounter constructs a new PNCounter with an initial value of 0.
func NewPNCounter(name string, node string, nodes []string) *PNCounter {
	counter := new(PNCounter)
	counter.name = name
	counter.node = node
	counter.positives = NewGCounter(name+"_positives", node, nodes)
	counter.negatives = NewGCounter(name+"_negatives", node, nodes)
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
// This operation will be ignored if that.name != counter.name.
func (counter *PNCounter) Merge(that *PNCounter) {
	if counter.name != that.name {
		return
	}
	counter.positives.Merge(that.positives)
	counter.negatives.Merge(that.negatives)
}
