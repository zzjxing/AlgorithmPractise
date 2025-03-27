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

func TestTrap(t *testing.T) {
	tests := []struct {
		height   []int
		expected int
	}{
		{[]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}, 6}, // 标准用例
		{[]int{4, 2, 0, 3, 2, 5}, 9},                   // 多层积水
		{[]int{1, 2, 3, 4, 5}, 0},                      // 递增，无法积水
		{[]int{5, 4, 3, 2, 1}, 0},                      // 递减，无法积水
		{[]int{5}, 0},                                  // 只有一个柱子
		{[]int{1, 2, 1}, 0},                            // 只有两个柱子
		{[]int{2, 0, 2}, 2},                            // 形成一个小池子
		{[]int{}, 0},                                   // 空数组
	}

	for _, tt := range tests {
		got := Trap(tt.height)
		if got != tt.expected {
			t.Errorf("Trap(%v) = %d, want %d", tt.height, got, tt.expected)
		}
	}
}

// sortSlices 对子集进行排序，使结果一致
func sortSlices(slices [][]int) {
	for _, s := range slices {
		sort.Ints(s) // 先对每个子集内部排序
	}
	sort.Slice(slices, func(i, j int) bool {
		// 然后对整个结果排序
		if len(slices[i]) != len(slices[j]) {
			return len(slices[i]) < len(slices[j])
		}
		for k := range slices[i] {
			if slices[i][k] != slices[j][k] {
				return slices[i][k] < slices[j][k]
			}
		}
		return false
	})
}

func TestSubSets(t *testing.T) {
	tests := []struct {
		nums     []int
		expected [][]int
	}{
		// 标准用例
		{[]int{1, 2, 3}, [][]int{
			{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3},
		}},
		// 只有一个元素
		{[]int{1}, [][]int{
			{}, {1},
		}},
		// 空数组
		{[]int{}, [][]int{
			{},
		}},
		// 包含重复数字（通常假设唯一，这里测试行为）
		{[]int{1, 2, 2}, [][]int{
			{}, {1}, {2}, {2}, {1, 2}, {1, 2}, {2, 2}, {1, 2, 2},
		}},
	}

	for _, tt := range tests {
		got := SubSets(tt.nums)
		sortSlices(got)         // 排序实际输出
		sortSlices(tt.expected) // 排序预期输出
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("SubSets(%v) = %v, want %v", tt.nums, got, tt.expected)
		}
	}
}

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}}, // 最大排列，回归最小
		{[]int{1, 1, 5}, []int{1, 5, 1}},
		{[]int{1, 3, 2}, []int{2, 1, 3}},
		{[]int{2, 3, 1}, []int{3, 1, 2}},
		{[]int{5, 4, 7, 5, 3, 2}, []int{5, 5, 2, 3, 4, 7}},
	}

	for _, tt := range tests {
		numsCopy := append([]int{}, tt.nums...) // 复制数组，防止修改原测试数据
		NextPermutation(numsCopy)
		if !reflect.DeepEqual(numsCopy, tt.expected) {
			t.Errorf("NextPermutation(%v) = %v, want %v", tt.nums, numsCopy, tt.expected)
		}
	}
}

func TestMaxSlidingWindow(t *testing.T) {
	tests := []struct {
		nums     []int
		k        int
		expected []int
	}{
		// 标准用例
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 3, []int{3, 4, 5, 6, 7, 8, 9}},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, 3, []int{9, 8, 7, 6, 5, 4, 3}},
		{[]int{1, 3, 1, 2, 0, 5}, 3, []int{3, 3, 2, 5}},
		// 窗口大小等于数组长度
		{[]int{1, 2, 3, 4}, 4, []int{4}},
		// k=1，每个元素都是最大值
		{[]int{1, 2, 3, 4, 5}, 1, []int{1, 2, 3, 4, 5}},
		// 只有一个元素
		{[]int{5}, 1, []int{5}},
		// 空数组
		{[]int{}, 3, nil},
		// k 大于数组长度（无效输入）
		{[]int{1, 2}, 3, nil},
	}

	for _, tt := range tests {
		got := MaxSlidingWindow(tt.nums, tt.k)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("MaxSlidingWindow(%v, %d) = %v, want %v", tt.nums, tt.k, got, tt.expected)
		}
	}
}

func TestCoinChange(t *testing.T) {
	tests := []struct {
		coins    []int
		target   int
		expected int
	}{
		// 标准用例
		{[]int{1, 2, 5}, 11, 3},              // 5+5+1=11
		{[]int{2}, 3, -1},                    // 目标 3，只有 2，无解
		{[]int{1}, 0, 0},                     // 目标 0，直接返回 0
		{[]int{1}, 2, 2},                     // 目标 2，只有 1，2 个 1
		{[]int{186, 419, 83, 408}, 6249, 20}, // 复杂大数测试
		// 目标不可达情况
		{[]int{3, 7}, 5, -1}, // 只能选 3 或 7，无法拼出 5
		// 目标刚好可达
		{[]int{2, 5, 10, 1}, 27, 4}, // 10+10+5+2=27
	}

	for _, tt := range tests {
		got := CoinChange(tt.coins, tt.target)
		if got != tt.expected {
			t.Errorf("CoinChange(%v, %d) = %d, want %d", tt.coins, tt.target, got, tt.expected)
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
