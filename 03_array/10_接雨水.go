package array

// Trap 接雨水
// 每个柱子能接的雨水为：两侧比这个柱子大的所有值中最小的一个。
// 维护一个 leftMax 和rightMax，用于记录每个柱子左侧和右侧柱子的最大值
// 则每个柱子能接的雨水为：min(leftMax[i],rightMax[i])-height[i]
func Trap(height []int) int {
	n := len(height)
	if n <= 1 {
		return 0
	}
	ret := 0
	leftMax, rightMax := make([]int, n), make([]int, n)
	leftMax[0] = height[0]
	rightMax[n-1] = height[n-1]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	for i := 0; i < n; i++ {
		ret += min(leftMax[i], rightMax[i]) - height[i]
	}
	return ret
}
