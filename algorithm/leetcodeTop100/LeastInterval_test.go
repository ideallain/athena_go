package leetcodetop100

import "sort"

func leastInterval(task []byte, n int) int {
	countArr := make([]int, 26)
	for _, value := range task {
		countArr[value-'A']++
	}
	sort.Ints(countArr)
	maxCount := 1
	for i := 24; i >= 0; i-- {
		if countArr[i] != countArr[25] {
			break
		}
		maxCount++
	}
	result := (countArr[25]-1)*(n+1) + maxCount
	if result < len(task) {
		result = len(task)
	}
	return result
}
