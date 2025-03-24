package mysort

import (
	"math/rand/v2"
)

func QuickSort(nums []int) {
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, start, end int) {
	if start < end {
		index := partition(nums, start, end)
		quickSort(nums, start, index-1)
		quickSort(nums, index+1, end)
	}
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
