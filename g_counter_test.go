package crdt

import "testing"

func TestGCounterInitialization(t *testing.T) {
	counter := NewGCounter("counter1", "srv1")
	if counter.Value() != 0 {
		t.Fatalf("counter value should initialize to 0")
	}
}

func TestGCounterSingleIncrement(t *testing.T) {
	counter := NewGCounter("counter1", "srv1")
	counter.Increment()
	if counter.Value() != 1 {
		t.Fatalf("counter value should be 1")
	}
}

func TestGCounterMultiIncrement(t *testing.T) {
	counter := NewGCounter("counter1", "srv1")
	counter.Increment()
	counter.Increment()
	counter.Increment()
	counter.Increment()
	if counter.Value() != 4 {
		t.Fatalf("counter value should be 4")
	}
}

func TestGCounterMerge(t *testing.T) {
	counter1 := NewGCounter("counter1", "srv1")
	counter1.Increment()
	counter1.Increment()
	counter2 := NewGCounter("counter1", "srv2")
	counter2.Increment()
	counter1.Merge(counter2)
	if counter1.Value() != 3 {
		t.Fatalf("merged counter value should be 3")
	}
}

func TestGCounterMergeMismatchedId(t *testing.T) {
	counter1 := NewGCounter("counter1", "srv1")
	counter1.Increment()
	counter1.Increment()
	counter2 := NewGCounter("counter2", "srv2")
	counter2.Increment()
	counter1.Merge(counter2)
	if counter1.Value() != 2 {
		t.Fatalf("merged counter value should be 2 since names did not match")
	}
}

func TestGCounterMergeIdempotent(t *testing.T) {
	counter1 := NewGCounter("counter1", "srv1")
	counter1.Increment()
	counter1.Increment()
	counter2 := NewGCounter("counter1", "srv2")
	counter2.Increment()
	counter1.Merge(counter2)
	if counter1.Value() != 3 {
		t.Fatalf("merged counter value should be 3")
	}
	counter1.Merge(counter2)
	if counter1.Value() != 3 {
		t.Fatalf("merged counter value should be 3")
	}
}

func TestGCounterMergeSameNode(t *testing.T) {
	counter1 := NewGCounter("counter1", "srv1")
	counter1.Increment()
	counter1.Increment()
	counter2 := NewGCounter("counter1", "srv2")
	counter2.Increment()
	counter1.Merge(counter2)
	if counter1.Value() != 3 {
		t.Fatalf("merged counter value should be 3")
	}
	counter2.Increment()
	counter1.Merge(counter2)
	if counter1.Value() != 4 {
		t.Fatalf("merged counter value should be 4")
	}
}
