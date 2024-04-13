package crdt

import "testing"

func initNodes() []string {
	return []string{"srv1", "srv2", "srv3", "srv4"}
}

func TestGCounterInitialization(t *testing.T) {
	counter := NewGCounter("counter1", "srv1", initNodes())
	if counter.Value() != 0 {
		t.Fatalf("counter value should initialize to 0")
	}
}

func TestGCounterSingleIncrement(t *testing.T) {
	counter := NewGCounter("counter1", "srv1", initNodes())
	counter.Increment()
	if counter.Value() != 1 {
		t.Fatalf("counter value should be 1")
	}
}

func TestGCounterMultiIncrement(t *testing.T) {
	counter := NewGCounter("counter1", "srv1", initNodes())
	counter.Increment()
	counter.Increment()
	counter.Increment()
	counter.Increment()
	if counter.Value() != 4 {
		t.Fatalf("counter value should be 4")
	}
}

func TestGCounterMerge(t *testing.T) {
	counter1 := NewGCounter("counter1", "srv1", initNodes())
	counter1.Increment()
	counter1.Increment()
	counter2 := NewGCounter("counter1", "srv2", initNodes())
	counter2.Increment()
	counter1.Merge(counter2)
	if counter1.Value() != 3 {
		t.Fatalf("merged counter value should be 3")
	}
}

func TestGCounterMergeMismatchedId(t *testing.T) {
	counter1 := NewGCounter("counter1", "srv1", initNodes())
	counter1.Increment()
	counter1.Increment()
	counter2 := NewGCounter("counter2", "srv2", initNodes())
	counter2.Increment()
	counter1.Merge(counter2)
	if counter1.Value() != 2 {
		t.Fatalf("merged counter value should be 2 since names did not match")
	}
}

func TestGCounterMergeIdempotent(t *testing.T) {
	counter1 := NewGCounter("counter1", "srv1", initNodes())
	counter1.Increment()
	counter1.Increment()
	counter2 := NewGCounter("counter1", "srv2", initNodes())
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
