package goset

type number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~byte | ~float32 | ~float64
}

type set[T comparable] interface {
	// Add adds elements
	Add(v ...T)
	// Del deletes elements
	Del(v ...T)
	// Clear empties set
	Clear()

	// Length returns elements count
	Length() int

	// Has checks if v is int set
	Has(v T) bool

	// Copy returns a deep copy
	Copy() *set[T]

	// Equals checks if it has same elements with set t
	Equals(t *set[T]) bool
	// IsSub checks if it's sub set of set t
	IsSub(t *set[T]) bool

	// Union returns a union of itself and set t
	Union(t *set[T]) *set[T]
	// Intersect returns an intersection of itself and set t
	Intersect(t *set[T]) *set[T]
	// Subtract returns a subtraction of itself and set t
	Subtract(t *set[T]) *set[T]
	// Complement returns a complement of itself and set t
	Complement(t *set[T]) *set[T]
}
