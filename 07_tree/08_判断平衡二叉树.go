package tree

import "AlgorithmPractise/common"

func IsBalanced(root *common.TreeNode) bool {
	var check func(node *common.TreeNode) int
	// 如果 node 是平衡二叉树，则返回其最大深度，否则返回-1
	check = func(node *common.TreeNode) int {
		if node == nil {
			return 0
		}
		left := check(node.Left)
		right := check(node.Right)
		if left == -1 || right == -1 {
			return -1
		}
		if common.Abs(left-right) > 1 {
			return -1
		}
		return max(left, right) + 1
	}
	return check(root) != -1
}
