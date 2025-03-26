package search

import "math/rand/v2"

// FindKLargest // 第 K 大元素（返回排序后倒数第 k 个）
func FindKLargest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

// FindKthSmallest 第 K 小元素（返回排序后正数第 k 个）
func FindKthSmallest(nums []int, k int) int {
	return quickSelect(nums, 0, len(nums)-1, k-1)
}

// quickSelect 寻找第 k 小元素（k 是索引）
func quickSelect(nums []int, start, end, k int) int {
	if start <= end {
		index := partition(nums, start, end)
		if index == k {
			return nums[index]
		} else if index > k {
			return quickSelect(nums, start, index-1, k)
		} else {
			return quickSelect(nums, index+1, end, k)
		}
	}
	return -1 // fallback
}

func partition(nums []int, start, end int) int {
	mid := rand.IntN(end-start+1) + start
	nums[start], nums[mid] = nums[mid], nums[start]
	pivot := nums[start]
	l, r := start, end
	for l < r {
		for l < r && nums[r] >= pivot {
			r--
		}
		nums[l] = nums[r]
		for l < r && nums[l] < pivot {
			l++
		}
		nums[r] = nums[l]
	}
	nums[l] = pivot
	return l
}
