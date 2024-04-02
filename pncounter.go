package crdt

import "math"

type PNCounter struct {
	positives, negatives *GCounter
}

func NewPNCounter(value int, clock *VectorClock) *PNCounter {
	counter := new(PNCounter)
	counter.positives = NewGCounter(clock)
	counter.negatives = NewGCounter(clock)
	if value >= 0 {
		counter.positives.value = value
	} else {
		counter.negatives.value = int(math.Abs(float64(value)))
	}
	return counter
}

func NewEmptyPNCounter(clock *VectorClock) *PNCounter {
	return NewPNCounter(0, clock)
}

func (counter *PNCounter) Value() int {
	return int(counter.positives.value - counter.negatives.value)
}

func (counter *PNCounter) Increment() {
	counter.positives.value++
}

func (counter *PNCounter) Decrement() {
	counter.negatives.value++
}

func (counter *PNCounter) Merge(that *PNCounter) {
	if counter.positives.value < that.positives.value {
		counter.positives.value = that.positives.value
	}
	if counter.negatives.value < that.negatives.value {
		counter.negatives.value = that.negatives.value
	}
}
