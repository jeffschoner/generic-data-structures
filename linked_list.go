package genericsds

import (
	"errors"
)

type LinkedList[T any] struct {
	head *linkedListNode[T]
	tail *linkedListNode[T]

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
	if l.tail == nil {
		l.tail = &linkedListNode[T]{
			Datum:    item,
			Previous: nil,
			Next:     nil,
		}
		l.head = l.tail
	} else {
		l.tail.Next = &linkedListNode[T]{
			Datum:    item,
			Previous: l.tail,
			Next:     nil,
		}
		l.tail = l.tail.Next
	}

	l.size++
}

func (l *LinkedList[T]) Add(item T) {
	l.Append(item)
}

func (l *LinkedList[T]) Prepend(item T) {
	if l.head == nil {
		l.head = &linkedListNode[T]{
			Datum:    item,
			Previous: nil,
			Next:     nil,
		}
		l.tail = l.head
	} else {
		l.head.Previous = &linkedListNode[T]{
			Datum:    item,
			Previous: nil,
			Next:     l.head,
		}
		l.head = l.head.Previous
	}

	l.size++
}

func (l LinkedList[T]) ForEach(f func(item T)) {
	for node := l.head; node != nil; node = node.Next {
		f(node.Datum)
	}
}

func (l LinkedList[T]) IsEmpty() bool {
	return l.head == nil
}

func (l LinkedList[T]) Size() int {
	return l.size
}

func (l *LinkedList[T]) RemoveLast() (T, error) {
	if l.tail == nil {
		var defaultValue T
		return defaultValue, EmptyError
	}

	l.size--

	toReturn := l.tail.Datum

	l.tail = l.tail.Previous
	if l.tail == nil {
		l.head = nil
	} else {
		l.tail.Next = nil
	}

	return toReturn, nil
}

func (l *LinkedList[T]) RemoveFirst() (T, error) {
	if l.head == nil {
		var defaultValue T
		return defaultValue, EmptyError
	}

	l.size--

	toReturn := l.head.Datum

	l.head = l.head.Next
	if l.head == nil {
		l.tail = nil
	} else {
		l.head.Previous = nil
	}

	return toReturn, nil
}
