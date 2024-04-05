package crdt

type PNCounter struct {
	node                 string
	positives, negatives *GCounter
}

func NewPNCounter(node string, nodes []string) *PNCounter {
	counter := new(PNCounter)
	counter.node = node
	counter.positives = NewGCounter(node, nodes)
	counter.negatives = NewGCounter(node, nodes)
	return counter
}

func (counter *PNCounter) Value() int {
	return counter.positives.Value() - counter.negatives.Value()
}

func (counter *PNCounter) Increment() {
	counter.positives.values[counter.node]++
}

func (counter *PNCounter) Decrement() {
	counter.negatives.values[counter.node]++
}

func (counter *PNCounter) Merge(that *PNCounter) {
	counter.positives.Merge(that.positives)
	counter.negatives.Merge(that.negatives)
}
