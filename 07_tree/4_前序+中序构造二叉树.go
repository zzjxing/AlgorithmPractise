package tree

import "AlgorithmPractise/common"

func BuildTreeFromInPreOrder(preOrder, inOrder []int) *common.TreeNode {
	indexMap := make(map[int]int)
	for i := range inOrder {
		indexMap[inOrder[i]] = i
	}
	var build func(preStart, preEnd, inStart, inEnd int) *common.TreeNode
	build = func(preStart, preEnd, inStart, inEnd int) *common.TreeNode {
		if preStart > preEnd {
			return nil
		}
		rootVal := preOrder[preStart]
		root := &common.TreeNode{Val: rootVal}
		rootIndex := indexMap[rootVal]
		leftSize := rootIndex - inStart

		root.Left = build(preStart+1, preStart+leftSize, inStart, rootIndex-1)
		root.Right = build(preStart+leftSize+1, preEnd, rootIndex+1, inEnd)
		return root

	}
	return build(0, len(preOrder)-1, 0, len(inOrder)-1)
}
