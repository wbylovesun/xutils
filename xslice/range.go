package xslice

// Range 生成 [from, to] 范围内以 delta 为步长的整数切片（包含两端）；from > to 时返回 nil。
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
