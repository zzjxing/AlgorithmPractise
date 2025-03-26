package array

func TwoSum(nums []int, target int) []int {
	hash := make(map[int]int)
	for i, num := range nums {
		if idx, exist := hash[target-num]; exist {
			return []int{i, idx}
		}
		hash[num] = i
	}
	return nil
}
