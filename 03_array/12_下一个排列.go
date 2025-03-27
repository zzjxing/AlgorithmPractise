package array

// NextPermutation 返回比 nums 字典序大的下一个排列，eg. 2 5 4 3 2
// 从右侧开始寻找，找到第一个nums[i]<nums[i+1]的位置：0
// 在i右侧找到比nums[i]大的最小数的位置j：3
// 交换nums[i] 和 nums[j]: 2 5 4 3 2 -> 3 5 4 2 2
// 反转 i 后面的数组: 3 5 4 2 2 -> 3 2 2 4 5
func NextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}
	i := len(nums) - 2
	for ; i >= 0; i-- {
		if nums[i] < nums[i+1] {
			break
		}
	}
	if i >= 0 {
		j := len(nums) - 1
		for j > i && nums[j] <= nums[i] {
			j--
		}
		nums[j], nums[i] = nums[i], nums[j]
	}
	reverseSlice(nums, i+1, len(nums)-1)
}

func reverseSlice(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
