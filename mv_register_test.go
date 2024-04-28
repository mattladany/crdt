package crdt

import (
	"testing"
	"time"
)

func TestMVRegisterInitialization(t *testing.T) {
	reg := NewMVRegister("mvreg", "node1", 5)
	for _, subValue := range reg.Value() {
		if subValue != 5 {
			t.Fatalf("subValue should be 5")
		}
	}

}

func TestMVRegisterAssign(t *testing.T) {
	reg := NewMVRegister("mvreg", "node1", 0)
	reg.Assign(5)
	if reg.Value()["node1"] != 5 {
		t.Fatalf("value should be 5")
	}
}

func TestMVRegisterMergeNoChange(t *testing.T) {
	reg := NewMVRegister("mvreg", "node1", 0)
	reg.Assign(5)
	if reg.Value()["node1"] != 5 {
		t.Fatalf("value should be 5")
	}

	reg2 := NewMVRegister("mvreg", "node2", 0)
	reg.Merge(reg2)
	if reg.Value()["node1"] != 5 {
		t.Fatalf("reg[node1] should be 5, was %d", reg.Value()["node1"])
	}
	if reg.Value()["node2"] != 0 {
		t.Fatalf("reg[node2] should be 0, was %d", reg.Value()["node2"])
	}
}

func TestMVRegisterMergeMismatchedName(t *testing.T) {
	reg := NewMVRegister("mvreg", "node1", 0)
	reg.Assign(5)
	if reg.Value()["node1"] != 5 {
		t.Fatalf("value should be 5")
	}

	reg2 := NewMVRegister("mvreg2", "node2", 0)
	reg.Merge(reg2)
	if reg.Value()["node1"] != 5 {
		t.Fatalf("reg[node1] should be 5, was %d", reg.Value()["node1"])
	}
	if _, exists := reg.Value()["node2"]; exists {
		t.Fatalf("reg[node2] should not exist since the name was mismatched")
	}
}

func TestMVRegisterMergeChange(t *testing.T) {
	reg := NewMVRegister("mvreg", "node1", 0)
	reg.Assign(5)
	if reg.Value()["node1"] != 5 {
		t.Fatalf("value should be 5")
	}

	reg2 := NewMVRegister("mvreg", "node2", 0)
	time.Sleep(50 * time.Millisecond)
	reg2.Assign(10)
	reg.Merge(reg2)
	if reg.Value()["node1"] != 5 {
		t.Fatalf("reg[node1] should be 5, was %d", reg.Value()["node1"])
	}
	if reg.Value()["node2"] != 10 {
		t.Fatalf("reg[node2] should be 10, was %d", reg.Value()["node2"])
	}
}

func TestMVRegisterMergeIdempotence(t *testing.T) {
	reg := NewMVRegister("mvreg", "node1", 0)
	reg.Assign(5)
	if reg.Value()["node1"] != 5 {
		t.Fatalf("value should be 5")
	}

	reg2 := NewMVRegister("mvreg", "node2", 0)
	time.Sleep(50 * time.Millisecond)
	reg2.Assign(10)

	reg.Merge(reg2)
	reg.Merge(reg2)
	reg.Merge(reg2)

	if reg.Value()["node1"] != 5 {
		t.Fatalf("reg[node1] should be 5, was %d", reg.Value()["node1"])
	}
	if reg.Value()["node2"] != 10 {
		t.Fatalf("reg[node2] should be 10, was %d", reg.Value()["node2"])
	}
}
