package crdt

import (
	"time"
)

// LWWESet is a CRDT set in which each element has a timestamp associated with
// it.
// A last-writer-wins implementation is used to determine whether an element
// exists within the set or not.
// The integer `time.Now().UnixNano()` is used to identify a total ordering of
// updates and is assumed to be universal across all replicas of this LWWESet.
// The types of elements (T) must be set on initialization.
type LWWESet[T comparable] struct {
	name      string
	addSet    map[T]int64
	removeSet map[T]int64
}

// NewLLWESet constructs an empty LWWESet with the associated name.
// It is assumed the name of this specific LWWESet uniquely identifies this
// set throughout the cluster.
func NewLWWESet[T comparable](name string) *LWWESet[T] {
	lwweSet := new(LWWESet[T])
	lwweSet.name = name
	lwweSet.addSet = make(map[T]int64)
	lwweSet.removeSet = make(map[T]int64)
	return lwweSet
}

// Add adds the specified value to the set (i.e., its timestamp in the addSet
// gets updated to now).
func (lwweSet *LWWESet[T]) Add(value T) {
	lwweSet.addSet[value] = time.Now().UnixNano()
}

// Remove removes the specified value from the set  (i.e., its timestamp in
// the removeSet gets updated to now).
// This is a no-op if the value has never been added to the set previously.
func (lwweSet *LWWESet[T]) Remove(value T) {
	if _, exists := lwweSet.addSet[value]; exists {
		lwweSet.removeSet[value] = time.Now().UnixNano()
	}
}

// Lookup reports whether the T value exists within the set.
// An element exists within the set if its timestamp in the addSet is greater
// than its timestamp in the removeSet.
func (lwweSet *LWWESet[T]) Lookup(value T) bool {
	if addSetValue, existsInAdd := lwweSet.addSet[value]; existsInAdd {
		removeSetValue, existsInRemove := lwweSet.removeSet[value]
		if !existsInRemove || removeSetValue < addSetValue {
			return true
		}
	}
	return false
}

// Size returns the number of elements currently within the set.
func (lwweSet *LWWESet[T]) Size() int {
	size := 0
	for value, addTimestamp := range lwweSet.addSet {
		removeTimestamp, existsInRemove := lwweSet.removeSet[value]
		if !existsInRemove || removeTimestamp < addTimestamp {
			size++
		}
	}
	return size
}

// Merge adds all elements in that.addSet and that.removeSet to lwweSet.addSet
// and llweSet.removeSet respectively, accounting for colissions by storing
// the most recent of the two timestamps.
// This is an idempotent operation and is a no-op if lwweSet.name != that.name.
func (lwweSet *LWWESet[T]) Merge(that *LWWESet[T]) {
	if lwweSet.name != that.name {
		return
	}
	// Merge addSet
	for thatValue, thatTimestamp := range that.addSet {
		if thisTimestamp, thisExists := lwweSet.addSet[thatValue]; thisExists {
			if thisTimestamp > thatTimestamp {
				continue
			}
		}
		lwweSet.addSet[thatValue] = thatTimestamp
	}
	// Merge removeSet
	for thatValue, thatTimestamp := range that.removeSet {
		if thisTimestamp, thisExists := lwweSet.removeSet[thatValue]; thisExists {
			if thisTimestamp > thatTimestamp {
				continue
			}
		}
		lwweSet.removeSet[thatValue] = thatTimestamp
	}
}
