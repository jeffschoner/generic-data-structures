package genericsds

func Filter[T any](input Collection[T], f func(item T) bool, init CollectionInit[T]) Collection[T] {
	result := init()
	input.ForEach(func(item T) {
		if f(item) {
			result.Add(item)
		}
	})

	return result
}
