package sliceMethods

type Integer interface {
	Signed | Unsigned
}

type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Ordered interface {
	Integer | Float | ~string
}

type Float interface {
	~float32 | ~float64
}

// basic methods similar to ruby and js for golang
// all methods take a pointer to a slice
// important with map to have a callback that returns any
func Filter[T Ordered](s *[]T, callback func(T) bool) []T {
	newSlice := make([]T, 0)

	for _, v := range *s {
		if callback(v) {
			newSlice = append(newSlice, v)
		}
	}

	return newSlice
}

func Map[T Ordered](s *[]T, callback func(T) any) []any  {
	newSlice := make([]any, len(*s))

	for i, v := range *s {
		newVal := callback(v)
		newSlice[i] = newVal
	}

	return newSlice
}

func Reduce[T Ordered](s *[]T, callback func(T, T) T, acc T) T {
	for _, v := range *s {
		acc = callback(v, acc)
	}

	return acc
}
