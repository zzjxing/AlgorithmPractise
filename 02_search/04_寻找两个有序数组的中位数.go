package search

// FindMedianSortArrays
// 解法一：双指针，每个都指向两个数组的头节点，较小值向右移动，直到移动到中心
// 解法二：重新定义函数 findK(nums1,nums2,k)，寻找 nums1 nums2中的第 K 小数，
//
//	findK：二分查找实现
//		1. 取两个数组的第k/2个元素
//		2. 如果nums1[k/2]<nums2[k/2]: nums1和 nums2的第 k 大数一定不在 nums1的前 k/2 个数中
//		3. 丢弃 nums1的钱k/2个数，重新调整 k 值，进行循环，直到 nums1 或者 nums2长度为 0，或者 k 值为 1
func FindMedianSortArrays(nums1, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2) // 3 4
	mid := (l1 + l2) / 2             // 3
	if (l1+l2)%2 == 1 {
		return float64(findK(nums1, nums2, mid+1))
	}
	return float64(findK(nums1, nums2, mid)+findK(nums1, nums2, mid+1)) / 2.0
}

// findK 返回数组 nums1 和 nums2 的第 k 小数
func findK(nums1, nums2 []int, k int) int {
	for {
		if len(nums1) == 0 {
			return nums2[k-1]
		}
		if len(nums2) == 0 {
			return nums1[k-1]
		}
		if k == 1 {
			return min(nums1[0], nums2[0])
		}
		m1 := min(len(nums1), k/2) - 1
		m2 := min(len(nums2), k/2) - 1
		if nums1[m1] < nums2[m2] {
			nums1 = nums1[m1+1:]
			k -= m1 + 1
		} else {
			nums2 = nums2[m2+1:]
			k -= m2 + 1
		}
	}
}
