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
