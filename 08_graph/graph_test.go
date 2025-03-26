package graph

import "testing"

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
