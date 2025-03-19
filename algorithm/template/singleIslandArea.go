package template

import "fmt"

func sumSingleIsland(grid [][]int) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	for i := 0; i < len(grid); i++ {
		visitIsland2(grid, i, 0)
		visitIsland2(grid, i, len(grid[0])-1)
	}
	for i := 0; i < len(grid[0]); i++ {
		visitIsland2(grid, 0, i)
		visitIsland2(grid, len(grid)-1, i)
	}
	sum := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				sum += grid[i][j]
			}

		}
	}
	return sum
}

func visitIsland2(grid [][]int, x int, y int) {
	if grid[x][y] != 1 {
		return
	}
	grid[x][y] = 2
	diff := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, value := range diff {
		nextX := x + value[0]
		nextY := y + value[1]
		if nextX < 0 || nextX >= len(grid) || nextY < 0 || nextY >= len(grid[0]) {
			continue
		}
		if grid[nextX][nextY] == 1 {
			visitIsland2(grid, nextX, nextY)
		}
	}
}
func TestSumSingleIsland() {
	var n, m int
	fmt.Scan(&n, &m)

	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&grid[i][j])
		}
	}
	sumSingleIsland(grid)
}
