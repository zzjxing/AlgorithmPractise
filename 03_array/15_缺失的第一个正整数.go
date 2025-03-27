package array

import (
	"AlgorithmPractise/common"
	"math"
)

// FirstMissingPositive 缺失的第一个正数
// 原地哈希，置换法，在 i 索引存储 i+1 的值，哈希函数：f(nums[i])=nums[i]-1
func FirstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); i++ {
		// 置换 nums[i] 到正确的位置，直到nums[i]中存储了i+1
		for nums[i] > 0 && nums[i] <= len(nums) && nums[i] != nums[nums[i]-1] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}

// FirstMissingPositive2
// 方法2：预处理非正数，置为MaxInt，然后遍历nums，将nums[i]-1的位置置为负数，标记nums[i]出现过
func FirstMissingPositive2(nums []int) int {
	// 预处理，将数组中所有数都置为正整数
	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 {
			nums[i] = math.MaxInt
		}
	}
	// 如果1-n出现过，则将其对应索引置为负数
	for i := 0; i < len(nums); i++ {
		absX := common.Abs(nums[i]) // 2.由于下面的标记，nums[i]可能为负数
		if absX <= len(nums) {
			nums[absX-1] = -common.Abs(nums[absX-1]) //1. 标记nums[i]出现过
		}
	}
	for i := range nums {
		if nums[i] >= 0 { // 索引 i 的值未被置为负数，则 i+1 未出现过
			return i + 1
		}
	}
	return len(nums) + 1
}
