package tree

import (
	"AlgorithmPractise/common"
	"reflect"
	"sort"
	"testing"
)

func buildTestTree() *common.TreeNode {
	// 构造如下的二叉树：
	//        1
	//       / \
	//      2   3
	//     / \   \
	//    4   5   6
	return &common.TreeNode{
		Val: 1,
		Left: &common.TreeNode{
			Val:   2,
			Left:  &common.TreeNode{Val: 4},
			Right: &common.TreeNode{Val: 5},
		},
		Right: &common.TreeNode{
			Val:   3,
			Right: &common.TreeNode{Val: 6},
		},
	}
}

func TestInorder(t *testing.T) {
	root := buildTestTree()
	expected := []int{4, 2, 5, 1, 3, 6}
	if res := Inorder(root); !reflect.DeepEqual(res, expected) {
		t.Errorf("Inorder failed, expected %v, got %v", expected, res)
	}
}

func TestInorder2(t *testing.T) {
	root := buildTestTree()
	expected := []int{4, 2, 5, 1, 3, 6}
	if res := Inorder2(root); !reflect.DeepEqual(res, expected) {
		t.Errorf("Inorder2 failed, expected %v, got %v", expected, res)
	}
}

func TestLevelOrder(t *testing.T) {
	root := buildTestTree()
	expected := [][]int{{1}, {2, 3}, {4, 5, 6}}
	if res := LevelOrder(root); !reflect.DeepEqual(res, expected) {
		t.Errorf("LevelOrder failed, expected %v, got %v", expected, res)
	}
}

func TestZigLevelOrder(t *testing.T) {
	root := buildTestTree()
	expected := [][]int{{1}, {3, 2}, {4, 5, 6}}
	if res := ZigLevelOrder(root); !reflect.DeepEqual(res, expected) {
		t.Errorf("ZigLevelOrder failed, expected %v, got %v", expected, res)
	}
}

func TestRightView(t *testing.T) {
	root := buildTestTree()
	expected := []int{1, 3, 6}
	if res := RightView(root); !reflect.DeepEqual(res, expected) {
		t.Errorf("RightView failed, expected %v, got %v", expected, res)
	}
}

func findNodeByVal(root *common.TreeNode, val int) *common.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	if left := findNodeByVal(root.Left, val); left != nil {
		return left
	}
	return findNodeByVal(root.Right, val)
}

func TestLowestCommonAncestor(t *testing.T) {
	// 构建树：层序数组构建
	// 树结构与之前例子一致
	//        3
	//       / \
	//      5   1
	//     / \ / \
	//    6  2 0  8
	//      / \
	//     7   4
	data := []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}
	root := common.BuildTree(data)

	tests := []struct {
		p, q     int
		expected int
	}{
		{5, 1, 3},
		{6, 4, 5},
		{7, 8, 3},
		{0, 8, 1},
		{7, 4, 2},
	}

	for _, tt := range tests {
		p := findNodeByVal(root, tt.p)
		q := findNodeByVal(root, tt.q)
		want := tt.expected

		got := LowestCommonAncestor(root, p, q)
		if got == nil || got.Val != want {
			t.Errorf("LCA(%d, %d) = %v; want %d", tt.p, tt.q, got, want)
		}
	}
}

func TestMaxPathSum(t *testing.T) {
	tests := []struct {
		data     []int
		expected int
	}{
		//        1
		//       / \
		//      2   3
		// 最大路径：2 -> 1 -> 3，和为 6
		{[]int{1, 2, 3}, 6},

		//         -10
		//         /  \
		//        9   20
		//            / \
		//           15  7
		// 最大路径：15 -> 20 -> 7，和为 42
		{[]int{-10, 9, 20, -1, -1, 15, 7}, 42},

		// 只有一个节点
		{[]int{-3}, -3},

		//       2
		//      / \
		//     -1  3
		// 最大路径：3 -> 2，和为 5
		{[]int{2, -1, 3}, 5},
	}

	for _, tt := range tests {
		root := common.BuildTree(tt.data)
		res := MaxPathSum(root)
		if res != tt.expected {
			t.Errorf("MaxPathSum(%v) = %d; want %d", tt.data, res, tt.expected)
		}
	}
}

