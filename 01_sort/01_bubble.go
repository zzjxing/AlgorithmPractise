package mysort

// BubbleSort 对整数切片 nums 进行原地冒泡排序
func BubbleSort(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		swapped := false // 标记本轮是否有交换
		for j := 0; j < n-1-i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}
		if !swapped {
			// 如果本轮没有发生交换，说明已经有序，可以提前结束
			break
		}
	}
}
