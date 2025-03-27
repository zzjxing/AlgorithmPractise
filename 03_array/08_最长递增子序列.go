package array

// LengthOfLIS 数组的最长递增子序列
// 思路：dp，dp[i]标识以 nums[i]为结尾的最长递增子序列
// 对于每个nums[i]，从 0 开始向 i 遍历，去寻找所有比nums[i]小的nums[j]
// nums[j]<nums[i]，表明nums[j]可以被加入到以 nums[i]为结尾的递增子序列中
// 更新dp[i]，更新 ret
func LengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}
	var ret int = 1
	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ret = max(ret, dp[i])
	}
	return ret
}