func TestBuildTreeFromInPreOrder(t *testing.T) {
	tests := []struct {
		preOrder []int
		inOrder  []int
	}{
		{
			preOrder: []int{3, 9, 20, 15, 7},
			inOrder:  []int{9, 3, 15, 20, 7},
		},
		{
			preOrder: []int{1, 2, 4, 5, 3},
			inOrder:  []int{4, 2, 5, 1, 3},
		},
		{
			preOrder: []int{1},
			inOrder:  []int{1},
		},
	}

	for _, tt := range tests {
		root := BuildTreeFromInPreOrder(tt.preOrder, tt.inOrder)
		got := Inorder(root)
		if !reflect.DeepEqual(got, tt.inOrder) {
			t.Errorf("Inorder(tree built from %v + %v) = %v; want %v", tt.preOrder, tt.inOrder, got, tt.inOrder)
		}
	}
}

func TestSumNumbers(t *testing.T) {
	tests := []struct {
		data     []int
		expected int
	}{
		//    1
		//   / \
		//  2   3
		// 路径 12 + 13 = 25
		{[]int{1, 2, 3}, 25},

		//    4
		//   / \
		//  9   0
		// /
		//5
		// 路径：495 + 40 = 535
		{[]int{4, 9, 0, 5, -1}, 535},

		// 空树
		{[]int{}, 0},

		// 单节点
		{[]int{7}, 7},
	}

	for _, tt := range tests {
		root := common.BuildTree(tt.data)
		got := SumNumbers(root)
		if got != tt.expected {
			t.Errorf("SumNumbers(%v) = %d; want %d", tt.data, got, tt.expected)
		}
	}
}

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		data     []int
		expected bool
	}{
		// 对称的树：
		//     1
		//   /   \
		//  2     2
		// / \   / \
		//3  4  4   3
		{[]int{1, 2, 2, 3, 4, 4, 3}, true},

		// 非对称的树：
		//     1
		//   /   \
		//  2     2
		//   \     \
		//   3      3
		{[]int{1, 2, 2, -1, 3, -1, 3}, false},

		// 空树
		{[]int{}, true},

		// 单节点
		{[]int{1}, true},

		// 不对称
		{[]int{1, 2, 2, -1, 3, -1, -1}, false},
	}

	for _, tt := range tests {
		root := common.BuildTree(tt.data)
		got := IsSymmetric(root)
		if got != tt.expected {
			t.Errorf("IsSymmetric(%v) = %v; want %v", tt.data, got, tt.expected)
		}
	}
}

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		data     []int
		expected int
	}{
		// 空树
		{[]int{}, 0},

		// 单节点
		{[]int{1}, 1},

		//     1
		//    / \
		//   2   3
		//  /
		// 4
		// 深度 = 3
		{[]int{1, 2, 3, 4}, 3},

		//     1
		//    /
		//   2
		//  /
		// 3
		///
		//4
		// 深度 = 4
		{[]int{1, 2, -1, 3, -1, 4}, 4},
	}

	for _, tt := range tests {
		root := common.BuildTree(tt.data)
		got := MaxDepth(root)
		if got != tt.expected {
			t.Errorf("MaxDepth(%v) = %d; want %d", tt.data, got, tt.expected)
		}
	}
}

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		data     []int
		expected bool
	}{
		// 空树
		{[]int{}, true},

		// 单节点
		{[]int{1}, true},

		// 完全平衡树
		//       1
		//     /   \
		//    2     2
		//   / \   / \
		//  3  4  4   3
		{[]int{1, 2, 2, 3, 4, 4, 3}, true},

		// 非平衡树
		//     1
		//    /
		//   2
		//  /
		// 3
		{[]int{1, 2, -1, 3}, false},

		// 深度差刚好为1，仍平衡
		{[]int{1, 2, 2, 3, -1, -1, 3}, true},
	}

	for _, tt := range tests {
		root := common.BuildTree(tt.data)
		got := IsBalanced(root)
		if got != tt.expected {
			t.Errorf("IsBalanced(%v) = %v; want %v", tt.data, got, tt.expected)
		}
	}
}

func TestDiameterOfBT(t *testing.T) {
	tests := []struct {
		data     []int
		expected int
	}{
		// 空树
		{[]int{}, 0},

		// 单节点
		{[]int{1}, 0},

		//     1
		//    / \
		//   2   3
		//  / \
		// 4   5
		// 最长路径：4 -> 2 -> 5，长度为 2（边数）
		// 或者 4 -> 2 -> 1 -> 3，长度为 3（边数）
		{[]int{1, 2, 3, 4, 5}, 3},

		//     1
		//    /
		//   2
		//  /
		// 3
		// 最长路径：3->2->1，长度 2
		{[]int{1, 2, -1, 3}, 2},

		// 平衡树
		//        1
		//      /   \
		//     2     3
		//    / \     \
		//   4   5     6
		//  /
		// 7
		// 最长路径：7->4->2->1->3->6，长度为 5
		{[]int{1, 2, 3, 4, 5, -1, 6, 7}, 5},
	}

	for _, tt := range tests {
		root := common.BuildTree(tt.data)
		got := DiameterOfBT(root)
		if got != tt.expected {
			t.Errorf("DiameterOfBT(%v) = %d; want %d", tt.data, got, tt.expected)
		}
	}
}

