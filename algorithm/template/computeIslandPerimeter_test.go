package template

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func computeIslandPerimeter(grid [][]int) int {
	islandPerimeter := 0
	if grid == nil || len(grid) == 0 {
		return islandPerimeter
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				islandPerimeter += computePerimeter(grid, i, j)
			}
		}
	}
	return islandPerimeter
}

func computePerimeter(grid [][]int, x, y int) int {
	diffArr := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	perimeter := 0
	for i := 0; i < 4; i++ {
		nextX := x + diffArr[i][0]
		nextY := y + diffArr[i][1]
		if nextX < 0 || nextY < 0 || nextX >= len(grid) || nextY >= len(grid[0]) {
			perimeter++
			continue
		}
		if grid[nextX][nextY] == 0 {
			perimeter++
		}
	}
	return perimeter
}

func TestComputeIslandPerimeter(t *testing.T) {
	// var grid = [][]int{
	// 	{0, 1, 0, 0},
	// 	{1, 1, 1, 0},
	// 	{0, 1, 0, 0},
	// }
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()

	n, m := parseInput(line)

	// 初始化 grid
	grid := make([][]int, n)
	for i := range grid {
		grid[i] = make([]int, m)
	}
	// 读入 grid 数据
	for i := 0; i < n; i++ {
		scanner.Scan()
		line := scanner.Text()
		values := parseLine(line, m)
		for j := 0; j < m; j++ {
			grid[i][j] = values[j]
		}
	}
	fmt.Println(computeIslandPerimeter(grid)) // 10
}

func parseInput(line string) (int, int) {
	parts := strings.Split(line, " ")
	n, _ := strconv.Atoi(parts[0])
	m, _ := strconv.Atoi(parts[1])
	return n, m
}

// parseLine 解析一行中的多个值
func parseLine(line string, count int) []int {
	parts := strings.Split(line, " ")
	values := make([]int, count)
	for i := 0; i < count; i++ {
		values[i], _ = strconv.Atoi(parts[i])
	}
	return values
}
