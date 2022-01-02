package genericsds

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyInt(t *testing.T) {
	s := NewSet[int]()
	assert.Equal(t, 0, s.Size())
	assert.True(t, s.IsEmpty())
}

func TestAddInt(t *testing.T) {
	s := NewSet[int]()
	s.Add(23)
	s.Add(-22)
	s.Add(0)
	assert.Equal(t, 3, s.Size())
	assert.True(t, s.Contains(23))
	assert.True(t, s.Contains(-22))
	assert.True(t, s.Contains(0))
	assert.False(t, s.IsEmpty())
}

func TestAddDupeInt(t *testing.T) {
	s := NewSet[int]()
	s.Add(23)
	s.Add(0)
	s.Add(23)
	assert.Equal(t, 2, s.Size())
	assert.True(t, s.Contains(23))
	assert.True(t, s.Contains(0))
}

func TestRemoveInt(t *testing.T) {
	s := NewSet[int]()
	s.Add(23)
	s.Add(-22)
	s.Add(0)
	s.Remove(0)
	assert.Equal(t, 2, s.Size())
	assert.True(t, s.Contains(23))
	assert.True(t, s.Contains(-22))
}

func TestRemoveUnknownInt(t *testing.T) {
	s := NewSet[int]()
	s.Add(23)
	s.Remove(12)
	assert.Equal(t, 1, s.Size())
	assert.True(t, s.Contains(23))
	assert.False(t, s.Contains(12))
}

func TestString(t *testing.T) {
	s := NewSet[string]()
	s.Add("hello")
	s.Add("foo")
	assert.Equal(t, 2, s.Size())
	assert.True(t, s.Contains("hello"))
	assert.True(t, s.Contains("foo"))

	s.Remove("unknown")
	assert.Equal(t, 2, s.Size())
	assert.False(t, s.Contains("unknown"))
}

func TestMerge(t *testing.T) {
	s1 := NewSet[string]()
	s1.Add("hi")
	s1.Add("foo")

	s2 := NewSet[string]()
	s2.Add("bar")
	s2.Add("hi")
	s2.Add("baz")

	s1.Merge(s2)

	assert.Equal(t, 3, s2.Size())

	assert.Equal(t, 4, s1.Size())
	assert.True(t, s1.Contains("foo"))
	assert.True(t, s1.Contains("bar"))
	assert.True(t, s1.Contains("hi"))
	assert.True(t, s1.Contains("baz"))
}

func TestUnion(t *testing.T) {
	s1 := NewSet[string]()
	s1.Add("hi")
	s1.Add("foo")

	s2 := NewSet[string]()
	s2.Add("bar")
	s2.Add("hi")
	s2.Add("baz")

	s3 := s1.Union(s2)

	assert.Equal(t, 2, s1.Size())
	assert.Equal(t, 3, s2.Size())

	assert.Equal(t, 4, s3.Size())
	assert.True(t, s3.Contains("foo"))
	assert.True(t, s3.Contains("bar"))
	assert.True(t, s3.Contains("hi"))
	assert.True(t, s3.Contains("baz"))
}

func TestIntersection(t *testing.T) {
	s1 := NewSet[string]()
	s1.Add("hi")
	s1.Add("foo")

	s2 := NewSet[string]()
	s2.Add("bar")
	s2.Add("hi")
	s2.Add("baz")

	s3 := s1.Intersection(s2)

	assert.Equal(t, 2, s1.Size())
	assert.Equal(t, 3, s2.Size())

	assert.Equal(t, 1, s3.Size())
	assert.True(t, s3.Contains("hi"))
}

func TestIntersectionFlipLong(t *testing.T) {
	s1 := NewSet[string]()
	s1.Add("hi")
	s1.Add("foo")
	s1.Add("baz")

	s2 := NewSet[string]()
	s2.Add("bar")
	s2.Add("hi")

	s3 := s1.Intersection(s2)

	assert.Equal(t, 3, s1.Size())
	assert.Equal(t, 2, s2.Size())

	assert.Equal(t, 1, s3.Size())
	assert.True(t, s3.Contains("hi"))
}

func TestDifference(t *testing.T) {
	s1 := NewSet[string]()
	s1.Add("hi")
	s1.Add("foo")

	s2 := NewSet[string]()
	s2.Add("bar")
	s2.Add("hi")
	s2.Add("baz")

	s3 := s2.Difference(s1)

	assert.Equal(t, 2, s1.Size())
	assert.Equal(t, 3, s2.Size())

	assert.Equal(t, 2, s3.Size())
	assert.True(t, s3.Contains("bar"))
	assert.True(t, s3.Contains("baz"))
}
