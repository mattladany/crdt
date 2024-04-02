package crdt

type VectorClock struct {
	name   string
	clocks map[string]int
}

func NewVectorClock(name string, clocks map[string]int) *VectorClock {
	lc := new(VectorClock)
	lc.name = name
	lc.clocks = clocks
	if _, exists := clocks[name]; !exists {
		lc.clocks[name] = 0
	}
	return lc
}

func (vc *VectorClock) inc() {
	vc.clocks[vc.name]++
}

func (clock *VectorClock) merge(that *VectorClock) bool {
	updated := false
	for name, value := range clock.clocks {
		if that.clocks[name] > value {
			clock.clocks[name] = that.clocks[name]
			updated = true
		}
	}

	if updated {
		clock.clocks[clock.name]++
	}

	return updated
}
