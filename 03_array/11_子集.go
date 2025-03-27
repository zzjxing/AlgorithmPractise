package array

func SubSets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	var ret [][]int
	var backTrack func(start int, path []int)
	backTrack = func(start int, path []int) {
		tmp := append([]int{}, path...)
		ret = append(ret, tmp)
		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backTrack(i+1, path)
			path = path[:len(path)-1]
		}
	}
	backTrack(0, []int{})
	return ret
}
