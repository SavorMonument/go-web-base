package util

func SetToSlice[T, V comparable](elems map[T]V) []T {
	slice := make([]T, 0)
	for k, _ := range elems {
		slice = append(slice, k)
	}

	return slice
}

func KeysToSlice[T any, V comparable](elems map[V]T) []V {
	slice := make([]V, 0)
	for key := range elems {
		slice = append(slice, key)
	}

	return slice
}

func ValuesToSlice[T any, V comparable](elems map[V]T) []T {
	slice := make([]T, 0)
	for _, v := range elems {
		slice = append(slice, v)
	}

	return slice
}

func SliceToMap[T any, V comparable](elems []T, getKey func(T) V) map[V]T {
	result := make(map[V]T)
	for _, e := range elems {
		result[getKey(e)] = e
	}

	return result
}

func Contains[T comparable](elems []T, value T) bool {
	for _, elem := range elems {
		if value == elem {
			return true
		}
	}

	return false
}

func ContainsFunc[T any](slice []T, condition func(T) bool) bool {
	for _, v := range slice {
		if condition(v) {
			return true
		}
	}
	return false
}

func Find[T comparable](elems []T, f func(T) bool) T {

	for idx, elem := range elems {
		if f(elem) {
			return elems[idx]
		}
	}
	var e T
	return e
}

func SliceIndexFunc[T comparable](elems []T, f func(T) bool) int {
	for idx, elem := range elems {
		if f(elem) {
			return idx
		}
	}
	return -1
}

func Index[T comparable](elems []T, value T) int {
	for idx, elem := range elems {
		if value == elem {
			return idx
		}
	}
	return -1
}

func Remove[T comparable](elems []T, value T) []T {
	for idx, elem := range elems {
		if value == elem {
			return append(elems[:idx], elems[idx+1:]...)
		}
	}

	return elems
}

func AppendIfNotPresent[T comparable](elems []T, value T) []T {
	if !Contains(elems, value) {
		return append(elems, value)
	}
	return elems
}

func Filter[T any](elems []T, f func(T) bool) []T {
	var result []T
	for _, e := range elems {
		if f(e) {
			result = append(result, e)
		}
	}
	return result
}

func FindFirst[T comparable](elems []T, f func(*T) bool) *T {
	for _, e := range elems {
		if f(&e) {
			return &e
		}
	}
	return nil
}

func Map[T any, V any](elems []T, f func(T) V) []V {
	var result []V
	for _, e := range elems {
		result = append(result, f(e))
	}
	return result
}

