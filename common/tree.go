package common

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BuildTree 用层序构造二叉树（nil 用 -1 表示）
func BuildTree(data []int) *TreeNode {
	if len(data) == 0 || data[0] == -1 {
		return nil
	}

	root := &TreeNode{Val: data[0]}
	queue := []*TreeNode{root}
	i := 1

	for i < len(data) {
		node := queue[0]
		queue = queue[1:]

		// 左孩子
		if data[i] != -1 {
			node.Left = &TreeNode{Val: data[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i >= len(data) {
			break
		}

		// 右孩子
		if data[i] != -1 {
			node.Right = &TreeNode{Val: data[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// TreeToSlice 将树转为层序数组（nil 节点为 -1）
func TreeToSlice(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var res []int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node != nil {
			res = append(res, node.Val)
			queue = append(queue, node.Left, node.Right)
		} else {
			res = append(res, -1)
		}
	}

	// 清除末尾多余的 -1
	for len(res) > 0 && res[len(res)-1] == -1 {
		res = res[:len(res)-1]
	}
	return res
}

// IsSameTree 判断两棵树是否完全相同
func IsSameTree(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil || p.Val != q.Val {
		return false
	}
	return IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

// MaxDepth 返回二叉树最大深度
func MaxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := MaxDepth(root.Left)
	right := MaxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}

// PrintTree 层序打印树结构（调试用）
func PrintTree(root *TreeNode) {
	if root == nil {
		fmt.Println("nil")
		return
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			if node == nil {
				fmt.Print("nil ")
				continue
			}
			fmt.Print(node.Val, " ")
			queue = append(queue, node.Left, node.Right)
		}
		fmt.Println()
	}
}
