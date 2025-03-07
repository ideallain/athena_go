package template

import common "athena_go/algorithm/common"

func inorderTraversal(root *common.TreeNode) []int {
	arr := []int{}
	var inOrder func(node *common.TreeNode)
	inOrder = func(node *common.TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		arr = append(arr, node.Val)
		inOrder(node.Right)
	}
	inOrder(root)
	return arr
}
