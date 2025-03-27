package graph

import (
	"reflect"
	"testing"
)

func TestNumIslands(t *testing.T) {
	tests := []struct {
		grid     [][]byte
		expected int
	}{
		{
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 1,
		},
		{
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expected: 3,
		},
		{
			grid: [][]byte{
				{'1'},
			},
			expected: 1,
		},
		{
			grid: [][]byte{
				{'0'},
			},
			expected: 0,
		},
		{
			grid:     [][]byte{},
			expected: 0,
		},
	}

	for i, tt := range tests {
		// 深拷贝 grid，防止修改影响后续测试
		gridCopy := make([][]byte, len(tt.grid))
		for i := range tt.grid {
			gridCopy[i] = make([]byte, len(tt.grid[i]))
			copy(gridCopy[i], tt.grid[i])
		}

		got := NumIslands(gridCopy)
		if got != tt.expected {
			t.Errorf("Test case %d: got %d, want %d", i+1, got, tt.expected)
		}
	}
}

func TestRotateImage(t *testing.T) {
	tests := []struct {
		matrix   [][]int
		expected [][]int
	}{
		// 3x3 矩阵
		{
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		// 4x4 矩阵
		{
			matrix: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			expected: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
		// 1x1 矩阵
		{
			matrix:   [][]int{{1}},
			expected: [][]int{{1}},
		},
		// 2x2 矩阵
		{
			matrix: [][]int{
				{1, 2},
				{3, 4},
			},
			expected: [][]int{
				{3, 1},
				{4, 2},
			},
		},
	}

	for _, tt := range tests {
		matrixCopy := make([][]int, len(tt.matrix))
		for i := range tt.matrix {
			matrixCopy[i] = append([]int{}, tt.matrix[i]...)
		}

		got := RotateImage(matrixCopy)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("RotateImage(%v) = %v, want %v", tt.matrix, got, tt.expected)
		}
	}
}
