package leetcodetop100

func findTargetSumWays(nums []int, target int) int {
	//2sum(p) = target + sum(nums)
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	if (target+sum)%2 == 1 || target > sum {
		return 0
	}
	sum = (target + sum) / 2
	if sum < 0 {
		return 0
	}
	dp := make([]int, sum+1)
	dp[0] = 1
	for _, num := range nums {
		for i := sum; i >= num; i-- {
			dp[i] += dp[i-num]
		}
	}
	return dp[sum]
}
