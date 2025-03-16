package leetcodetop100

import (
	common "athena_go/algorithm/common"
)

func rob(root *common.TreeNode) int {
	//dp数组，表示某个节点偷或者不偷能够获取的最大金额
	//如果dp[i]节点偷，那么他的子节点肯定不能偷 dp[i] = root[i].val + left不偷 + right不偷
	//如果dp[i]节点不偷，那么他的子节点可以选择偷，也可以选择不偷，但是我们直接返回每个节点偷或者不偷时的最大值，来给上面的递归函数选择
	result := robTree(root)
	return getMax(result[0], result[1])
}

func robTree(root *common.TreeNode) [2]int {
	//0表示当前节点不偷
	//1表示当前节点偷
	result := [2]int{}
	if root == nil {
		return result
	}
	leftResult := robTree(root.Left)
	rightResult := robTree(root.Right)
	result[1] = root.Val + leftResult[0] + rightResult[0]
	result[0] = getMax(leftResult[0], leftResult[1]) + getMax(rightResult[0], rightResult[1])
	return result
}

func getMax(i int, j int) int {
	if i > j {
		return i
	}
	return j
}
