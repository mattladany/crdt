package crdt

// Uses a PNCounter that is assumed to have only one node
type PNSet[T comparable] struct {
	name string
	set  map[T]*PNCounter
}

func NewPNSet[T comparable](name string) *PNSet[T] {
	pnset := new(PNSet[T])
	pnset.name = name
	pnset.set = make(map[T]*PNCounter)
	return pnset
}

func (pnset *PNSet[T]) Add(value T) {
	if counter, exists := pnset.set[value]; exists {
		counter.Increment()
	} else {
		newCounter := NewPNCounter(pnset.name, pnset.name, []string{pnset.name})
		newCounter.Increment()
		pnset.set[value] = newCounter
	}
}

func (pnset *PNSet[T]) Remove(value T) {
	if counter, exists := pnset.set[value]; exists {
		counter.Decrement()
	}
}

func (pnset *PNSet[T]) Lookup(value T) bool {
	if counter, exists := pnset.set[value]; exists {
		return counter.Value() > 0
	}
	return false
}

func (pnset *PNSet[T]) Size() int {
	size := 0
	for _, counter := range pnset.set {
		if counter.Value() > 0 {
			size++
		}
	}
	return size
}

func (pnset *PNSet[T]) Merge(that *PNSet[T]) {
	for value, counter := range that.set {
		if _, exists := pnset.set[value]; !exists {
			pnset.set[value] = counter
		} else {
			pnset.set[value].Merge(counter)
		}
	}
}
