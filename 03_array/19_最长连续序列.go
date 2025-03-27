package array

// LongestConsecutive 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
// nums = [100,4,200,1,3,2]
// 输出：4： 1-2-3-4
// 思路：以时间换空间,map[int]bool表明某个值是否出现过，然后对于每个num，递增去查 num+1是否出现过
func LongestConsecutive(nums []int) int {
	hash := make(map[int]bool)
	for i := range nums {
		hash[nums[i]] = true
	}
	var ret = 0
	for num := range hash {
		if hash[num-1] {
			continue
		}
		currLen := 0
		currNum := num
		for hash[currNum] {
			currLen++
			currNum++
		}
		ret = max(ret, currLen)
	}
	return ret
}
