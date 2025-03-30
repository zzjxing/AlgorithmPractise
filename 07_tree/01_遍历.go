package tree

import "AlgorithmPractise/common"

// Inorder 中序遍历-递归实现
func Inorder(root *common.TreeNode) []int {
	var ret []int
	var dfs func(node *common.TreeNode)
	dfs = func(node *common.TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		ret = append(ret, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return ret
}

// Inorder2 中序遍历-循环实现
func Inorder2(root *common.TreeNode) []int {
	var ret []int
	var stack []*common.TreeNode
	curr := root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ret = append(ret, curr.Val)
		curr = curr.Right
	}
	return ret
}

// LevelOrder 层次遍历
func LevelOrder(root *common.TreeNode) [][]int {
	var ret [][]int
	if root == nil {
		return nil
	}
	var queue []*common.TreeNode
	queue = append(queue, root)
	for len(queue) != 0 {
		l := len(queue)
		currLevel := make([]int, 0)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			currLevel = append(currLevel, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ret = append(ret, currLevel)
	}
	return ret
}

// ZigLevelOrder 树的锯齿形遍历
func ZigLevelOrder(root *common.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var ret [][]int
	queue := []*common.TreeNode{root}
	left2right := true
	for len(queue) != 0 {
		l := len(queue)
		currLevel := make([]int, 0)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			currLevel = append(currLevel, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if !left2right {
			currLevel = common.ReverseSlice(currLevel)
		}
		ret = append(ret, currLevel)
		left2right = !left2right
	}
	return ret
}

// RightView 二叉树的右视图
func RightView(root *common.TreeNode) []int {
	if root == nil {
		return nil
	}
	var ret []int
	queue := []*common.TreeNode{root}
	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			if i == l-1 {
				ret = append(ret, node.Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return ret
}
