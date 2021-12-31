package collection

type Collection[T any] interface {
	ForEach(f func(item T))
	Add(item T)
}

type CollectionInit[T any] func() Collection[T]
