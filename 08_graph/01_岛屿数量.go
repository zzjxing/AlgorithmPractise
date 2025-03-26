package graph

func numIslands(grid [][]byte) int {
	ret := 0
	directions := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	var bfs func(i, j int)
	bfs = func(i, j int) {
		if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		for _, dir := range directions {
			bfs(i+dir[0], j+dir[1])
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == '1' {
				ret++
				bfs(i, j)
			}
		}
	}
	return ret
}
