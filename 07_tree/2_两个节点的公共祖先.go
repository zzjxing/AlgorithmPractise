package tree

import "AlgorithmPractise/common"

// LowestCommonAncestor LCA问题，寻找两个节点的最近公共祖先
// 递归实现，后续遍历所有节点，先在左右子树分别寻找 p 或者 q
// 如果都找到了，说明 p 和 q 分布在当前节点的两侧，也即当前节点就是 p 和 q 的最近祖先，返回当前节点即可
// 否则返回找到的那一侧
// 递归终止：当前节点就是 p 或者 q，或者当前节点为nil
func LowestCommonAncestor(root, p, q *common.TreeNode) *common.TreeNode {
	var findLCA func(node *common.TreeNode) *common.TreeNode
	findLCA = func(node *common.TreeNode) *common.TreeNode {
		if node == nil || node == p || node == q {
			return node
		}
		left := findLCA(node.Left)
		right := findLCA(node.Right)
		if left != nil && right != nil {
			return node
		}
		if left != nil {
			return left
		}
		return right
	}
	return findLCA(root)
}
