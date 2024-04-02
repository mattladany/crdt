package crdt

import "testing"

func initPNCounterClock(name string) *VectorClock {
	clocks := make(map[string]int)
	clocks["srv1"] = 0
	clocks["srv2"] = 0
	clocks["srv3"] = 0
	clocks["srv4"] = 0
	return NewVectorClock(name, clocks)
}

func TestPNCounterEmptyInitialization(t *testing.T) {
	counter := NewEmptyPNCounter(initPNCounterClock("srv1"))
	if counter.Value() != 0 {
		t.Fatalf("counter value should initialize to 0")
	}
}

func TestPNCounterExplicitInitialization(t *testing.T) {
	counter := NewPNCounter(5, initPNCounterClock("srv1"))
	if counter.Value() != 5 {
		t.Fatalf("counter value should initialize to 5")
	}
}

func TestPNCounterSingleIncrement(t *testing.T) {
	counter := NewEmptyPNCounter(initPNCounterClock("srv1"))
	counter.Increment()
	if counter.Value() != 1 {
		t.Fatalf("counter value should be 1")
	}
}

func TestPNCounterMultiIncrement(t *testing.T) {
	counter := NewEmptyPNCounter(initPNCounterClock("srv1"))
	counter.Increment()
	counter.Increment()
	counter.Increment()
	counter.Increment()
	if counter.Value() != 4 {
		t.Fatalf("counter value should be 4")
	}
}

func TestPNCounterSingleDecrement(t *testing.T) {
	counter := NewEmptyPNCounter(initPNCounterClock("srv1"))
	counter.Decrement()
	if counter.Value() != -1 {
		t.Fatalf("counter value should be -1")
	}
}

func TestPNCounterMultiDecrement(t *testing.T) {
	counter := NewEmptyPNCounter(initPNCounterClock("srv1"))
	counter.Decrement()
	counter.Decrement()
	counter.Decrement()
	counter.Decrement()
	if counter.Value() != -4 {
		t.Fatalf("counter value should be -4")
	}
}

func TestPNCounterEqualIncrementAndDecrement(t *testing.T) {
	counter := NewEmptyPNCounter(initPNCounterClock("srv1"))
	counter.Increment()
	counter.Increment()
	counter.Increment()
	counter.Increment()
	counter.Decrement()
	counter.Decrement()
	counter.Decrement()
	counter.Decrement()
	if counter.Value() != 0 {
		t.Fatalf("counter value should be 0")
	}
}

func TestPNCounterPositiveIncrementAndDecrement(t *testing.T) {
	counter := NewEmptyPNCounter(initPNCounterClock("srv1"))
	counter.Increment()
	counter.Increment()
	counter.Increment()
	counter.Increment()
	counter.Decrement()
	counter.Decrement()
	if counter.positives.value != 4 {
		t.Fatalf("counter.positives.value should be 4")
	}
	if counter.negatives.value != 2 {
		t.Fatalf("counter.positives.value should be 2")
	}
	if counter.Value() != 2 {
		t.Fatalf("counter value should be 2")
	}
}

func TestPNCounterNegativeIncrementAndDecrement(t *testing.T) {
	counter := NewEmptyPNCounter(initPNCounterClock("srv1"))
	counter.Increment()
	counter.Increment()
	counter.Decrement()
	counter.Decrement()
	counter.Decrement()
	counter.Decrement()
	if counter.positives.value != 2 {
		t.Fatalf("counter.positives.value should be 2")
	}
	if counter.negatives.value != 4 {
		t.Fatalf("counter.positives.value should be 4")
	}
	if counter.Value() != -2 {
		t.Fatalf("counter value should be -2")
	}
}

func TestPNCounterMergeIdempotent(t *testing.T) {
	counter1 := NewEmptyPNCounter(initPNCounterClock("srv1"))
	counter1.Increment()
	counter1.Increment()
	counter1.Decrement()
	counter1.Decrement()
	counter2 := NewEmptyPNCounter(initPNCounterClock("srv2"))
	counter2.Increment()
	counter2.Decrement()

	counter1.Merge(counter2)

	if counter1.positives.value != 2 {
		t.Fatalf("counter1.positives.value should be 2")
	}
	if counter1.negatives.value != 2 {
		t.Fatalf("counter1.positives.value should be 2")
	}
	if counter1.Value() != 0 {
		t.Fatalf("counter value should be 0")
	}
}

func TestPNCounterMergePositiveChange(t *testing.T) {
	counter1 := NewEmptyPNCounter(initPNCounterClock("srv1"))
	counter1.Increment()
	counter1.Decrement()
	counter1.Decrement()
	counter2 := NewEmptyPNCounter(initPNCounterClock("srv2"))
	counter2.Increment()
	counter2.Increment()
	counter2.Decrement()

	counter1.Merge(counter2)

	if counter1.positives.value != 2 {
		t.Fatalf("counter1.positives.value should be 2")
	}
	if counter1.negatives.value != 2 {
		t.Fatalf("counter1.positives.value should be 2")
	}
	if counter1.Value() != 0 {
		t.Fatalf("counter value should be 0")
	}
}

func TestPNCounterMergeNegativeChange(t *testing.T) {
	counter1 := NewEmptyPNCounter(initPNCounterClock("srv1"))
	counter1.Increment()
	counter1.Increment()
	counter1.Decrement()
	counter2 := NewEmptyPNCounter(initPNCounterClock("srv2"))
	counter2.Increment()
	counter2.Decrement()
	counter2.Decrement()

	counter1.Merge(counter2)

	if counter1.positives.value != 2 {
		t.Fatalf("counter1.positives.value should be 2")
	}
	if counter1.negatives.value != 2 {
		t.Fatalf("counter1.positives.value should be 2")
	}
	if counter1.Value() != 0 {
		t.Fatalf("counter value should be 0")
	}
}

func TestPNCounterMergeFullChange(t *testing.T) {
	counter1 := NewEmptyPNCounter(initPNCounterClock("srv1"))
	counter1.Increment()
	counter1.Decrement()
	counter2 := NewEmptyPNCounter(initPNCounterClock("srv2"))
	counter2.Increment()
	counter2.Increment()
	counter2.Decrement()
	counter2.Decrement()

	counter1.Merge(counter2)

	if counter1.positives.value != 2 {
		t.Fatalf("counter1.positives.value should be 2")
	}
	if counter1.negatives.value != 2 {
		t.Fatalf("counter1.positives.value should be 2")
	}
	if counter1.Value() != 0 {
		t.Fatalf("counter value should be 0")
	}
}
