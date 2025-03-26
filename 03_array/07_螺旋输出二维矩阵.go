package array

// SpiralOrder 螺旋输出矩阵，每次外围的循环输出一个圈
func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	var ret []int
	top, bottom := 0, len(matrix)-1
	left, right := 0, len(matrix[0])-1
	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			ret = append(ret, matrix[top][i])
		}
		top++ // 可能会导致 top<=bottom 不成立
		for i := top; i <= bottom; i++ {
			ret = append(ret, matrix[i][right])
		}
		right-- // 可能会导致 left<=right 不成立
		if top <= bottom {
			for i := right; i >= left; i-- {
				ret = append(ret, matrix[bottom][i])
			}
			bottom--
		}
		if left <= right {
			for i := bottom; i >= top; i-- {
				ret = append(ret, matrix[i][left])
			}
			left++
		}
	}
	return ret
}
