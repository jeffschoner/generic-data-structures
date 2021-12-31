package operations

import (
	"github.com/jeffschoner/generic-data-structures/collection"
)

func Map[TIn any, TOut any](input collection.Collection[TIn], init collection.CollectionInit[TOut], f func(item TIn) TOut) collection.Collection[TOut] {
	result := init()
	input.ForEach(func(item TIn) {
		out := f(item)
		result.Add(out)
	})

	return result
}
