package array

func MergeTwoSlice(nums1, nums2 []int) []int {
	ret := make([]int, 0)
	i, j := 0, 0
	for i < len(nums1) || j < len(nums2) {
		if i >= len(nums1) {
			ret = append(ret, nums2[j])
			j++
			continue
		}
		if j >= len(nums2) {
			ret = append(ret, nums1[i])
			i++
			continue
		}
		if nums1[i] < nums2[j] {
			ret = append(ret, nums1[i])
			i++
		} else {
			ret = append(ret, nums2[j])
			j++
		}
	}
	return ret
}
