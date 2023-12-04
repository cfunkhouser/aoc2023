package util

// Pointy returns a pointer to v, regardless of type.
func Pointy[T any](v T) *T {
	return &v
}

type numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

// Sum of all elements in a slice of numeric values.
func Sum[N numeric](s []N) (ret N) {
	for _, v := range s {
		ret += v
	}
	return
}
