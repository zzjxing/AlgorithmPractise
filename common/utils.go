package common

// Abs 计算任意整数或浮点数的绝对值
func Abs[T int | int8 | int16 | int32 | int64 | float32 | float64](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
