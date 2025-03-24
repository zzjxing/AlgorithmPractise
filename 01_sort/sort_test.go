// 文件名: quick_sort_test.go
package mysort

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{5, 3, 8, 4, 2}, []int{2, 3, 4, 5, 8}},
		{[]int{9, -1, 3, 0, 2, 1}, []int{-1, 0, 1, 2, 3, 9}},
		{[]int{5, 5, 5, 5}, []int{5, 5, 5, 5}},
	}

	for _, tt := range tests {
		// 拷贝输入，避免修改原始切片
		inputCopy := make([]int, len(tt.input))

		// 快排测试
		copy(inputCopy, tt.input)
		QuickSort(inputCopy)
		if !reflect.DeepEqual(inputCopy, tt.expected) {
			t.Errorf("QuickSort(%v) = %v; want %v", tt.input, inputCopy, tt.expected)
		}
		// 冒泡测试
		copy(inputCopy, tt.input)
		BubbleSort(inputCopy)
		if !reflect.DeepEqual(inputCopy, tt.expected) {
			t.Errorf("Bubble(%v) = %v; want %v", tt.input, inputCopy, tt.expected)
		}
	}
}
