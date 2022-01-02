package genericsds

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMapEmpty(t *testing.T) {
	l := NewLinkedList[int]()

	lOut := Map[int](l, NewListCollection[int], func(item int) int {
		return item * item
	})

	var collect []int
	lOut.ForEach(func(item int) {
		collect = append(collect, item)
	})
	assert.Equal(t, 0, len(collect))
}

func TestMapList(t *testing.T) {
	l := NewLinkedList[int]()
	l.Add(1)
	l.Add(2)
	l.Add(3)

	lOut := Map[int](l, NewListCollection[int], func(item int) int {
		return item * item
	})

	var collect []int
	lOut.ForEach(func(item int) {
		collect = append(collect, item)
	})
	assert.ElementsMatch(t, []int{1, 4, 9}, collect)
}

func TestMapSet(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	s.Add(2)
	s.Add(3)

	sOut := Map[int](s, NewSetCollection[string], func(item int) string {
		return strconv.Itoa(item)
	})

	var collect []string
	sOut.ForEach(func(item string) {
		collect = append(collect, item)
	})
	assert.ElementsMatch(t, []string{"1", "2", "3"}, collect)
}
