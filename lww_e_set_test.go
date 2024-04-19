package crdt

import (
	"testing"
	"time"
)

func TestLWWESetInitialization(t *testing.T) {
	set := NewLWWESet[int]("lwweset")
	if set.Size() != 0 {
		t.Fatalf("lwweset should initialize to a size of 0")
	}
}

func TestLSSESetAdd(t *testing.T) {
	set := NewLWWESet[int]("lwweset")
	set.Add(5)
	if set.Size() != 1 {
		t.Fatalf("lwweset size should be 1")
	}
}

func TestLWWESetLookupTrue(t *testing.T) {
	set := NewLWWESet[int]("lwweset")
	set.Add(5)
	if !set.Lookup(5) {
		t.Fatalf("lwweset lookup should have been true")
	}
}

func TestLWWESetLookupFalse(t *testing.T) {
	set := NewLWWESet[int]("lwweset")
	set.Add(5)
	if set.Lookup(10) {
		t.Fatalf("lwweset lookup should have been false")
	}
}

func TestLWWESetValidRemove(t *testing.T) {
	set := NewLWWESet[int]("lwweset")
	set.Add(5)
	set.Remove(5)
	if set.Size() != 0 {
		t.Fatalf("lwweset size should be 0")
	}
	if set.Lookup(5) {
		t.Fatalf("lwweset lookup should have been false")
	}
}

func TestLWWESetInvalidRemove(t *testing.T) {
	set := NewLWWESet[int]("lwweset")
	set.Remove(5)
	set.Add(5)
	if set.Size() != 1 {
		t.Fatalf("lwweset size should be 1")
	}
	if !set.Lookup(5) {
		t.Fatalf("lwweset lookup should have been true")
	}
}

func TestLWWESetValidRemoveWithReAdd(t *testing.T) {
	set := NewLWWESet[int]("lwweset")
	set.Add(5)
	set.Remove(5)
	time.Sleep(1 * time.Millisecond)
	set.Add(5)
	if set.Size() != 1 {
		t.Fatalf("lwweset size should be 1")
	}
	if !set.Lookup(5) {
		t.Fatalf("lwweset lookup should have been true")
	}
}

func TestLWWESetMerge(t *testing.T) {
	set1 := NewLWWESet[int]("lwweset")
	set1.Add(5)
	set2 := NewLWWESet[int]("lwweset")
	set2.Add(10)
	set2.Add(15)
	set2.Remove(10)

	if set1.Lookup(10) {
		t.Fatalf("lwweset lookup for 10 should have been false before merge")
	}
	if set1.Lookup(15) {
		t.Fatalf("lwweset lookup for 15 should have been false before merge")
	}

	set1.Merge(set2)

	if set1.Size() != 2 {
		t.Fatalf("lwweset size should have been 2 after merge")
	}
	if set1.Lookup(10) {
		t.Fatalf("lwweset lookup for 10 should have been false after merge")
	}
	if !set1.Lookup(15) {
		t.Fatalf("lwweset lookup for 10 should have been true after merge")
	}
}

func TestLWWESetMergeWithDelayedAdd(t *testing.T) {
	set1 := NewLWWESet[int]("lwweset")
	set1.Add(5)
	time.Sleep(1 * time.Millisecond)
	set1.Remove(5)

	set2 := NewLWWESet[int]("lwweset")
	time.Sleep(1 * time.Millisecond)
	set2.Add(5)

	if set1.Size() != 0 {
		t.Fatalf("lwweset size should have been 0 before merge")
	}
	if set1.Lookup(5) {
		t.Fatalf("lwweset lookup for 5 should have been false before merge")
	}

	set1.Merge(set2)

	if set1.Size() != 1 {
		t.Fatalf("lwweset size should have been 1 after merge")
	}
	if !set1.Lookup(5) {
		t.Fatalf("lwweset lookup for 5 should have been true after merge")
	}
}

func TestLWWESetMergeIdempotent(t *testing.T) {
	set1 := NewLWWESet[int]("lwweset")
	set1.Add(5)
	time.Sleep(1 * time.Millisecond)
	set1.Remove(5)

	set2 := NewLWWESet[int]("lwweset")
	time.Sleep(1 * time.Millisecond)
	set2.Add(5)

	if set1.Size() != 0 {
		t.Fatalf("lwweset size should have been 0 before merge")
	}
	if set1.Lookup(5) {
		t.Fatalf("lwweset lookup for 5 should have been false before merge")
	}

	set1.Merge(set2)

	if set1.Size() != 1 {
		t.Fatalf("lwweset size should have been 1 after merge")
	}
	if !set1.Lookup(5) {
		t.Fatalf("lwweset lookup for 5 should have been true after merge")
	}

	set1.Merge(set2)
	set1.Merge(set2)
	set1.Merge(set2)

	if set1.Size() != 1 {
		t.Fatalf("lwweset size should have been 2 after merge")
	}
	if !set1.Lookup(5) {
		t.Fatalf("lwweset lookup for 5 should have been true after merge")
	}
}

func TestLWWESetMergeUnmatchedName(t *testing.T) {
	set1 := NewLWWESet[int]("lwweset")
	set1.Add(5)
	set2 := NewLWWESet[int]("lwweset2")
	set2.Add(10)
	set2.Add(15)
	set2.Remove(10)

	if set1.Lookup(10) {
		t.Fatalf("lwweset lookup for 10 should have been false before merge")
	}
	if set1.Lookup(15) {
		t.Fatalf("lwweset lookup for 15 should have been false before merge")
	}

	set1.Merge(set2)

	if set1.Size() != 1 {
		t.Fatalf("lwweset size should have been 2 after merge")
	}
	if set1.Lookup(10) {
		t.Fatalf("lwweset lookup for 10 should have been false after merge")
	}
	if set1.Lookup(15) {
		t.Fatalf("lwweset lookup for 10 should have been false after merge")
	}
}
