package crdt

import (
	"testing"
	"time"
)

func TestLWWRegisterInitialization(t *testing.T) {
	reg := NewLWWRegister("reg1", 0)
	if reg.Value() != 0 {
		t.Fatalf("reg.value should be 0")
	}
}

func TestLWWRegisterSet(t *testing.T) {
	reg := NewLWWRegister("reg1", 0)
	reg.Assign(5)
	if reg.Value() != 5 {
		t.Fatalf("reg.value should be 5")
	}
}

func TestLWWRegisterMergeNoChange(t *testing.T) {
	reg := NewLWWRegister("reg1", 0)
	reg.Assign(5)
	if reg.Value() != 5 {
		t.Fatalf("reg.value should be 5, was %d", reg.Value())
	}

	reg2 := NewLWWRegister("reg2", 0)
	reg2.Assign(3)
	reg.Assign(10)

	reg.Merge(reg2)
	if reg.Value() != 10 {
		t.Fatalf("reg.value should be 10, was %d", reg.Value())
	}
}

func TestLWWRegisterMergeChange(t *testing.T) {
	reg := NewLWWRegister("reg1", 0)
	reg.Assign(5)
	reg2 := NewLWWRegister("reg2", 0)
	time.Sleep(1 * time.Millisecond) // Sleeping to force time increment
	reg2.Assign(7)

	reg.Merge(reg2)

	if reg.Value() != 7 {
		t.Fatalf("reg.value should be 7, was %d", reg.Value())
	}
}

func TestLWWRegisterMergeIdempotence(t *testing.T) {
	reg := NewLWWRegister("reg1", 0)
	reg.Assign(5)

	reg2 := NewLWWRegister("reg2", 0)
	time.Sleep(1 * time.Millisecond) // Sleeping to force time increment
	reg2.Assign(7)

	reg.Merge(reg2)

	if reg.Value() != 7 {
		t.Fatalf("reg.value should be 7, was %d", reg.Value())
	}

	reg.Merge(reg2)

	if reg.Value() != 7 {
		t.Fatalf("reg.value should be 7, was %d", reg.Value())
	}

	reg.Merge(reg2)
	reg.Merge(reg2)

	if reg.Value() != 7 {
		t.Fatalf("reg.value should be 7, was %d", reg.Value())
	}
}
