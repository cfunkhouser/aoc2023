package util

import "slices"

// Pointy returns a pointer to v, regardless of type.
func Pointy[T any](v T) *T {
	return &v
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Sum of all elements in a slice of numeric values.
func Sum[N Number](s []N) (ret N) {
	for _, v := range s {
		ret += v
	}
	return
}

// Set of values.
type Set[N Number] struct {
	values map[N]bool
}

// Values from the set as a slice.
func (s *Set[N]) Values() (ret []N) {
	for v := range s.values {
		ret = append(ret, v)
	}
	slices.SortStableFunc(ret, func(l, r N) int {
		return int(l - r)
	})
	return
}

// Intersection of two sets is the values which appear in each.
func (s *Set[N]) Intersection(other *Set[N]) *Set[N] {
	var ret []N
	for v := range other.values {
		if s.values[v] {
			ret = append(ret, v)
		}
	}
	return NewSet[N](ret)
}

// NewSet creates a new set populate with all distinct values.
func NewSet[N Number](values []N) *Set[N] {
	ret := &Set[N]{
		values: make(map[N]bool),
	}
	for _, v := range values {
		ret.values[v] = true
	}
	return ret
}
