package tree

import "AlgorithmPractise/common"

func DiameterOfBT(root *common.TreeNode) int {
	var ret int
	var dfs func(node *common.TreeNode) int
	dfs = func(node *common.TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		ret = max(right+left, ret)
		return max(left, right) + 1
	}
	dfs(root)
	return ret
}
