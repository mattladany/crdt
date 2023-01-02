package crdt

import "testing"

func TestGCounterInitialization(t *testing.T) {
	counter := NewGCounter()
	if counter.Value() != 0 {
		t.Fatalf("counter value should initialize to 0")
	}
}

func TestGCounterSingleIncrement(t *testing.T) {
	counter := NewGCounter()
	counter.Increment()
	if counter.Value() != 1 {
		t.Fatalf("counter value should be 1")
	}
}

func TestGCounterMultiIncrement(t *testing.T) {
	counter := NewGCounter()
	counter.Increment()
	counter.Increment()
	counter.Increment()
	counter.Increment()
	if counter.Value() != 4 {
		t.Fatalf("counter value should be 4")
	}
}

func TestGCounterMergeStaySame(t *testing.T) {
	counter1 := NewGCounter()
	counter1.Increment()
	counter1.Increment()
	counter2 := NewGCounter()
	counter2.Increment()
	counter1.Merge(counter2)
	if counter1.Value() != 2 {
		t.Fatalf("merged counter value should be 2")
	}
}

func TestGCounterMergeChange(t *testing.T) {
	counter1 := NewGCounter()
	counter1.Increment()
	counter2 := NewGCounter()
	counter2.Increment()
	counter2.Increment()
	counter1.Merge(counter2)
	if counter1.Value() != 2 {
		t.Fatalf("merged counter value should be 2")
	}
}
