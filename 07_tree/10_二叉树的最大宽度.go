package tree

import "AlgorithmPractise/common"

func WidthOfBT(root *common.TreeNode) int {
	if root == nil {
		return 0
	}
	type NodeInfo struct {
		node *common.TreeNode
		idx  int
	}
	var queue []NodeInfo
	var ret int
	queue = append(queue, NodeInfo{root, 0})
	for len(queue) > 0 {
		size := len(queue)
		start := queue[0].idx
		end := queue[len(queue)-1].idx
		ret = max(end-start+1, ret)
		for i := 0; i < size; i++ {
			ni := queue[0]
			queue = queue[1:]
			if ni.node.Left != nil {
				queue = append(queue, NodeInfo{ni.node.Left, 2 * (ni.idx - start)})
			}
			if ni.node.Right != nil {
				queue = append(queue, NodeInfo{ni.node.Right, 2*(ni.idx-start) + 1})
			}
		}
	}
	return ret
}
