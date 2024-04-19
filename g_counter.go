package crdt

// GCounter is a CRDT counter that can only be incremented.
type GCounter struct {
	name   string
	node   string
	values map[string]int
}

// NewGCounter constructs a GCounter based on the parameters provided.
// The GCounter will be initialized to a value of 0.
// It is assumed the name of this specific GCounter uniquely identifies this
// counter throughout the cluster.
func NewGCounter(name string, node string, nodes []string) *GCounter {
	counter := new(GCounter)
	counter.name = name
	counter.node = node
	counter.values = make(map[string]int)
	for _, n := range nodes {
		counter.values[n] = 0
	}
	return counter
}

// Value returns the current value of this counter.
func (counter *GCounter) Value() int {
	sum := 0
	for _, val := range counter.values {
		sum += val
	}
	return sum
}

// Increment increments this counter by 1.
func (counter *GCounter) Increment() {
	counter.values[counter.node]++
}

// Merge performs an idempotent operation which combines that with counter.
// The maximum value for each value in values is set to each counter.values.
// This is an idempotent operation and is a no-op if counter.name != that.name.
func (counter *GCounter) Merge(that *GCounter) {
	if counter.name != that.name {
		return
	}
	for key, thatValue := range that.values {
		if counterValue, exists := counter.values[key]; exists {
			if counterValue < thatValue {
				counter.values[key] = thatValue
			}
		} else {
			counter.values[key] = thatValue
		}
	}
}
