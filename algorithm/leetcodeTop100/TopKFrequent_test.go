package leetcodetop100

import (
	"fmt"
	"sort"
	"testing"
)

func topKFrequent(nums []int, k int) []int {
	sortMap := make(map[int]int)
	for _, num := range nums {
		sortMap[num]++
	}
	type kv struct {
		key   int
		value int
	}
	println("the map is ", fmt.Sprintf("%v", sortMap))
	arr := make([]kv, 0)
	for num, count := range sortMap {
		arr = append(arr, kv{num, count})
	}
	println("the arr is ", fmt.Sprintf("%v", arr))
	sort.Slice(arr, func(i int, j int) bool {
		return arr[j].value < arr[i].value
	})
	println("the sorted arr is ", fmt.Sprintf("%v", arr))

	result := make([]int, 0)
	for i := range k {
		result = append(result, arr[i].key)
	}
	return result
}

func TestTopKFrequent(t *testing.T) {
	arr := []int{1, 1, 1, 2, 2, 3}
	result := topKFrequent(arr, 2)
	println("the result is ", fmt.Sprintf("%v", result))
}
