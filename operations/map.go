package operations

import (
	generics "github.com/jeffschoner/generic-data-structures"
)

func Map[TIn any, TOut any](input generics.Collection[TIn], init generics.CollectionInit[TOut], f func(item TIn) TOut) generics.Collection[TOut] {
	result := init()
	input.ForEach(func(item TIn) {
		out := f(item)
		result.Add(out)
	})

	return result
}
