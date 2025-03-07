package sword2offer

import common "athena_go/algorithm/common"

func HasSubtree(root1, root2 *common.TreeNode) bool {
	if root1 == nil || root2 == nil {
		return false
	}
	if root1.Val == root2.Val {
		if doesRoot2IsRoo1Sub(root1, root2) {
			return true
		}
	}
	return HasSubtree(root1.Left, root2) || HasSubtree(root1.Right, root2)
}

func doesRoot2IsRoo1Sub(root1, root2 *common.TreeNode) bool {
	if root2 == nil {
		return true
	}
	if root1 == nil {
		return false
	}
	if root1.Val != root2.Val {
		return false
	}
	return doesRoot2IsRoo1Sub(root1.Left, root2.Left) && doesRoot2IsRoo1Sub(root1.Right, root2.Right)
}
