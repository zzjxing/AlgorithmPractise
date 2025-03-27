package graph

// MaxSquare matrix 中的最大正方形的面积
// 思路：dp，dp[i][j]表示以matrix[i][j]为右下角的最大正方形的边长
// 状态转移方程：matrix[i][j]为 1 的情况下，dp[i][j]为其上侧、左侧、左上侧的最小正方形值+1
func MaxSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	var ret = 0
	// 初始化 上侧和左侧的 dp 数组
	for i := 0; i < m; i++ {
		dp[i][0] = int(matrix[i][0] - '0')
		if dp[i][0] == 1 {
			ret = 1
		}
	}
	for j := 0; j < n; j++ {
		dp[0][j] = int(matrix[0][j] - '0')
		if dp[0][j] == 1 {
			ret = 1
		}
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == '1' {
				dp[i][j] = min(dp[i-1][j-1], min(dp[i-1][j], dp[i][j-1])) + 1
			}
			ret = max(ret, dp[i][j])
		}
	}
	return ret * ret
}
