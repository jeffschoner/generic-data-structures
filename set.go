package genericsds

type Set[T comparable] struct {
	storage map[T]struct{}
}

// Create a new generic-typed set. T must be comparable. You need to use this function
// to create new sets. Set[int]{} will not work.
func NewSet[T comparable]() Set[T] {
	return Set[T]{
		storage: make(map[T]struct{}),
	}
}

func NewSetCollection[T comparable]() Collection[T] {
	return NewSet[T]()
}

// Add an item to the set. Has no effect if the item already exists.
func (r Set[T]) Add(item T) {
	r.storage[item] = struct{}{}
}

// Determine if the set is empty
func (r Set[T]) IsEmpty() bool {
	return len(r.storage) == 0
}

// Returns true if the set contains the specified item.
func (r Set[T]) Contains(item T) bool {
	_, has := r.storage[item]
	return has
}

// Removes the specified item from the set. Has no effect if the item is not present.
func (r Set[T]) Remove(item T) {
	delete(r.storage, item)
}

// Returns the number of items in the set.
func (r Set[T]) Size() int {
	return len(r.storage)
}

// Adds all items from the specified set to this one. This function mutates the set.
func (r Set[T]) Merge(s Set[T]) {
	for k := range s.storage {
		r.Add(k)
	}
}

// Iterate through all items, call the delegate function on each
func (r Set[T]) ForEach(f func(item T)) {
	for k := range r.storage {
		f(k)
	}
}

// Creates a new set from all items in this one and the one specified. It is
// non-destructive to both sets involved.
func (r Set[T]) Union(s Set[T]) Set[T] {
	newSet := NewSet[T]()

	for k := range r.storage {
		newSet.Add(k)
	}

	for k := range s.storage {
		newSet.Add(k)
	}

	return newSet
}

// Creates a new set with all items common to this one and the one specified. It is
// non-destructive to both sets involved.
func (r Set[T]) Intersection(s Set[T]) Set[T] {
	newSet := NewSet[T]()

	var shorterSet *Set[T]
	var longerSet *Set[T]

	if r.Size() > s.Size() {
		shorterSet = &s
		longerSet = &r
	} else {
		shorterSet = &r
		longerSet = &s
	}

	for k := range shorterSet.storage {
		if longerSet.Contains(k) {
			newSet.Add(k)
		}
	}

	return newSet
}

// Creates a new set with the items in this one less any of the items in the specified
// set. It is non-destructive to both sets involved.
func (r Set[T]) Difference(s Set[T]) Set[T] {
	newSet := NewSet[T]()

	for k := range r.storage {
		newSet.Add(k)
	}

	for k := range s.storage {
		newSet.Remove(k)
	}

	return newSet
}
