package crdt

import (
	"time"
)

type LWWESet[T comparable] struct {
	name      string
	addSet    map[T]int64
	removeSet map[T]int64
}

func NewLWWESet[T comparable](name string) *LWWESet[T] {
	lwweSet := new(LWWESet[T])
	lwweSet.name = name
	lwweSet.addSet = make(map[T]int64)
	lwweSet.removeSet = make(map[T]int64)
	return lwweSet
}

func (lwweSet *LWWESet[T]) Add(value T) {
	lwweSet.addSet[value] = time.Now().UnixNano()
}

func (lwweSet *LWWESet[T]) Remove(value T) {
	if _, exists := lwweSet.addSet[value]; exists {
		lwweSet.removeSet[value] = time.Now().UnixNano()
	}
}

func (lwweSet *LWWESet[T]) Lookup(value T) bool {
	if addSetValue, existsInAdd := lwweSet.addSet[value]; existsInAdd {
		removeSetValue, existsInRemove := lwweSet.removeSet[value]
		if !existsInRemove || removeSetValue < addSetValue {
			return true
		}
	}
	return false
}

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

func (lwweSet *LWWESet[T]) Merge(that *LWWESet[T]) {
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
