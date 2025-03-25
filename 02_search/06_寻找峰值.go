package search

/*
峰值元素是指其值严格大于左右相邻值的元素。
给你一个整数数组 nums，找到峰值元素并返回其索引。数组可能包含多个峰值，在这种情况下，返回 任何一个峰值 所在位置即可。
你可以假设 nums[-1] = nums[n] = -∞ 。
你必须实现时间复杂度为 O(log n) 的算法来解决此问题。
*/
// 题目假设 nums[-1] = nums[n] = -∞, 且相邻元素不相同，因此峰值一定存在，而且一定是相邻元素中更大的那个

// FindPeakElement 利用二分查找：mid 指针一直向更大的方向移动，一定可以找到峰值
//   - 如果 nums[mid] < nums[mid+1]，说明峰值在右边
//   - 如果 nums[mid] > nums[mid+1]，说明峰值在左边（mid 可能是峰）
func FindPeakElement(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid // mid有可能是峰值，因此不能舍弃
		}
	}
	return left
}
