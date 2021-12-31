package operations

import (
	"github.com/jeffschoner/generic-data-structures/collection"
)

func FlatMap[TIn any, TOut any](input collection.Collection[TIn], f func(item TIn) collection.Collection[TOut], init collection.CollectionInit[TOut]) collection.Collection[TOut] {
	result := init()
	input.ForEach(func(item TIn) {
		out := f(item)
		out.ForEach(func(outItem TOut) {
			result.Add(outItem)
		})
	})

	return result
}
