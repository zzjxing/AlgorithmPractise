package array

import "sort"

// MergeInterval 合并相交的区间
// 例如[1,3][2,4]合并为[1,4], [1,5][2,3]合并为[1,5]
// 思路：先按照区间头进行排序，然后循环合并
func MergeInterval(intervals [][]int) [][]int {
	var ret [][]int
	if len(intervals) == 0 {
		return ret
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	curr := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if curr[1] >= intervals[i][0] {
			curr[1] = max(curr[1], intervals[i][1])
		} else {
			ret = append(ret, curr)
			curr = intervals[i]
		}
	}
	ret = append(ret, curr)
	return ret
}
