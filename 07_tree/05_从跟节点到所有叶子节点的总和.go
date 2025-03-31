package tree

import "AlgorithmPractise/common"

func SumNumbers(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	var ret int
	var dfs func(node *common.TreeNode, val int)
	dfs = func(node *common.TreeNode, val int) {
		if node.Left == nil && node.Right == nil {
			tmp := val*10 + node.Val
			ret += tmp
		}
		val = val*10 + node.Val
		if node.Left != nil {
			dfs(node.Left, val)
		}
		if node.Right != nil {
			dfs(node.Right, val)
		}
	}
	dfs(root, 0)
	return ret
}
