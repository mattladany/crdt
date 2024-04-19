package crdt

import "testing"

func TestPNSetInitialization(t *testing.T) {
	pnset := NewPNSet[int]("pnset")
	if pnset.Size() != 0 {
		t.Fatalf("pnset should initialize to a size of 0")
	}
}

func TestPNSetAdd(t *testing.T) {
	pnset := NewPNSet[int]("pnset")
	pnset.Add(5)
	if pnset.Size() != 1 {
		t.Fatalf("pnset size should be 1")
	}
}

func TestPNSetLookupTrue(t *testing.T) {
	pnset := NewPNSet[int]("pnset")
	pnset.Add(5)
	if !pnset.Lookup(5) {
		t.Fatalf("pnset lookup should have been true")
	}
}

func TestPNSetLookupFalse(t *testing.T) {
	pnset := NewPNSet[int]("pnset")
	pnset.Add(5)
	if pnset.Lookup(10) {
		t.Fatalf("pnset lookup should have been false")
	}
}

func TestPNSetValidRemove(t *testing.T) {
	pnset := NewPNSet[int]("pnset")
	pnset.Add(5)
	pnset.Remove(5)
	if pnset.Size() != 0 {
		t.Fatalf("pnset size should be 0")
	}
	if pnset.Lookup(5) {
		t.Fatalf("pnset lookup should have been false")
	}
}

func TestPNSetInvalidRemove(t *testing.T) {
	pnset := NewPNSet[int]("pnset")
	pnset.Remove(5)
	pnset.Add(5)
	if pnset.Size() != 1 {
		t.Fatalf("pnset size should be 1")
	}
	if !pnset.Lookup(5) {
		t.Fatalf("pnset lookup should have been true")
	}
}

func TestPNSetMerge(t *testing.T) {
	pnset := NewPNSet[int]("pnset")
	pnset.Add(5)
	pnset2 := NewPNSet[int]("pnset")
	pnset2.Add(10)
	pnset2.Add(15)
	pnset2.Remove(10)

	if pnset.Lookup(10) {
		t.Fatalf("pnset lookup for 10 should have been false before merge")
	}
	if pnset.Lookup(15) {
		t.Fatalf("pnset lookup for 15 should have been false before merge")
	}

	pnset.Merge(pnset2)

	if pnset.Size() != 2 {
		t.Fatalf("pnset size should have been 2 after merge")
	}
	if pnset.Lookup(10) {
		t.Fatalf("pnset lookup for 10 should have been false after merge")
	}
	if !pnset.Lookup(15) {
		t.Fatalf("pnset lookup for 15 should have been true after merge")
	}
}

func TestPNSetMergeRemoveRaceConditionCausingEmptyAdd(t *testing.T) {
	pnset1 := NewPNSet[int]("pnset")
	pnset1.Add(5)
	pnset2 := NewPNSet[int]("pnset")
	pnset2.Merge(pnset1)

	pnset1.Remove(5)
	pnset2.Remove(5)

	pnset2.Merge(pnset1)
	pnset2.Add(5)

	if pnset2.Size() != 0 {
		t.Fatalf("pnset size should have remained 0 due to the remove race condition")
	}
}

func TestPNSetMergeIdempotent(t *testing.T) {
	pnset := NewPNSet[int]("pnset")
	pnset.Add(5)
	pnset2 := NewPNSet[int]("pnset")
	pnset2.Add(10)
	pnset2.Add(15)
	pnset2.Remove(10)

	if pnset.Lookup(10) {
		t.Fatalf("pnset lookup for 10 should have been false before merge")
	}
	if pnset.Lookup(15) {
		t.Fatalf("pnset lookup for 15 should have been false before merge")
	}

	pnset.Merge(pnset2)

	if pnset.Size() != 2 {
		t.Fatalf("pnset size should have been 2 after merge")
	}
	if pnset.Lookup(10) {
		t.Fatalf("pnset lookup for 10 should have been false after merge")
	}
	if !pnset.Lookup(15) {
		t.Fatalf("pnset lookup for 15 should have been true after merge")
	}

	pnset.Merge(pnset2)
	pnset.Merge(pnset2)
	pnset.Merge(pnset2)

	if pnset.Size() != 2 {
		t.Fatalf("pnset size should have been 2 after merge")
	}
	if pnset.Lookup(10) {
		t.Fatalf("pnset lookup for 10 should have been false after merge")
	}
	if !pnset.Lookup(15) {
		t.Fatalf("pnset lookup for 15 should have been true after merge")
	}
}

func TestPNSetMergeUnmatchedName(t *testing.T) {
	pnset := NewPNSet[int]("pnset")
	pnset.Add(5)
	pnset2 := NewPNSet[int]("pnset2")
	pnset2.Add(10)
	pnset2.Add(15)
	pnset2.Remove(10)

	if pnset.Lookup(10) {
		t.Fatalf("pnset lookup for 10 should have been false before merge")
	}
	if pnset.Lookup(15) {
		t.Fatalf("pnset lookup for 15 should have been false before merge")
	}

	pnset.Merge(pnset2)

	if pnset.Size() != 1 {
		t.Fatalf("pnset size should have been 1 after merge")
	}
	if pnset.Lookup(10) {
		t.Fatalf("pnset lookup for 10 should have been false after merge")
	}
	if pnset.Lookup(15) {
		t.Fatalf("pnset lookup for 15 should have been false after merge")
	}
}
