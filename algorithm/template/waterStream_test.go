package template

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
)

func VisitWaterGrid(grid [][]int) [][]int {
	result := make([][]int, 0)

	if grid == nil || len(grid) == 0 {
		return result
	}
	used1 := make([][]bool, len(grid))
	used2 := make([][]bool, len(grid))
	for index, _ := range used1 {
		used1[index] = make([]bool, len(grid[0]))
		used2[index] = make([]bool, len(grid[0]))
	}
	for i := 0; i < len(grid); i++ {
		//在测试页面编译时没有math.MinInt这个变量，只能初始化为矩阵的值
		visit3(grid, used1, i, 0, grid[i][0])
		visit3(grid, used2, i, len(grid[0])-1, grid[i][len(grid[0])-1])
	}
	for j := 0; j < len(grid[0]); j++ {
		visit3(grid, used1, 0, j, grid[0][j])
		visit3(grid, used2, len(grid)-1, j, grid[len(grid)-1][j])
	}
	for i := 0; i < len(used1); i++ {
		for j := 0; j < len(used1[0]); j++ {
			if used1[i][j] && used2[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}
	return result
}

func visit3(grid [][]int, visited [][]bool, x int, y int, preValue int) {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) { // 修复边界条件
		return
	}
	if grid[x][y] < preValue || visited[x][y] {
		return
	}
	visited[x][y] = true
	visit3(grid, visited, x+1, y, grid[x][y])
	visit3(grid, visited, x-1, y, grid[x][y])
	visit3(grid, visited, x, y+1, grid[x][y])
	visit3(grid, visited, x, y-1, grid[x][y])
}

func TestVisitWaterGrid(t *testing.T) {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	lineList := strings.Fields(scanner.Text())
	N, _ := strconv.Atoi(lineList[0])
	M, _ := strconv.Atoi(lineList[1])

	grid := make([][]int, N)
	for i := 0; i < N; i++ {
		grid[i] = make([]int, M)
		scanner.Scan()
		lineList = strings.Fields(scanner.Text())

		for j := 0; j < M; j++ {
			grid[i][j], _ = strconv.Atoi(lineList[j])
		}
	}

	grid2 := [][]int{
		{1, 3, 1, 2, 4},
		{1, 2, 1, 3, 2},
		{2, 4, 7, 2, 1},
		{4, 5, 6, 1, 1},
		{1, 4, 1, 2, 1},
	} // 示例测试数据

	result := VisitWaterGrid(grid2)
	for _, arr := range result {
		fmt.Println(strconv.Itoa(arr[0]) + " " + strconv.Itoa(arr[1]))
	}
}
