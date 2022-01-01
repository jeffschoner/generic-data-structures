package generics

type Collection[T any] interface {
	ForEach(f func(item T))
	Add(item T)
	IsEmpty() bool
}

type CollectionInit[T any] func() Collection[T]
