package xslice

func Range(from, to, delta int) []int {
	var rng []int
	if from > to {
		return rng
	}
	for i := from; i <= to; i += delta {
		rng = append(rng, i)
	}
	return rng
}
