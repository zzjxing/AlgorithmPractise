package array

// MaxSlidingWindow 滑动窗口最大值
// 思路：维护一个单调队列，存储窗口内可以成为最大值的索引。保持队列单调递减，确保队首元素是最大值
// 关键：1. 维护队首索引，如果队首索引超过窗口范围，则溢出
//  2. 保持队列单调递减，在将 i append进入队列之前，先移除队尾所有小于nums[i]的值
func MaxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 || len(nums) < k {
		return nil
	}
	var ret []int
	var deque []int
	for i := 0; i < len(nums); i++ {
		// 弹出队首 1 [3  -1  -3] 5  3  6  7
		// 1 2 3  4 k=3 i=4 deque[0]=1
		if len(deque) > 0 && i-k >= deque[0] {
			deque = deque[1:]
		}
		// 弹出 deque 中比 nums[i] 小的数
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		ret = append(ret, nums[deque[0]])
	}
	return ret[k-1:] // 前 k-1个数字组不成长度为 k 的窗口
}
