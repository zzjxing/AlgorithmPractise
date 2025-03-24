package search

// BinarySearch :给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target
// 写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。
// https://leetcode.cn/problems/binary-search
func BinarySearch(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
