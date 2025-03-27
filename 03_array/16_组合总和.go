package array

func CombinationSum(nums []int, target int) [][]int {
	var ret [][]int
	var backTrack func(path []int, target int, start int)
	backTrack = func(path []int, target int, start int) {
		if target == 0 {
			tmp := append([]int{}, path...)
			ret = append(ret, tmp)
			return
		}
		if target < 0 {
			return
		}
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backTrack(path, target-nums[i], i)
			path = path[:len(path)-1]
		}
	}
	backTrack([]int{}, target, 0)
	return ret
}
