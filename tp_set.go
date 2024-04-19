package crdt

// TwoPhaseSet is a CRDT set in which an element can be added and removed only
// once in the lifetime of this set.
// This is executed by having to GSets internally, one that holds elements
// that have been added, the other elements that have been removed.
// The types of elements (T) must be set on initialization.
type TwoPhaseSet[T comparable] struct {
	name      string
	gset      *GSet[T]
	tombstone *GSet[T]
}

// NewTwoPhaseSet constructs a TwoPhaseSet with the associated name.
// It is assumed the name of this specific TwoPhaseSet uniquely identifies
// this set throughout the cluster.
func NewTwoPhaseSet[T comparable](name string) *TwoPhaseSet[T] {
	twoPhaseSet := new(TwoPhaseSet[T])
	twoPhaseSet.name = name
	twoPhaseSet.gset = NewGSet[T](name)
	twoPhaseSet.tombstone = NewGSet[T](name)

	return twoPhaseSet
}

// Add adds the specified value to the set.
// Is a no-op if the value has ever previously been added to the set.
func (tpset *TwoPhaseSet[T]) Add(value T) {
	tpset.gset.Add(value)
}

// Remove adds the specified value to tpset.tombstone if it exists within
// tpset.gset.
// This functionally removes it from this TwoPhaseSet forever.
func (tpset *TwoPhaseSet[T]) Remove(value T) {
	if tpset.gset.Lookup(value) {
		tpset.tombstone.Add(value)
	}
}

// Lookup reports whether the T value exists within the set.
func (tpset *TwoPhaseSet[T]) Lookup(value T) bool {
	return tpset.gset.Lookup(value) && !tpset.tombstone.Lookup(value)
}

// Size returns the number of elements that currently exist within the set.
func (tpset *TwoPhaseSet[T]) Size() int {
	return tpset.gset.Size() - tpset.tombstone.Size()
}

// Merge calls Merge on both GSets within this TwoPhaseSet.
// This is an idempotent operation and is a no-op if tpset.name != that.name.
func (tpset *TwoPhaseSet[T]) Merge(that *TwoPhaseSet[T]) {
	if tpset.name == that.name {
		tpset.gset.Merge(that.gset)
		tpset.tombstone.Merge(that.tombstone)
	}
}
