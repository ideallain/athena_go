package template

import "fmt"

func buildMaxIsland(grid [][]int) int {
	result := 0
	if grid == nil || len(grid) == 0 {
		return result
	}
	hashMap := make(map[int]int)
	islandArea := 0
	gridMark := 2
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				islandArea = 0
				visitGrid4(grid, i, j, &islandArea, gridMark)
				hashMap[gridMark] = islandArea
				if islandArea > result {
					result = islandArea
				}
				gridMark++
			}
		}
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 0 {
				sum := findNearByIsland(grid, i, j, hashMap)
				if sum > result {
					result = sum
				}
			}
		}
	}

	return result
}

func visitGrid4(grid [][]int, x int, y int, preValue *int, gridMark int) {
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[0]) {
		return
	}
	//可以省略掉visited的原因是因为grid的每个位置如果被遍历过就会被标记为gridMark，能识别出来是不是被遍历过
	if grid[x][y] != 1 {
		return
	}
	// visited[x][y] = true
	grid[x][y] = gridMark
	*preValue++
	visitGrid4(grid, x+1, y, preValue, gridMark)
	visitGrid4(grid, x-1, y, preValue, gridMark)
	visitGrid4(grid, x, y+1, preValue, gridMark)
	visitGrid4(grid, x, y-1, preValue, gridMark)
}

func findNearByIsland(grid [][]int, x, y int, hashMap map[int]int) int {
	markSet := make(map[int]bool)
	sum := 1
	coordinate := [][]int{{x + 1, y}, {x - 1, y}, {x, y + 1}, {x, y - 1}}
	for _, arr := range coordinate {
		if arr[0] < 0 || arr[1] < 0 || arr[0] >= len(grid) || arr[1] >= len(grid[0]) {
			continue
		}
		if grid[arr[0]][arr[1]] == 0 {
			continue
		}
		gridMark := grid[arr[0]][arr[1]]
		if !markSet[gridMark] {
			markSet[gridMark] = true
			sum += hashMap[gridMark]
		}
	}
	return sum
}

func TestBuildMaxIsland() {
	var m, n int
	fmt.Scan(&m, &n)

	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
		for j := range grid[i] {
			fmt.Scan(&grid[i][j])
		}
	}

	sum := buildMaxIsland(grid)
	fmt.Println(sum)
}
