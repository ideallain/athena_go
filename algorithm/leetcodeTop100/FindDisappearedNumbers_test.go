package leetcodetop100

import (
	"math"
)

func findDisappearedNumbers(nums []int) []int {
	//1,2,3,4
	for _, num := range nums {
		tempIndex := math.Abs(float64(num)) - 1
		index := int64(tempIndex)
		if nums[index] > 0 {
			nums[index] = -nums[index]
		}
	}
	var result []int
	for index, value := range nums {
		if value > 0 {
			result = append(result, index+1)
		}
	}

	return result
}
