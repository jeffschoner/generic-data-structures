package genericsds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	l := NewLinkedList[int]()
	assert.True(t, l.IsEmpty())
	assert.Equal(t, 0, l.Size())
}

func TestAppendNotEmpty(t *testing.T) {
	l := NewLinkedList[int]()
	l.Append(23)
	assert.False(t, l.IsEmpty())
	assert.Equal(t, 1, l.Size())
}

func TestPrependNotEmpty(t *testing.T) {
	l := NewLinkedList[int]()
	l.Prepend(23)
	assert.False(t, l.IsEmpty())
	assert.Equal(t, 1, l.Size())
}

func TestRemoveFirstToEmpty(t *testing.T) {
	l := NewLinkedList[int]()
	l.Append(23)
	l.Append(50)
	assert.False(t, l.IsEmpty())
	assert.Equal(t, 2, l.Size())

	removed, err := l.RemoveFirst()
	assert.NoError(t, err)
	assert.Equal(t, 23, removed)
	assert.False(t, l.IsEmpty())
	assert.Equal(t, 1, l.Size())

	removed, err = l.RemoveFirst()
	assert.NoError(t, err)
	assert.Equal(t, 50, removed)
	assert.True(t, l.IsEmpty())
	assert.Equal(t, 0, l.Size())

	_, err = l.RemoveFirst()
	assert.Error(t, err)
	assert.Equal(t, 0, l.Size())
}

func TestRemoveLastToEmpty(t *testing.T) {
	l := NewLinkedList[int]()
	l.Append(23)
	l.Append(50)
	assert.False(t, l.IsEmpty())
	assert.Equal(t, 2, l.Size())

	removed, err := l.RemoveLast()
	assert.NoError(t, err)
	assert.Equal(t, 50, removed)
	assert.False(t, l.IsEmpty())
	assert.Equal(t, 1, l.Size())

	removed, err = l.RemoveLast()
	assert.NoError(t, err)
	assert.Equal(t, 23, removed)
	assert.True(t, l.IsEmpty())
	assert.Equal(t, 0, l.Size())

	_, err = l.RemoveLast()
	assert.Error(t, err)
	assert.Equal(t, 0, l.Size())
}

func TestForEach(t *testing.T) {
	l := NewLinkedList[string]()
	l.Append("foo")
	l.Prepend("bar")
	l.Append("baz")

	var order []string
	l.ForEach(func(item string) {
		order = append(order, item)
	})

	assert.ElementsMatch(t, []string{"bar", "foo", "baz"}, order)
}
