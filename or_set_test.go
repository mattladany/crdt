package crdt

import "testing"

func TestORSetInitialization(t *testing.T) {
	orset := NewORSet[int]("orset")
	if orset.Size() != 0 {
		t.Fatalf("orset should initialize to a size of 0")
	}
}

func TestORSetAdd(t *testing.T) {
	orset := NewORSet[int]("orset")
	orset.Add(5)
	if orset.Size() != 1 {
		t.Fatalf("orset size should be 1")
	}
}

func TestORSetLookupTrue(t *testing.T) {
	orset := NewORSet[int]("orset")
	orset.Add(5)
	if !orset.Lookup(5) {
		t.Fatalf("orset lookup should have been true")
	}
}

func TestORSetLookupFalse(t *testing.T) {
	orset := NewORSet[int]("orset")
	orset.Add(5)
	if orset.Lookup(10) {
		t.Fatalf("orset lookup should have been false")
	}
}

func TestORSetValidRemove(t *testing.T) {
	orset := NewORSet[int]("orset")
	orset.Add(5)
	orset.Remove(5)
	if orset.Size() != 0 {
		t.Fatalf("orset size should be 0")
	}
	if orset.Lookup(5) {
		t.Fatalf("orset lookup should have been false")
	}
}

func TestORSetInvalidRemove(t *testing.T) {
	orset := NewORSet[int]("orset")
	orset.Remove(5)
	orset.Add(5)
	if orset.Size() != 1 {
		t.Fatalf("orset size should be 1")
	}
	if !orset.Lookup(5) {
		t.Fatalf("orset lookup should have been true")
	}
}

func TestORSetMerge(t *testing.T) {
	orset := NewORSet[int]("orset")
	orset.Add(5)
	orset2 := NewORSet[int]("orset")
	orset2.Add(10)
	orset2.Add(15)
	orset2.Remove(10)

	if orset.Lookup(10) {
		t.Fatalf("orset lookup for 10 should have been false before merge")
	}
	if orset.Lookup(15) {
		t.Fatalf("orset lookup for 15 should have been false before merge")
	}

	if orset2.Lookup(10) {
		t.Fatalf("orset2 lookup for 10 should have been false")
	}
	if !orset2.Lookup(15) {
		t.Fatalf("orset2 lookup for 15 should have been true")
	}

	orset.Merge(orset2)

	if orset.Size() != 2 {
		t.Fatalf("orset size should have been 2 after merge, was %d", orset.Size())
	}
	if orset.Lookup(10) {
		t.Fatalf("orset lookup for 10 should have been false after merge")
	}
	if !orset.Lookup(15) {
		t.Fatalf("orset lookup for 15 should have been true after merge")
	}

}

func TestORSetMergeIdempotent(t *testing.T) {
	orset := NewORSet[int]("orset")
	orset.Add(5)
	orset2 := NewORSet[int]("orset")
	orset2.Add(10)
	orset2.Add(15)
	orset2.Remove(10)

	if orset.Lookup(10) {
		t.Fatalf("orset lookup for 10 should have been false before merge")
	}
	if orset.Lookup(15) {
		t.Fatalf("orset lookup for 15 should have been false before merge")
	}

	orset.Merge(orset2)

	if orset.Size() != 2 {
		t.Fatalf("orset size should have been 2 after merge, was %d", orset.Size())
	}
	if orset.Lookup(10) {
		t.Fatalf("orset lookup for 10 should have been false after merge")
	}
	if !orset.Lookup(15) {
		t.Fatalf("orset lookup for 15 should have been true after merge")
	}

	orset.Merge(orset2)
	orset.Merge(orset2)
	orset.Merge(orset2)

	if orset.Size() != 2 {
		t.Fatalf("orset size should have been 2 after merge, was %d", orset.Size())
	}
	if orset.Lookup(10) {
		t.Fatalf("orset lookup for 10 should have been false after merge")
	}
	if !orset.Lookup(15) {
		t.Fatalf("orset lookup for 10 should have been true after merge")
	}
}
