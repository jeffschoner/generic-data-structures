package genericsds

func FlatMap[TIn any, TOut any](input Collection[TIn], f func(item TIn) Collection[TOut], init CollectionInit[TOut]) Collection[TOut] {
	result := init()
	input.ForEach(func(item TIn) {
		out := f(item)
		out.ForEach(func(outItem TOut) {
			result.Add(outItem)
		})
	})

	return result
}
