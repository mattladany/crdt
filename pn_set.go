package crdt

// PNSet is a CRDT set in which each element can be added/removed from the set.
// Uses a PNCounter that is assumed to have only one node to abstract
// additions and removals of elements from the set.
// Note, it is possible for a no-op add to occur if two nodes remove the same
// element from the set and are then merged again, making the counter equal to
// -1.
// The types of elements (T) must be set on initialization.
type PNSet[T comparable] struct {
	name string
	set  map[T]*PNCounter
}

// NewPNSet constructs a PNSet with the associated name.
// It is assumed the name of this specific PNSet uniquely identifies this
// set throughout the cluster.
func NewPNSet[T comparable](name string) *PNSet[T] {
	pnset := new(PNSet[T])
	pnset.name = name
	pnset.set = make(map[T]*PNCounter)
	return pnset
}

// Add adds the specified value to the set (i.e., its counter is incremented
// by 1).
func (pnset *PNSet[T]) Add(value T) {
	if counter, exists := pnset.set[value]; exists {
		counter.Increment()
	} else {
		newCounter := NewPNCounter(pnset.name, pnset.name, []string{pnset.name})
		newCounter.Increment()
		pnset.set[value] = newCounter
	}
}

// Remove removes the specified value from the set (i.e., its counter is
// decremented by 1).
// This is functionally a no-op if value does not exist in the set.
func (pnset *PNSet[T]) Remove(value T) {
	if counter, exists := pnset.set[value]; exists {
		counter.Decrement()
	}
}

// Lookup reports whether the T value exists within the set.
func (pnset *PNSet[T]) Lookup(value T) bool {
	if counter, exists := pnset.set[value]; exists {
		return counter.Value() > 0
	}
	return false
}

// Size returns the number of elements that currently exist within the set.
func (pnset *PNSet[T]) Size() int {
	size := 0
	for _, counter := range pnset.set {
		if counter.Value() > 0 {
			size++
		}
	}
	return size
}

// Merge adds all elements in that.set to pnset.set that did not exist
// previously, and merges the counters for each element that exists in both.
// This is an idempotent operation and is a no-op if pnset.name != that.name.
func (pnset *PNSet[T]) Merge(that *PNSet[T]) {
	if pnset.name != that.name {
		return
	}
	for value, counter := range that.set {
		if _, exists := pnset.set[value]; !exists {
			pnset.set[value] = counter
		} else {
			pnset.set[value].Merge(counter)
		}
	}
}
