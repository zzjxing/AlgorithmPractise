package graph

// RotateImage 旋转图像
// 思路，1. 转秩 2. 反转
// 1 2 3     1 4 7     7 4 1
// 4 5 6 ->  2 5 8 ->  8 5 2
// 7 8 9     3 6 9     9 6 3
func RotateImage(matrix [][]int) [][]int {
	n := len(matrix)
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
	return matrix
}
