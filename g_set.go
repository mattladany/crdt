package crdt

type GSet[T comparable] struct {
	name string
	set  map[T]bool
}

func NewGSet[T comparable](name string) *GSet[T] {
	gset := new(GSet[T])
	gset.name = name
	gset.set = make(map[T]bool)
	return gset
}

func (gset *GSet[T]) Add(value T) {
	gset.set[value] = true
}

func (gset *GSet[T]) Lookup(value T) (exists bool) {
	_, exists = gset.set[value]
	return
}

func (gset *GSet[T]) Size() int {
	return len(gset.set)
}

func (gset *GSet[T]) Merge(that *GSet[T]) {
	for key := range that.set {
		gset.set[key] = true
	}
}
