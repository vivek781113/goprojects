package leetcode

func maxAreaOfIsland(grid [][]int) int {
	row := len(grid) - 1
	col := len(grid[0]) - 1
	maxArea := 0
	for r := 0; r <= row; r++ {
		for c := 0; c <= col; c++ {
			curr := grid[r][c]
			if curr == 1 {
				area := dfs(grid, r, c, row, col, 0)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}
	return maxArea
}

func dfs(grid [][]int, x, y, row, col, areaCount int) int {
	if x > row || x < 0 || y > col || y < 0 || grid[x][y] == 0 {
		return areaCount
	}
	grid[x][y] = 0
	areaCount++
	dx := []int{0, 0, 1, -1}
	dy := []int{-1, 1, 0, 0}
	for i := 0; i < 4; i++ {
		areaCount = dfs(grid, x+dx[i], y+dy[i], row, col, areaCount)
	}

	return areaCount
}
