package crdt

import (
	"github.com/google/uuid"
	"golang.org/x/exp/maps"
)

// ORSet is a CRDT set in which specific elements that are added to the set
// are given a unique ID that must be removed explicitly for an element to be
// perceived as not being in the set.
// The types of elements (T) must be set on initialization.
type ORSet[T comparable] struct {
	name      string
	set       map[T]map[string]bool
	tombstone map[T]map[string]bool
}

// NewORSet constructs an empty ORSet with the associated name.
// It is assumed the name of this specific ORSet uniquely identifies this
// set throughout the cluster.
func NewORSet[T comparable](name string) *ORSet[T] {
	orset := new(ORSet[T])
	orset.name = name
	orset.set = make(map[T]map[string]bool)
	orset.tombstone = make(map[T]map[string]bool)
	return orset
}

// Add adds the specified value to the set.
// Functionally, this creates a new unique ID and adds it to this element's
// List of currently added IDs.
func (orset *ORSet[T]) Add(value T) {
	uniqueId := uuid.NewString()
	if uniqueIds, exists := orset.set[value]; exists {
		uniqueIds[uniqueId] = true
		return
	}
	uniqueIds := make(map[string]bool)
	uniqueIds[uniqueId] = true
	orset.set[value] = uniqueIds
}

// Remove removes the specified value from the set.
// Functionally, this adds all unique IDs for the specified value in orset.set
// to orset.tombstone and removes them from orset.set.
func (orset *ORSet[T]) Remove(value T) {
	if existingIds, existsInSet := orset.set[value]; existsInSet {
		tombstoneIds := make(map[string]bool)
		if uniqueIdTombstone, existsInTombstone := orset.tombstone[value]; existsInTombstone {
			tombstoneIds = uniqueIdTombstone
		}

		maps.Copy(tombstoneIds, existingIds)
		orset.tombstone[value] = tombstoneIds
		delete(orset.set, value)
	}
}

// Lookup reports whether T value exists within the set.
func (orset *ORSet[T]) Lookup(value T) bool {
	if uniqueIds, exists := orset.set[value]; exists {
		return len(uniqueIds) > 0
	}
	return false
}

// Size returns the number of elements that currently exist within the set.
func (orset *ORSet[T]) Size() int {
	return len(orset.set)
}

// Merge performs the following three actions:
//   - Add all elements (and their unique IDs) from that.set to orset.set
//   - Add all elements (and their unique IDs) from that.tombstone to
//     orset.tombstone
//   - Remove all elements from orset.set if all their unique IDs exist within
//     orset.tombstone
//
// This is an idempotent operation and is a no-op if orset.name != that.name.
func (orset *ORSet[T]) Merge(that *ORSet[T]) {
	if orset.name != that.name {
		return
	}
	for value, thatUuids := range that.set {
		if uuids, exists := orset.set[value]; exists {
			maps.Copy(uuids, thatUuids)
		} else {
			orset.set[value] = thatUuids
		}
	}

	for value, thatUuids := range that.tombstone {
		if uuids, exists := orset.tombstone[value]; exists {
			maps.Copy(uuids, thatUuids)
		} else {
			orset.tombstone[value] = thatUuids
		}
	}

	maps.DeleteFunc(orset.set, func(value T, uuids map[string]bool) bool {
		tombstoneUuids, exists := orset.tombstone[value]
		if !exists {
			return false
		}
		for uuid := range uuids {
			if _, exists := tombstoneUuids[uuid]; !exists {
				return false
			}
		}
		return true
	})
}
