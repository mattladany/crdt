package crdt

import "testing"

func initClock(name string) *VectorClock {
	clocks := make(map[string]int)
	clocks["srv1"] = 0
	clocks["srv2"] = 0
	clocks["srv3"] = 0
	clocks["srv4"] = 0
	return NewVectorClock(name, clocks)
}

func TestGCounterInitialization(t *testing.T) {
	counter := NewGCounter(initClock("srv1"))
	if counter.Value() != 0 {
		t.Fatalf("counter value should initialize to 0")
	}
}

func TestGCounterSingleIncrement(t *testing.T) {
	counter := NewGCounter(initClock("srv1"))
	counter.Increment()
	if counter.Value() != 1 {
		t.Fatalf("counter value should be 1")
	}
}

func TestGCounterMultiIncrement(t *testing.T) {
	counter := NewGCounter(initClock("srv1"))
	counter.Increment()
	counter.Increment()
	counter.Increment()
	counter.Increment()
	if counter.Value() != 4 {
		t.Fatalf("counter value should be 4")
	}
}

func TestGCounterMergeStaySame(t *testing.T) {
	counter1 := NewGCounter(initClock("srv1"))
	counter1.Increment()
	counter1.Increment()
	counter2 := NewGCounter(initClock("srv2"))
	counter2.Increment()
	counter1.Merge(counter2)
	if counter1.Value() != 2 {
		t.Fatalf("merged counter value should be 2")
	}
}

func TestGCounterMergeChange(t *testing.T) {
	counter1 := NewGCounter(initClock("srv1"))
	counter1.Increment()
	counter2 := NewGCounter(initClock("srv2"))
	counter2.Increment()
	counter2.Increment()
	counter1.Merge(counter2)
	if counter1.Value() != 2 {
		t.Fatalf("merged counter value should be 2")
	}
}
