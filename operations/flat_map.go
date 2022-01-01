package operations

import (
	generics "github.com/jeffschoner/generic-data-structures"
)

func FlatMap[TIn any, TOut any](input generics.Collection[TIn], f func(item TIn) generics.Collection[TOut], init generics.CollectionInit[TOut]) generics.Collection[TOut] {
	result := init()
	input.ForEach(func(item TIn) {
		out := f(item)
		out.ForEach(func(outItem TOut) {
			result.Add(outItem)
		})
	})

	return result
}
