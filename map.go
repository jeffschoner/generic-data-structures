package genericsds

func Map[TIn any, TOut any](input Collection[TIn], init CollectionInit[TOut], f func(item TIn) TOut) Collection[TOut] {
	result := init()
	input.ForEach(func(item TIn) {
		out := f(item)
		result.Add(out)
	})

	return result
}
