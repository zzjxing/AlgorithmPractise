package search

// SearchRange 查找某元素出现的第一个位置和最后一个位置
// 思想：查找target 出现的左边界和右边界即可
// 使用二分查找，在二分查找找到的 target 的时候，如果寻找左边界，就继续在左侧寻找，否则在右侧寻找
// 因此关键在于 nums[mid]==target 时，如何处理。
func SearchRange(nums []int, target int) []int {
	return []int{binarySearch(nums, target, false), binarySearch(nums, target, true)}
}

// 二分查找变种，如果 flag 为 true，则寻找右边界，否则寻找左边界
func binarySearch(nums []int, target int, flag bool) int {
	left, right := 0, len(nums)-1
	var ret = -1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			ret = mid
			if flag { //寻找右边界，在右侧继续寻找
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ret
}
