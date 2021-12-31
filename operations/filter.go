package operations

import (
	"github.com/jeffschoner/generic-data-structures/collection"
)

func Filter[T any](input collection.Collection[T], f func(item T) bool, init collection.CollectionInit[T]) collection.Collection[T] {
	result := init()
	input.ForEach(func(item T) {
		if f(item) {
			result.Add(item)
		}
	})

	return result
}
