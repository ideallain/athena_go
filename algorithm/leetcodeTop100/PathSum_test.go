package leetcodetop100

import (
	common "athena_go/algorithm/common"
)

func pathSum(root *common.TreeNode, targetSum int) int {
	count := 0
	if root == nil {
		return count
	}
	count += findSum(root, 0, targetSum)
	count += pathSum(root.Left, targetSum)
	count += pathSum(root.Right, targetSum)
	return count
}

func findSum(root *common.TreeNode, sum int, targetSum int) int {
	if root == nil {
		return 0
	}
	count := 0
	sum += root.Val
	if sum == targetSum {
		count++
	}
	count += findSum(root.Left, sum, targetSum)
	count += findSum(root.Right, sum, targetSum)
	return count
}
