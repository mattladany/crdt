package crdt

import "testing"

func testLWWRegisterInitialization(t *testing.T) {
	reg := NewLWWRegister(0)
	if reg.Get() != 0 {
		t.Fatalf("reg.value should be 0")
	}
}

func testLWWRegisterSet(t *testing.T) {
	reg := NewLWWRegister(0)
	reg.Set(5)
	if reg.Get() != 5 {
		t.Fatalf("reg.value should be 5")
	}
}

func testLWWRegisterSetClockInc(t *testing.T) {
	reg := NewLWWRegister(0)
	reg.Set(5)
	if reg.Get() != 5 {
		t.Fatalf("reg.value should be 5")
	}
	if reg.clock != 1 {
		t.Fatalf("reg.clock should be 1")
	}
}

func testLWWRegisterMergeNoChange(t *testing.T) {
	reg := NewLWWRegister(0)
	reg.Set(5)
	reg.Set(10)
	if reg.Get() != 10 {
		t.Fatalf("reg.value should be 10")
	}
	if reg.clock != 2 {
		t.Fatalf("reg.clock should be 2")
	}

	reg2 := NewLWWRegister(0)
	reg2.Set(3)

	reg.Merge(reg2)
	if reg.Get() != 10 {
		t.Fatalf("reg.value should be 10")
	}
	if reg.clock != 2 {
		t.Fatalf("reg.clock should be 2")
	}
}

func testLWWRegisterMergeChange(t *testing.T) {
	reg := NewLWWRegister(0)
	reg.Set(5)
	if reg.Get() != 5 {
		t.Fatalf("reg.value should be 5")
	}
	if reg.clock != 2 {
		t.Fatalf("reg.clock should be 1")
	}

	reg2 := NewLWWRegister(0)
	reg2.Set(3)
	reg2.Set(7)

	reg.Merge(reg2)

	if reg.Get() != 7 {
		t.Fatalf("reg.value should be 7")
	}
	if reg.clock != 2 {
		t.Fatalf("reg.clock should be 2")
	}
}

func testLWWRegisterMergeIdempotence(t *testing.T) {
	reg := NewLWWRegister(0)
	reg.Set(5)
	if reg.Get() != 5 {
		t.Fatalf("reg.value should be 5")
	}
	if reg.clock != 2 {
		t.Fatalf("reg.clock should be 1")
	}

	reg2 := NewLWWRegister(0)
	reg2.Set(3)
	reg2.Set(7)

	reg.Merge(reg2)

	if reg.Get() != 7 {
		t.Fatalf("reg.value should be 7")
	}
	if reg.clock != 2 {
		t.Fatalf("reg.clock should be 2")
	}

	reg.Merge(reg2)
	reg.Merge(reg2)
	reg.Merge(reg2)

	if reg.Get() != 7 {
		t.Fatalf("reg.value should be 7")
	}
	if reg.clock != 2 {
		t.Fatalf("reg.clock should be 2")
	}
}
