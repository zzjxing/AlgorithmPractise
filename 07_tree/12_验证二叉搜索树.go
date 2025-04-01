package tree

import (
	"AlgorithmPractise/common"
)

func IsBST(root *common.TreeNode) bool {
	var prev *common.TreeNode
	var dfs func(node *common.TreeNode) bool
	dfs = func(node *common.TreeNode) bool {
		if node == nil {
			return true
		}
		if !dfs(node.Left) {
			return false
		}
		if prev != nil && node.Val <= prev.Val {
			return false
		}
		prev = node
		return dfs(node.Right)
	}
	return dfs(root)
}
