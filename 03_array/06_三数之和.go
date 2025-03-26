package array

import "sort"

// ThreeSum 返回所有三数之和为 0 的组合，组合不允许重复
// 排序数组：方便去重
//  1. 遍历数组，当前数字为 a
//  2. 在 a 的右侧使用双指针查找a+b+c==0的组合
//     a+b+c==0: 记录三数之和，双指针向中间移动(注意去重)
//     a+b+c>0: 缩小三数之和，右指针左移
//     a+b+c<0: 增大三数之和，左指针右移
func ThreeSum(nums []int) [][]int {
	var ret [][]int
	if len(nums) < 3 {
		return ret
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		a := nums[i]
		left, right := i+1, len(nums)-1
		for left < right {
			b, c := nums[left], nums[right]
			if a+b+c == 0 {
				ret = append(ret, []int{a, b, c})
				for left < right && nums[left] == nums[left+1] { // 去重
					left++
				}
				for left < right && nums[right] == nums[right-1] { // 去重
					right--
				}
				left++
				right--
			} else if a+b+c > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return ret
}
