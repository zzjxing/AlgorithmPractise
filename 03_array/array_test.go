package array

import (
	"reflect"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	res := TwoSum(nums, target)
	if nums[res[0]]+nums[res[1]] != target {
		t.Errorf("TwoSum failed, got %v", res)
	}
}

func TestPermute(t *testing.T) {
	nums := []int{1, 2, 3}
	res := Permute(nums)
	if len(res) != 6 {
		t.Errorf("Permute length mismatch: got %d", len(res))
	}
}

func TestMergeTwoSlice(t *testing.T) {
	a := []int{1, 3, 5}
	b := []int{2, 4, 6}
	want := []int{1, 2, 3, 4, 5, 6}
	got := MergeTwoSlice(a, b)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MergeTwoSlice failed: got %v, want %v", got, want)
	}
}

func TestMaxProfitOnce(t *testing.T) {
	input := []int{7, 1, 5, 3, 6, 4}
	want := 5
	got := MaxProfitOnce(input)
	if got != want {
		t.Errorf("MaxProfitOnce failed: got %d, want %d", got, want)
	}
}

func TestMaxProfitNoLimit(t *testing.T) {
	input := []int{7, 1, 5, 3, 6, 4}
	want := 7
	got := MaxProfitNoLimit(input)
	if got != want {
		t.Errorf("MaxProfitNoLimit failed: got %d, want %d", got, want)
	}
}

func TestMaxSumArray(t *testing.T) {
	input := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	want := 6
	got := MaxSumArray(input)
	if got != want {
		t.Errorf("MaxSumArray failed: got %d, want %d", got, want)
	}
}

func TestThreeSum(t *testing.T) {
	input := []int{-1, 0, 1, 2, -1, -4}
	got := ThreeSum(input)
	expected := [][]int{{-1, -1, 2}, {-1, 0, 1}}

	// 排序结果统一顺序便于判断
	sortTriplets := func(tri [][]int) {
		for _, t := range tri {
			sort.Ints(t)
		}
		sort.Slice(tri, func(i, j int) bool {
			for x := range tri[i] {
				if tri[i][x] != tri[j][x] {
					return tri[i][x] < tri[j][x]
				}
			}
			return false
		})
	}
	sortTriplets(got)
	sortTriplets(expected)

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("ThreeSum failed: got %v, want %v", got, expected)
	}
}

func TestSpiralOrder(t *testing.T) {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	want := []int{1, 2, 3, 6, 9, 8, 7, 4, 5}
	got := SpiralOrder(matrix)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("SpiralOrder failed: got %v, want %v", got, want)
	}
}

func TestLengthOfLIS(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4}, // 最长递增子序列: [2,3,7,101]
		{[]int{0, 1, 0, 3, 2, 3}, 4},           // 最长递增子序列: [0,1,2,3]
		{[]int{7, 7, 7, 7, 7, 7, 7}, 1},        // 所有元素相同，只能取 1
		{[]int{1, 3, 5, 4, 7}, 4},              // 最长递增子序列: [1,3,5,7]
		{[]int{1}, 1},                          // 只有一个元素
		{[]int{}, 0},                           // 空数组
	}

	for _, tt := range tests {
		got := LengthOfLIS(tt.nums)
		if got != tt.expected {
			t.Errorf("LengthOfLIS(%v) = %d, want %d", tt.nums, got, tt.expected)
		}
	}
}

func TestMergeInterval(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected [][]int
	}{
		// 标准用例
		{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
		// 完全重叠
		{[][]int{{1, 4}, {2, 3}}, [][]int{{1, 4}}},
		// 没有重叠
		{[][]int{{1, 2}, {3, 4}, {5, 6}}, [][]int{{1, 2}, {3, 4}, {5, 6}}},
		// 只有一个区间
		{[][]int{{1, 5}}, [][]int{{1, 5}}},
		// 空输入
		{[][]int{}, nil},
	}

	for _, tt := range tests {
		got := MergeInterval(tt.input)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("MergeInterval(%v) = %v, want %v", tt.input, got, tt.expected)
		}
	}
}

func TestLongestConsecutive(t *testing.T) {
	tests := []struct {
		nums     []int
		expected int
	}{
		{[]int{100, 4, 200, 1, 3, 2}, 4},              // 最长连续序列: [1,2,3,4]
		{[]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}, 9},      // 最长连续序列: [0,1,2,3,4,5,6,7,8]
		{[]int{1, 2, 0, 1}, 3},                        // [0,1,2]
		{[]int{1}, 1},                                 // 只有一个元素
		{[]int{}, 0},                                  // 空数组
		{[]int{9, 1, 4, 7, 3, -1, 0, 5, 8, -1, 6}, 7}, // [-1,0,1,3,4,5,6]
	}

	for _, tt := range tests {
		got := LongestConsecutive(tt.nums)
		if got != tt.expected {
			t.Errorf("LongestConsecutive(%v) = %d, want %d", tt.nums, got, tt.expected)
		}
	}
}
