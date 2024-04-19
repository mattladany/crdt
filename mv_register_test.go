package crdt

import (
	"testing"
	"time"
)

func initMVRegisterNodes() map[string]registerValue[int] {
	nodes := make(map[string]registerValue[int])
	node1 := new(registerValue[int])
	node1.value = 0
	node1.timestamp = time.Now().UnixNano()
	node2 := new(registerValue[int])
	node2.value = 0
	node2.timestamp = time.Now().UnixNano()
	node3 := new(registerValue[int])
	node3.value = 0
	node3.timestamp = time.Now().UnixNano()

	nodes["node1"] = *node1
	nodes["node2"] = *node2
	nodes["node3"] = *node3
	return nodes
}

func TestMVRegisterInitialization(t *testing.T) {
	reg := NewMVRegister("node1", initMVRegisterNodes())
	for _, subValue := range reg.Value() {
		if subValue != 0 {
			t.Fatalf("subValue should be 0")
		}
	}

}

func TestMVRegisterAssign(t *testing.T) {
	reg := NewMVRegister("node1", initMVRegisterNodes())
	reg.Assign(5)
	if reg.Value()["node1"] != 5 {
		t.Fatalf("value should be 5")
	}
}

func TestMVRegisterMergeNoChange(t *testing.T) {
	reg := NewMVRegister("node1", initMVRegisterNodes())
	reg.Assign(5)
	if reg.Value()["node1"] != 5 {
		t.Fatalf("value should be 5")
	}

	reg2 := NewMVRegister("node2", initMVRegisterNodes())
	reg.Merge(reg2)
	if reg.Value()["node1"] != 5 {
		t.Fatalf("reg[node1] should be 5, was %d", reg.Value()["node1"])
	}
	if reg.Value()["node2"] != 0 {
		t.Fatalf("reg[node2] should be 0, was %d", reg.Value()["node2"])
	}
}

func TestMVRegisterMergeChange(t *testing.T) {
	reg := NewMVRegister("node1", initMVRegisterNodes())
	reg.Assign(5)
	if reg.Value()["node1"] != 5 {
		t.Fatalf("value should be 5")
	}

	reg2 := NewMVRegister("node2", initMVRegisterNodes())
	time.Sleep(1 * time.Millisecond)
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
	reg := NewMVRegister("node1", initMVRegisterNodes())
	reg.Assign(5)
	if reg.Value()["node1"] != 5 {
		t.Fatalf("value should be 5")
	}

	reg2 := NewMVRegister("node2", initMVRegisterNodes())
	time.Sleep(1 * time.Millisecond)
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
