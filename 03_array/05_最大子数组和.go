package array

// MaxSumArray 给你一个整数数组 nums, 请你找出一个具有最大和的连续子数组(子数组最少包含一个元素),返回其最大和.
//  1. 动态规划，遍历数组，计算以 nums[i] 结尾的最大子数组和：
//     ○ 如果当前和 curSum + nums[i] 变小了，说明之前的子数组是负贡献，丢弃，重新从 nums[i] 开始计算。
//     ○ 如果 curSum + nums[i] 变大了，则继续扩展子数组。
//     ○ 状态转移方程：currSum=max(nums[i],nums[i]+currSum)
//  2. 维护一个全局最大值 maxSum，记录所有 curSum 中的最大值： maxSum=max(maxSum,currSum)
func MaxSumArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	currSum, maxSum := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		currSum = max(nums[i], currSum+nums[i])
		maxSum = max(maxSum, currSum)
	}
	return maxSum
}
