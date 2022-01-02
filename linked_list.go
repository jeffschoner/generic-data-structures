package genericsds

import (
	"errors"
)

type LinkedList[T any] struct {
	Head *linkedListNode[T]
	Tail *linkedListNode[T]

	size int
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func NewListCollection[T any]() Collection[T] {
	return NewLinkedList[T]()
}

var EmptyError = errors.New("List is empty")

type linkedListNode[T any] struct {
	Datum    T
	Previous *linkedListNode[T]
	Next     *linkedListNode[T]
}

func (l *LinkedList[T]) Append(item T) {
	if l.Tail == nil {
		l.Tail = &linkedListNode[T]{
			Datum:    item,
			Previous: nil,
			Next:     nil,
		}
		l.Head = l.Tail
	} else {
		l.Tail.Next = &linkedListNode[T]{
			Datum:    item,
			Previous: l.Tail,
			Next:     nil,
		}
		l.Tail = l.Tail.Next
	}

	l.size++
}

func (l *LinkedList[T]) Add(item T) {
	l.Append(item)
}

func (l *LinkedList[T]) Prepend(item T) {
	if l.Head == nil {
		l.Head = &linkedListNode[T]{
			Datum:    item,
			Previous: nil,
			Next:     nil,
		}
		l.Tail = l.Head
	} else {
		l.Head.Previous = &linkedListNode[T]{
			Datum:    item,
			Previous: nil,
			Next:     l.Head,
		}
		l.Head = l.Head.Previous
	}

	l.size++
}

func (l LinkedList[T]) ForEach(f func(item T)) {
	for node := l.Head; node != nil; node = node.Next {
		f(node.Datum)
	}
}

func (l LinkedList[T]) IsEmpty() bool {
	return l.Head == nil
}

func (l LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) RemoveLast() (T, error) {
	if l.Tail == nil {
		var defaultValue T
		return defaultValue, EmptyError
	}

	l.size--

	toReturn := l.Tail.Datum

	l.Tail = l.Tail.Previous
	if l.Tail == nil {
		l.Head = nil
	} else {
		l.Tail.Next = nil
	}

	return toReturn, nil
}

func (l *LinkedList[T]) RemoveFirst() (T, error) {
	if l.Head == nil {
		var defaultValue T
		return defaultValue, EmptyError
	}

	l.size--

	toReturn := l.Head.Datum

	l.Head = l.Head.Next
	if l.Head == nil {
		l.Tail = nil
	} else {
		l.Head.Previous = nil
	}

	return toReturn, nil
}
