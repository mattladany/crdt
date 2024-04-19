package crdt

import (
	"github.com/google/uuid"
	"golang.org/x/exp/maps"
)

type ORSet[T comparable] struct {
	name      string
	set       map[T]map[string]bool
	tombstone map[T]map[string]bool
}

func NewORSet[T comparable](name string) *ORSet[T] {
	orset := new(ORSet[T])
	orset.name = name
	orset.set = make(map[T]map[string]bool)
	orset.tombstone = make(map[T]map[string]bool)
	return orset
}

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

func (orset *ORSet[T]) Lookup(value T) bool {
	if uniqueIds, exists := orset.set[value]; exists {
		return len(uniqueIds) > 0
	}
	return false
}

func (orset *ORSet[T]) Size() int {
	return len(orset.set)
}

func (orset *ORSet[T]) Merge(that *ORSet[T]) {

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