func TestWidthOfBT(t *testing.T) {
	tests := []struct {
		data     []int
		expected int
	}{
		// 空树
		{[]int{}, 0},

		// 单节点
		{[]int{1}, 1},

		//        1
		//       / \
		//      3   2
		//     /     \
		//    5       9
		//   /         \
		//  6           7
		// 最大宽度是 8：第4层从 index=0 的 6 到 index=7 的 7
		{[]int{1, 3, 2, 5, -1, -1, 9, 6, -1, -1, 7}, 8},

		// 完全二叉树：最大宽度 = 最后一层节点数 = 4
		//       1
		//     /   \
		//    2     3
		//   / \   / \
		//  4  5  6   7
		{[]int{1, 2, 3, 4, 5, 6, 7}, 4},

		// 非完全树，缺右节点：最大宽度 = 2
		{[]int{1, 2, -1, 4, -1}, 1},

		// 两层中间空节点拉长宽度：最大宽度 = 4
		//     1
		//    / \
		//   3   2
		//  /     \
		// 5       9
		{[]int{1, 3, 2, 5, -1, -1, 9}, 4},
	}

	for _, tt := range tests {
		root := common.BuildTree(tt.data)
		got := WidthOfBT(root)
		if got != tt.expected {
			t.Errorf("WidthOfBT(%v) = %d; want %d", tt.data, got, tt.expected)
		}
	}
}

func sort2D(nums [][]int) {
	sort.Slice(nums, func(i, j int) bool {
		for k := 0; k < len(nums[i]) && k < len(nums[j]); k++ {
			if nums[i][k] != nums[j][k] {
				return nums[i][k] < nums[j][k]
			}
		}
		return len(nums[i]) < len(nums[j])
	})
}

func TestPathSum(t *testing.T) {
	tests := []struct {
		data     []int
		target   int
		expected [][]int
	}{
		// 树结构：
		//     5
		//    / \
		//   4   8
		//  /   / \
		// 11  13  4
		// / \      \
		//7  2       1
		//
		// target = 22
		// 路径：[5,4,11,2] 和 [5,8,4,5]
		{[]int{5, 4, 8, 11, -1, 13, 4, 7, 2, -1, -1, -1, 1}, 22, [][]int{
			{5, 4, 11, 2},
		}},

		// 空树
		{[]int{}, 0, nil},

		// 单节点等于目标
		{[]int{1}, 1, [][]int{{1}}},

		// 单节点不等于目标
		{[]int{1}, 2, nil},

		// 多路径匹配
		//       1
		//      / \
		//     2   3
		// target = 3
		{[]int{1, 2, 3}, 3, [][]int{{1, 2}}},
	}

	for _, tt := range tests {
		root := common.BuildTree(tt.data)
		got := PathSum(root, tt.target)

		// 排序结果，防止顺序导致对比失败
		sort2D(got)
		sort2D(tt.expected)

		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("PathSum(%v, %d) = %v; want %v", tt.data, tt.target, got, tt.expected)
		}
	}
}

func TestIsBST(t *testing.T) {
	tests := []struct {
		data     []int
		expected bool
	}{
		// 空树
		{[]int{}, true},

		// 单节点
		{[]int{1}, true},

		// 有效 BST
		//      2
		//     / \
		//    1   3
		{[]int{2, 1, 3}, true},

		// 非 BST：右子树节点小于根
		//      5
		//     / \
		//    1   4
		//       / \
		//      3   6
		{[]int{5, 1, 4, -1, -1, 3, 6}, false},

		// 有效 BST（多层）
		//        10
		//       /  \
		//      5   15
		//         /  \
		//        11  20
		{[]int{10, 5, 15, -1, -1, 11, 20}, true},
	}

	for _, tt := range tests {
		root := common.BuildTree(tt.data)
		got := IsBST(root)
		if got != tt.expected {
			t.Errorf("IsBST(%v) = %v; want %v", tt.data, got, tt.expected)
		}
	}
}
