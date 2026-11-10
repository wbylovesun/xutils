package xmap

func ToSlice[K comparable, V any, T map[K]V, O any](t T, f func(K, V) O) []O {
	s := make([]O, 0)
	for k, v := range t {
		s = append(s, f(k, v))
	}
	return s
}

func Values[K comparable, V any, T map[K]V](t T) []V {
	s := make([]V, 0)
	for _, v := range t {
		s = append(s, v)
	}
	return s
}

func ConvertTo[K comparable, V any, T map[K]V, NK comparable, NV any](t T, f func(K, V) (NK, NV)) map[NK]NV {
	nm := make(map[NK]NV, 0)
	for k, v := range t {
		nk, nv := f(k, v)
		nm[nk] = nv
	}
	return nm
}
