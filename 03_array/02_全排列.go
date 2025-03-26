package array

func Permute(nums []int) [][]int {
	ret := make([][]int, 0)
	var backTrack func(start int)
	backTrack = func(start int) {
		if start == len(nums) {
			tmp := append([]int{}, nums...)
			ret = append(ret, tmp)
			return
		}
		for i := start; i < len(nums); i++ {
			nums[i], nums[start] = nums[start], nums[i]
			backTrack(start + 1)
			nums[i], nums[start] = nums[start], nums[i]
		}

	}
	backTrack(0)
	return ret
}
