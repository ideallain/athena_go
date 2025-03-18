package template

import (
	"fmt"
	"testing"
)

func visitIsland(grid [][]int) int {
	row := len(grid)
	if row == 0 {
		return 0
	}
	visited := make([][]bool, row)
	//go的这种初始化方式真的丑陋
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, len(grid[0]))
	}
	ans := 0
	for i := 0; i < row; i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				visited[i][j] = true
				ans++
				visitGrid(grid, visited, i, j)
			}
		}
	}
	return ans
}

func visitGrid(grid [][]int, visited [][]bool, x int, y int) {
	diff := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	for _, arr := range diff {
		nextX := x + arr[0]
		nextY := y + arr[1]
		if nextX < 0 || nextX >= len(grid) || nextY < 0 || nextY >= len(grid[0]) {
			continue
		}
		if !visited[nextX][nextY] && grid[nextX][nextY] == 1 {
			visited[nextX][nextY] = true
			visitGrid(grid, visited, nextX, nextY)
		}
	}
}

func TestVisitIsland(t *testing.T) {
	var row, col int
	fmt.Scan(&row, &col)
	if row <= 0 || col <= 0 {
		return
	}
	grid := make([][]int, row)
	for i := 0; i < row; i++ {
		grid[i] = make([]int, col)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			fmt.Scan(&grid[i][j])
		}
	}
	//这里必须要打印，不然报错会显示潜在的数组越界
	fmt.Println(visitIsland(grid))
}
