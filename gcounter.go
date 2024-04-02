package crdt

type GCounter struct {
	value int
	clock *VectorClock
}

func NewGCounter(clock *VectorClock) *GCounter {
	counter := new(GCounter)
	counter.value = 0
	counter.clock = clock
	return counter
}

func (counter *GCounter) Value() int {
	return counter.value
}

func (counter *GCounter) Increment() {
	counter.value++
	counter.clock.inc()
}

func (counter *GCounter) Merge(that *GCounter) {
	if counter.value < that.value {
		counter.value = that.value
	}
	counter.clock.merge(that.clock)
}
