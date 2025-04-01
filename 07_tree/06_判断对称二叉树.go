package tree

import "AlgorithmPractise/common"

func IsSymmetric(root *common.TreeNode) bool {
	if root == nil {
		return true
	}
	var check func(left, right *common.TreeNode) bool
	check = func(left, right *common.TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil {
			return false
		}
		if left.Val != right.Val {
			return false
		}
		return check(left.Left, right.Right) && check(left.Right, right.Left)
	}
	return check(root.Left, root.Right)
}
