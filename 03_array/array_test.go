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
