package array

// SearMatrix 搜索二维矩阵
// 矩阵：从左到右升序，从上到下升序，在该矩阵中寻找 target 是否存在
// 思路：从右上角看，该矩阵可以抽象为一颗搜索树，因此可以从右上角寻找
func SearMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	i, j := 0, len(matrix[0])-1
	for i < len(matrix) && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] > target {
			j--
		} else {
			i++
		}
	}
	return false
}
