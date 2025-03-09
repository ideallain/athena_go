package leetcodetop100

func subarraySum(nums []int, target int) int {
	prefixMap := make(map[int]int)
	prefixMap[0] = 1
	prefixSum := make([]int, len(nums))
	prefixSum[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i]
	}
	count := 0
	for i := 0; i < len(nums); i++ {
		index := prefixSum[i] - target
		if value, exists := prefixMap[index]; exists {
			count += value
		}
		prefixMap[prefixSum[i]]++
	}
	return count
}
