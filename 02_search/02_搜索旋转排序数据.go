package search

// SearchInRotatedArray
// 题目描述
//
//	整数数组 nums 按升序排列，数组中的值 互不相同 。
//	在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了旋转，
//	使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。
//	例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
//	给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回 -1 。
//	https://leetcode.cn/problems/search-in-rotated-sorted-array
//
// 题解 二分查找 + 判断有序区域
//
//	每次 mid 将数组分为两半，一半是有序的（因为整体是升序+旋转）
//	判断 target 是否在有序半边中，来缩小搜索范围
//	时间复杂度：O(log n)
//	空间复杂度：O(1)
func SearchInRotatedArray(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] { // 左边有序
			// eg.5678912 找 1
			// mid=8 5<8 如果 target 在 5-8 之间，则在左半部分找，否则在右半部分找
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // 右边有序
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
