package linkedlist

import (
	"errors"
)

type List[T any] struct {
	Head *node[T]
	Tail *node[T]
}

func New[T any]() *List[T] {
	return &List[T]{}
}

var EmptyError = errors.New("List is empty")

func (l *List[T]) Append(item T) {
	if l.Tail == nil {
		l.Tail = &node[T]{
			Datum:    item,
			Previous: nil,
			Next:     nil,
		}
		l.Head = l.Tail
	} else {
		l.Tail.Next = &node[T]{
			Datum:    item,
			Previous: l.Tail,
			Next:     nil,
		}
		l.Tail = l.Tail.Next
	}
}

func (l *List[T]) Prepend(item T) {
	if l.Head == nil {
		l.Head = &node[T]{
			Datum:    item,
			Previous: nil,
			Next:     nil,
		}
		l.Tail = l.Head
	} else {
		l.Head.Previous = &node[T]{
			Datum:    item,
			Previous: nil,
			Next:     l.Head,
		}
		l.Head = l.Head.Previous
	}
}

func (l List[T]) ForEach(f func(item T)) {
	for node := l.Head; node != nil; node = node.Next {
		f(node.Datum)
	}
}

func (l List[T]) Empty() bool {
	return l.Head == nil
}

func (l *List[T]) RemoveLast() (T, error) {
	if l.Tail == nil {
		var defaultValue T
		return defaultValue, EmptyError
	}

	toReturn := l.Tail.Datum

	l.Tail = l.Tail.Previous
	if l.Tail == nil {
		l.Head = nil
	} else {
		l.Tail.Next = nil
	}

	return toReturn, nil
}

func (l *List[T]) RemoveFirst() (T, error) {
	if l.Head == nil {
		var defaultValue T
		return defaultValue, EmptyError
	}

	toReturn := l.Head.Datum

	l.Head = l.Head.Next
	if l.Head == nil {
		l.Tail = nil
	} else {
		l.Head.Previous = nil
	}

	return toReturn, nil
}
