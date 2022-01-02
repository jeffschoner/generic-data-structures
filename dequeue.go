package genericsds

import "errors"

var DequeueEmptyErr = errors.New("Dequeue is empty")

type Dequeue[T any] interface {
	Pop() (T, error)
	Push(item T)

	Dequeue() (T, error)
	Enqueue(item T)

	IsEmpty() bool
}
