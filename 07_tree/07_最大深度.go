package tree

import "AlgorithmPractise/common"

func MaxDepth(root *common.TreeNode) int {
	var ret int
	var dfs func(node *common.TreeNode, depth int)
	dfs = func(node *common.TreeNode, depth int) {
		if node == nil {
			ret = max(ret, depth)
			return
		}
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, 0)
	return ret
}
