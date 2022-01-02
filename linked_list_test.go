package genericsds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	l := NewLinkedList[int]()
	assert.True(t, l.IsEmpty())
}

func TestAppendNotEmpty(t *testing.T) {
	l := NewLinkedList[int]()
	l.Append(23)
	assert.False(t, l.IsEmpty())
}

func TestPrependNotEmpty(t *testing.T) {
	l := NewLinkedList[int]()
	l.Prepend(23)
	assert.False(t, l.IsEmpty())
}

func TestRemoveFirstToEmpty(t *testing.T) {
	l := NewLinkedList[int]()
	l.Append(23)
	l.Append(50)
	assert.False(t, l.IsEmpty())

	removed, err := l.RemoveFirst()
	assert.NoError(t, err)
	assert.Equal(t, 23, removed)
	assert.False(t, l.IsEmpty())

	removed, err = l.RemoveFirst()
	assert.NoError(t, err)
	assert.Equal(t, 50, removed)
	assert.True(t, l.IsEmpty())

	_, err = l.RemoveFirst()
	assert.Error(t, err)
}

func TestRemoveLastToEmpty(t *testing.T) {
	l := NewLinkedList[int]()
	l.Append(23)
	l.Append(50)
	assert.False(t, l.IsEmpty())

	removed, err := l.RemoveLast()
	assert.NoError(t, err)
	assert.Equal(t, 50, removed)
	assert.False(t, l.IsEmpty())

	removed, err = l.RemoveLast()
	assert.NoError(t, err)
	assert.Equal(t, 23, removed)
	assert.True(t, l.IsEmpty())

	_, err = l.RemoveLast()
	assert.Error(t, err)
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
