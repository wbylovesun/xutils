package xslice

type NumberSliceElementType interface {
	int | int8 | int16 | int32 | int64
}

type UnsignedNumberSliceElementType interface {
	uint | uint8 | uint16 | uint32 | uint64
}

type SliceElementType interface {
	int | int8 | int16 | int32 | int64 |
	uint | uint8 | uint16 | uint32 | uint64 |
	float32 | float64 | string
}
