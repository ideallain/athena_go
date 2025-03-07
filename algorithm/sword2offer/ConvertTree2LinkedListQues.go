package sword2offer

import common "athena_go/algorithm/common"

func ConvertTree2LinkedLst(root *common.TreeNode) *common.TreeNode {
	// write code here
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	left := ConvertTree2LinkedLst(root.Left)
	head := left
	for left != nil && left.Right != nil {
		left = left.Right
	}
	if left != nil {
		left.Right = root
		root.Left = left
	}
	right := ConvertTree2LinkedLst(root.Right)
	if right != nil {
		right.Left = root
		root.Right = right
	}
	if head == nil {
		return root
	}
	return head
}
