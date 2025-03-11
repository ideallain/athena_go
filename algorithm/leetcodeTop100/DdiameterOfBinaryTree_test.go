package leetcodetop100

import (
	"athena_go/algorithm/common"
	"testing"
)

func diameterOfBinaryTree(root *common.TreeNode) int {
	maxPath = 0
	findPath(root)
	//这个题目没搞明白为什么这里要减一
	return maxPath - 1
}

var maxPath int

func findPath(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	leftLen := findPath(root.Left)
	rightLen := findPath(root.Right)

	curLen := leftLen
	if leftLen < rightLen {
		curLen = rightLen
	}
	curLen += 1
	if maxPath < (leftLen + rightLen + 1) {
		maxPath = leftLen + rightLen + 1
	}
	return curLen
}

func TestDiameterOfBinaryTree(t *testing.T) {
	root := &common.TreeNode{}
	root.Val = 1
	leftNode := &common.TreeNode{}
	leftNode.Val = 2
	root.Left = leftNode
	len := diameterOfBinaryTree(root)
	println("the len of root is ", len)
}
