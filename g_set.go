package crdt

// GSet is a CRDT set in which elements can only be added to it.
// The types of elements (T) must be set on initialization.
type GSet[T comparable] struct {
	name string
	set  map[T]bool
}

// NewGSet constructs a GSet with the associated name.
// It is assumed the name of this specific GSet uniquely identifies this
// set throughout the cluster.
func NewGSet[T comparable](name string) *GSet[T] {
	gset := new(GSet[T])
	gset.name = name
	gset.set = make(map[T]bool)
	return gset
}

// Add adds the specified value to the set.
// Is a no-op if the value has ever previously been added to the set.
func (gset *GSet[T]) Add(value T) {
	gset.set[value] = true
}

// Lookup reports whether the T value exists within the set.
func (gset *GSet[T]) Lookup(value T) (exists bool) {
	_, exists = gset.set[value]
	return
}

// Size returns number of elements that currently exist within the set.
func (gset *GSet[T]) Size() int {
	return len(gset.set)
}

// Merge adds all elements in that.set to gset.set.
// This is an idempotent operation and is a no-op if gset.name != that.name.
func (gset *GSet[T]) Merge(that *GSet[T]) {
	if gset.name == that.name {
		for key := range that.set {
			gset.set[key] = true
		}
	}
}
