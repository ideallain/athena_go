package leetcodetop100

//这道题目和目标和那道非常像
func canPartition(nums []int) bool {
	sum := 0
	for _, value := range nums {
		sum += value
	}
	if sum%2 == 1 {
		return false
	}
	sum /= 2
	//dp[i]表示容量为i的数组最多能放进的物品之和
	//dp[i] = math.Max(dp[i - nums[j]]), dp[i]) 遍历j，不断更新do[i]的组合个数
	//0-1背包，组合问题，外层遍历物品
	dp := make([]int, sum+1)
	dp[0] = 0
	for _, num := range nums {
		for j := sum; j >= num; j-- {
			if dp[j-num]+num > dp[j] {
				dp[j] = dp[j-num] + num
			}
		}
	}
	return dp[sum] == sum
}
