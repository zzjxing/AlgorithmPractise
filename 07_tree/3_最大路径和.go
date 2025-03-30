package tree

import (
	"AlgorithmPractise/common"
	"math"
)

// MaxPathSum
// 思路：后续遍历，计算左子树的最大路径和leftSum 和 右子树的最大路径和rightSum
// 更新返回值 ret，ret=max(ret,leftSum+rightSum+node.Val)
// 返回当前节点的最大贡献值：max(leftSum,rightSum)+node.Val
func MaxPathSum(root *common.TreeNode) int {
	var ret = math.MinInt32
	var dfs func(node *common.TreeNode) int
	dfs = func(node *common.TreeNode) int {
		if node == nil {
			return 0
		}
		leftSum := max(dfs(node.Left), 0)
		rightSum := max(dfs(node.Right), 0)
		ret = max(ret, leftSum+rightSum+node.Val)
		return max(leftSum, rightSum) + node.Val
	}
	dfs(root)
	return ret
}
