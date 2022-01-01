package operations

import (
	generics "github.com/jeffschoner/generic-data-structures"
)

func Filter[T any](input generics.Collection[T], f func(item T) bool, init generics.CollectionInit[T]) generics.Collection[T] {
	result := init()
	input.ForEach(func(item T) {
		if f(item) {
			result.Add(item)
		}
	})

	return result
}
