package genericsds

type Collection[T any] interface {
	ForEach(f func(item T))
	Add(item T)
	IsEmpty() bool
	Size() int
}

type CollectionInit[T any] func() Collection[T]
