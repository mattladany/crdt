package crdt

type TwoPhaseSet[T comparable] struct {
	name      string
	gset      *GSet[T]
	tombstone *GSet[T]
}

func NewTwoPhaseSet[T comparable](name string) *TwoPhaseSet[T] {
	twoPhaseSet := new(TwoPhaseSet[T])
	twoPhaseSet.name = name
	twoPhaseSet.gset = NewGSet[T](name)
	twoPhaseSet.tombstone = NewGSet[T](name)

	return twoPhaseSet
}

func (tpset *TwoPhaseSet[T]) Add(value T) {
	tpset.gset.Add(value)
}

func (tpset *TwoPhaseSet[T]) Remove(value T) {
	if tpset.gset.Lookup(value) {
		tpset.tombstone.Add(value)
	}
}

func (tpset *TwoPhaseSet[T]) Lookup(value T) bool {
	return tpset.gset.Lookup(value) && !tpset.tombstone.Lookup(value)
}

func (tpset *TwoPhaseSet[T]) Size() int {
	return tpset.gset.Size() - tpset.tombstone.Size()
}

func (tpset *TwoPhaseSet[T]) Merge(that *TwoPhaseSet[T]) {
	tpset.gset.Merge(that.gset)
	tpset.tombstone.Merge(that.tombstone)
}
