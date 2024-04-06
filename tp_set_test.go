package crdt

import "testing"

func TestTPSetInitialization(t *testing.T) {
	tpset := NewTwoPhaseSet[int]("tpset")
	if tpset.Size() != 0 {
		t.Fatalf("tpset should initialize to a size of 0")
	}
}

func TestTPSetAdd(t *testing.T) {
	tpset := NewTwoPhaseSet[int]("tpset")
	tpset.Add(5)
	if tpset.Size() != 1 {
		t.Fatalf("tpset size should be 1")
	}
}

func TestTPSetLookupTrue(t *testing.T) {
	tpset := NewTwoPhaseSet[int]("tpset")
	tpset.Add(5)
	if !tpset.Lookup(5) {
		t.Fatalf("tpset lookup should have been true")
	}
}

func TestTPSetLookupFalse(t *testing.T) {
	tpset := NewTwoPhaseSet[int]("tpset")
	tpset.Add(5)
	if tpset.Lookup(10) {
		t.Fatalf("tpset lookup should have been false")
	}
}

func TestTPSetValidRemove(t *testing.T) {
	tpset := NewTwoPhaseSet[int]("tpset")
	tpset.Add(5)
	tpset.Remove(5)
	if tpset.Size() != 0 {
		t.Fatalf("tpset size should be 0")
	}
	if tpset.Lookup(5) {
		t.Fatalf("tpset lookup should have been false")
	}
}

func TestTPSetInvalidRemove(t *testing.T) {
	tpset := NewTwoPhaseSet[int]("tpset")
	tpset.Remove(5)
	tpset.Add(5)
	if tpset.Size() != 1 {
		t.Fatalf("tpset size should be 1")
	}
	if !tpset.Lookup(5) {
		t.Fatalf("tpset lookup should have been true")
	}
}

func TestTPSetMerge(t *testing.T) {
	tpset := NewTwoPhaseSet[int]("tpset")
	tpset.Add(5)
	tpset2 := NewTwoPhaseSet[int]("tpset")
	tpset2.Add(10)
	tpset2.Add(15)
	tpset2.Remove(10)

	if tpset.Lookup(10) {
		t.Fatalf("tpset lookup for 10 should have been false before merge")
	}
	if tpset.Lookup(15) {
		t.Fatalf("tpset lookup for 15 should have been false before merge")
	}

	tpset.Merge(tpset2)

	if tpset.Size() != 2 {
		t.Fatalf("tpset size should have been 2 after merge")
	}
	if tpset.Lookup(10) {
		t.Fatalf("tpset lookup for 10 should have been false after merge")
	}
	if !tpset.Lookup(15) {
		t.Fatalf("tpset lookup for 10 should have been true after merge")
	}

}

func TestTPSetMergeIdempotent(t *testing.T) {
	tpset := NewTwoPhaseSet[int]("tpset")
	tpset.Add(5)
	tpset2 := NewTwoPhaseSet[int]("tpset")
	tpset2.Add(10)
	tpset2.Add(15)
	tpset2.Remove(10)

	if tpset.Lookup(10) {
		t.Fatalf("tpset lookup for 10 should have been false before merge")
	}
	if tpset.Lookup(15) {
		t.Fatalf("tpset lookup for 15 should have been false before merge")
	}

	tpset.Merge(tpset2)

	if tpset.Size() != 2 {
		t.Fatalf("tpset size should have been 2 after merge")
	}
	if tpset.Lookup(10) {
		t.Fatalf("tpset lookup for 10 should have been false after merge")
	}
	if !tpset.Lookup(15) {
		t.Fatalf("tpset lookup for 10 should have been true after merge")
	}

	tpset.Merge(tpset2)
	tpset.Merge(tpset2)
	tpset.Merge(tpset2)

	if tpset.Size() != 2 {
		t.Fatalf("tpset size should have been 2 after merge")
	}
	if tpset.Lookup(10) {
		t.Fatalf("tpset lookup for 10 should have been false after merge")
	}
	if !tpset.Lookup(15) {
		t.Fatalf("tpset lookup for 10 should have been true after merge")
	}
}
