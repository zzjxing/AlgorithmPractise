package tree

import "AlgorithmPractise/common"

func PathSum(root *common.TreeNode, target int) [][]int {
	var ret [][]int
	var dfs func(node *common.TreeNode, path []int, target int)
	dfs = func(node *common.TreeNode, path []int, target int) {
		if node == nil {
			return
		}
		path = append(path, node.Val)
		target -= node.Val
		if node.Left == nil && node.Right == nil && target == 0 {
			tmp := append([]int{}, path...)
			ret = append(ret, tmp)
		}
		dfs(node.Left, path, target)
		dfs(node.Right, path, target)
	}
	dfs(root, []int{}, target)
	return ret
}
