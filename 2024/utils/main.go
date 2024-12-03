package utils

type numeric interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~float32 | ~float64
}

func Abs[T numeric](val T) T {
	if val < 0 {
		return -val
	}
	return val
}
