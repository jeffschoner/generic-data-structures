package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	l := New[int]()
	assert.True(t, l.Empty())
}

func TestAppendNotEmpty(t *testing.T) {
	l := New[int]()
	l.Append(23)
	assert.False(t, l.Empty())
}

func TestPrependNotEmpty(t *testing.T) {
	l := New[int]()
	l.Prepend(23)
	assert.False(t, l.Empty())
}

func TestRemoveFirstToEmpty(t *testing.T) {
	l := New[int]()
	l.Append(23)
	l.Append(50)
	assert.False(t, l.Empty())

	removed, err := l.RemoveFirst()
	assert.NoError(t, err)
	assert.Equal(t, 23, removed)
	assert.False(t, l.Empty())

	removed, err = l.RemoveFirst()
	assert.NoError(t, err)
	assert.Equal(t, 50, removed)
	assert.True(t, l.Empty())

	_, err = l.RemoveFirst()
	assert.Error(t, err)
}

func TestRemoveLastToEmpty(t *testing.T) {
	l := New[int]()
	l.Append(23)
	l.Append(50)
	assert.False(t, l.Empty())

	removed, err := l.RemoveLast()
	assert.NoError(t, err)
	assert.Equal(t, 50, removed)
	assert.False(t, l.Empty())

	removed, err = l.RemoveLast()
	assert.NoError(t, err)
	assert.Equal(t, 23, removed)
	assert.True(t, l.Empty())

	_, err = l.RemoveLast()
	assert.Error(t, err)
}

func TestForEach(t *testing.T) {
	l := New[string]()
	l.Append("foo")
	l.Prepend("bar")
	l.Append("baz")

	var order []string
	l.ForEach(func(item string) {
		order = append(order, item)
	})

	assert.ElementsMatch(t, []string{"bar", "foo", "baz"}, order)
}
