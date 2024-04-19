package crdt

import "testing"

func initPNNodes() []string {
	return []string{"srv1", "srv2", "srv3", "srv4"}
}

func TestPNCounterEmptyInitialization(t *testing.T) {
	counter := NewPNCounter("counter1", "srv1", initPNNodes())
	if counter.Value() != 0 {
		t.Fatalf("counter value should initialize to 0")
	}
}

func TestPNCounterSingleIncrement(t *testing.T) {
	counter := NewPNCounter("counter1", "srv1", initPNNodes())
	counter.Increment()
	if counter.Value() != 1 {
		t.Fatalf("counter value should be 1")
	}
}

func TestPNCounterMultiIncrement(t *testing.T) {
	counter := NewPNCounter("counter1", "srv1", initPNNodes())
	counter.Increment()
	counter.Increment()
	counter.Increment()
	counter.Increment()
	if counter.Value() != 4 {
		t.Fatalf("counter value should be 4")
	}
}

func TestPNCounterSingleDecrement(t *testing.T) {
	counter := NewPNCounter("counter1", "srv1", initPNNodes())
	counter.Decrement()
	if counter.Value() != -1 {
		t.Fatalf("counter value should be -1")
	}
}

func TestPNCounterMultiDecrement(t *testing.T) {
	counter := NewPNCounter("counter1", "srv1", initPNNodes())
	counter.Decrement()
	counter.Decrement()
	counter.Decrement()
	counter.Decrement()
	if counter.Value() != -4 {
		t.Fatalf("counter value should be -4")
	}
}

func TestPNCounterEqualIncrementAndDecrement(t *testing.T) {
	counter := NewPNCounter("counter1", "srv1", initPNNodes())
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
	counter := NewPNCounter("counter1", "srv1", initPNNodes())
	counter.Increment()
	counter.Increment()
	counter.Increment()
	counter.Increment()
	counter.Decrement()
	counter.Decrement()
	if counter.positives.Value() != 4 {
		t.Fatalf("counter.positives.value should be 4")
	}
	if counter.negatives.Value() != 2 {
		t.Fatalf("counter.positives.value should be 2")
	}
	if counter.Value() != 2 {
		t.Fatalf("counter value should be 2")
	}
}

func TestPNCounterNegativeIncrementAndDecrement(t *testing.T) {
	counter := NewPNCounter("counter1", "srv1", initPNNodes())
	counter.Increment()
	counter.Increment()
	counter.Decrement()
	counter.Decrement()
	counter.Decrement()
	counter.Decrement()
	if counter.positives.Value() != 2 {
		t.Fatalf("counter.positives.value should be 2")
	}
	if counter.negatives.Value() != 4 {
		t.Fatalf("counter.positives.value should be 4")
	}
	if counter.Value() != -2 {
		t.Fatalf("counter value should be -2")
	}
}

func TestPNCounterMergeStatic(t *testing.T) {
	counter1 := NewPNCounter("counter1", "srv1", initPNNodes())
	counter1.Increment()
	counter1.Increment()
	counter1.Decrement()
	counter1.Decrement()
	counter2 := NewPNCounter("counter1", "srv2", initPNNodes())
	counter2.Increment()
	counter2.Decrement()

	counter1.Merge(counter2)

	if counter1.positives.Value() != 3 {
		t.Fatalf("counter1.positives.value should be 2")
	}
	if counter1.negatives.Value() != 3 {
		t.Fatalf("counter1.positives.value should be 2")
	}
	if counter1.Value() != 0 {
		t.Fatalf("counter value should be 0")
	}
}

func TestPNCounterMergeIncrease(t *testing.T) {
	counter1 := NewPNCounter("counter1", "srv1", initPNNodes())
	counter1.Increment()
	counter1.Increment()
	counter1.Decrement()
	counter1.Decrement()
	counter2 := NewPNCounter("counter1", "srv2", initPNNodes())
	counter2.Increment()
	counter2.Increment()

	counter1.Merge(counter2)

	if counter1.positives.Value() != 4 {
		t.Fatalf("counter1.positives.value should be 2")
	}
	if counter1.negatives.Value() != 2 {
		t.Fatalf("counter1.positives.value should be 2")
	}
	if counter1.Value() != 2 {
		t.Fatalf("counter value should be 2")
	}
}

func TestPNCounterMergeDecrease(t *testing.T) {
	counter1 := NewPNCounter("counter1", "srv1", initPNNodes())
	counter1.Increment()
	counter1.Decrement()
	counter1.Decrement()
	counter2 := NewPNCounter("counter1", "srv2", initPNNodes())
	counter2.Decrement()
	counter2.Decrement()
	counter2.Decrement()

	counter1.Merge(counter2)

	if counter1.positives.Value() != 1 {
		t.Fatalf("counter1.positives.value should be 1")
	}
	if counter1.negatives.Value() != 5 {
		t.Fatalf("counter1.positives.value should be 2")
	}
	if counter1.Value() != -4 {
		t.Fatalf("counter value should be -0")
	}
}

func TestPNCounterMergeIdempotent(t *testing.T) {
	counter1 := NewPNCounter("counter1", "srv1", initPNNodes())
	counter1.Increment()
	counter1.Increment()
	counter1.Decrement()
	counter1.Decrement()
	counter2 := NewPNCounter("counter1", "srv2", initPNNodes())
	counter2.Increment()
	counter2.Decrement()

	counter1.Merge(counter2)

	if counter1.positives.Value() != 3 {
		t.Fatalf("counter1.positives.value should be 2")
	}
	if counter1.negatives.Value() != 3 {
		t.Fatalf("counter1.positives.value should be 2")
	}
	if counter1.Value() != 0 {
		t.Fatalf("counter value should be 0")
	}

	counter1.Merge(counter2)
	if counter1.positives.Value() != 3 {
		t.Fatalf("counter1.positives.value should be 2")
	}
	if counter1.negatives.Value() != 3 {
		t.Fatalf("counter1.positives.value should be 2")
	}
	if counter1.Value() != 0 {
		t.Fatalf("counter value should be 0")
	}
}

func TestPNCounterMergeUnmatchedName(t *testing.T) {
	counter1 := NewPNCounter("counter1", "srv1", initPNNodes())
	counter1.Increment()
	counter1.Increment()
	counter1.Decrement()
	counter1.Decrement()
	counter2 := NewPNCounter("counter2", "srv2", initPNNodes())
	counter2.Increment()
	counter2.Increment()

	counter1.Merge(counter2) // should be no-op due to mismatched name

	if counter1.positives.Value() != 2 {
		t.Fatalf("counter1.positives.value should be 2")
	}
	if counter1.negatives.Value() != 2 {
		t.Fatalf("counter1.positives.value should be 2")
	}
	if counter1.Value() != 0 {
		t.Fatalf("counter value should be 0")
	}
}
