package search

import (
	"reflect"
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{1, 2, 3}, 0, -1},
	}
	for _, tt := range tests {
		got := BinarySearch(tt.nums, tt.target)
		if got != tt.want {
			t.Errorf("BinarySearch(%v, %d) = %d; want %d", tt.nums, tt.target, got, tt.want)
		}
	}
}

func TestSearchInRotatedArray(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
	}
	for _, tt := range tests {
		got := SearchInRotatedArray(tt.nums, tt.target)
		if got != tt.want {
			t.Errorf("SearchInRotatedArray(%v, %d) = %d; want %d", tt.nums, tt.target, got, tt.want)
		}
	}
}

func TestFindMedianSortArrays(t *testing.T) {
	tests := []struct {
		nums1, nums2 []int
		want         float64
	}{
		{[]int{1, 3}, []int{2}, 2.0},
		{[]int{1, 2}, []int{3, 4}, 2.5},
	}
	for _, tt := range tests {
		got := FindMedianSortArrays(tt.nums1, tt.nums2)
		if got != tt.want {
			t.Errorf("FindMedianSortArrays(%v, %v) = %f; want %f", tt.nums1, tt.nums2, got, tt.want)
		}
	}
}

func TestSearchRange(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{5, 7, 7, 8, 8, 10}, 8, []int{3, 4}},
		{[]int{5, 7, 7, 8, 8, 10}, 6, []int{-1, -1}},
	}
	for _, tt := range tests {
		got := SearchRange(tt.nums, tt.target)
		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("SearchRange(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}

func TestFindPeakElement(t *testing.T) {
	tests := []struct {
		nums []int
		want int // 只要返回的是峰值的下标即可，不唯一
	}{
		{[]int{1, 2, 3, 1}, 2},
		{[]int{1, 2, 1, 3, 5, 6, 4}, 5}, // 5 或 1 也都是合法峰值下标
	}
	for _, tt := range tests {
		got := FindPeakElement(tt.nums)
		// 只验证是否为峰值即可
		if !isPeak(tt.nums, got) {
			t.Errorf("FindPeakElement(%v) = %d is not peak", tt.nums, got)
		}
	}
}

// 判断某个下标是否是峰值
func isPeak(nums []int, idx int) bool {
	n := len(nums)
	left := idx == 0 || nums[idx] > nums[idx-1]
	right := idx == n-1 || nums[idx] > nums[idx+1]
	return left && right
}
