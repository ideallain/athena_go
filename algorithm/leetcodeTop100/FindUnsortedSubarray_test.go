package leetcodetop100

import (
	"math"
	"testing"
)

func findUnsortedSubarray(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	//从左往右遍历，如果nums[i] < max ,记录下i，设置为区间的右边界
	//从右往左遍历，如果nums[i] > min,记录下i，设置为区间的左边界
	max := math.MinInt
	min := math.MaxInt
	var left int
	var right int
	for i := 0; i < len(nums); i++ {
		if nums[i] < max {
			println("in the index of ", i, " is less than left max ", max, " will adjustment ")
			right = i
		} else {
			println("record the max num in left ", nums[i])
			max = nums[i]
		}
	}
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] > min {
			println("in the index of ", i, " is more than right min ", min, " will adjustment ")
			left = i
		} else {
			println("record the min num in right ", nums[i])
			min = nums[i]
		}
	}
	if left < right {
		return right - left + 1
	}
	return 0
}

func TestFindUnsortedSubarray(t *testing.T) {
	nums := []int{2, 6, 4, 8, 10, 9, 15}
	count := findUnsortedSubarray(nums)
	println("the count is ", count)
}
