package xmap

func ToSlice[K comparable, V any, T map[K]V, O any](t T, f func(K, V) O) []O {
	slice := make([]O, 0)
	for k, v := range t {
		slice = append(slice, f(k, v))
	}
	return slice
}
