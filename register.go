package crdt

type Register[T any] interface {
	Get() T
	Set(T)
	Merge(*Register[T])
}
