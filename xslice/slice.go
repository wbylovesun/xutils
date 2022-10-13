package xslice

// Contains 判断数组和切片中是否存在某个值
func Contains[T SliceElementType](data []T, ct T) bool {
	if data == nil {
		return false
	}
	for _, v := range data {
		if v == ct {
			return true
		}
	}
	return false
}
