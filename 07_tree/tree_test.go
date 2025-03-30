package tree

import (
	"AlgorithmPractise/common"
	"reflect"
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
