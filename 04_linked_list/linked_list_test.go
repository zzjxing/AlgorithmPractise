package linked_list

import (
	"AlgorithmPractise/common"
	"testing"
)

func TestHasCycle(t *testing.T) {
	node1 := &common.ListNode{Val: 1}
	node2 := &common.ListNode{Val: 2}
	node3 := &common.ListNode{Val: 3}
	node1.Next = node2
	node2.Next = node3
	node3.Next = node1 // 成环
	if !HasCycle(node1) {
		t.Errorf("HasCycle failed: expected true")
	}

	noCycle := common.BuildList([]int{1, 2, 3})
	if HasCycle(noCycle) {
		t.Errorf("HasCycle failed: expected false")
	}
}

func TestDetectCycle(t *testing.T) {
	nodes := make([]*common.ListNode, 6)
	for i := 0; i < 6; i++ {
		nodes[i] = &common.ListNode{Val: i}
		if i > 0 {
			nodes[i-1].Next = nodes[i]
		}
	}
	nodes[5].Next = nodes[3] // 成环入口 3
	entry := DetectCycle(nodes[0])
	if entry != nodes[3] {
		t.Errorf("DetectCycle failed: expected node %d, got %v", nodes[3].Val, entry)
	}
}

func TestReverseList(t *testing.T) {
	head := common.BuildList([]int{1, 2, 3, 4})
	got := ReverseList(head)
	want := common.BuildList([]int{4, 3, 2, 1})
	if !common.EqualList(got, want) {
		t.Errorf("ReverseList failed: got %v, want %v", got, want)
	}
}

func TestReverseBetween(t *testing.T) {
	head := common.BuildList([]int{1, 2, 3, 4, 5})
	got := ReverseBetween(head, 2, 4)
	want := common.BuildList([]int{1, 4, 3, 2, 5})
	if !common.EqualList(got, want) {
		t.Errorf("ReverseBetween failed: got %v, want %v", got, want)
	}
}

func TestReverseKGroup(t *testing.T) {
	head := common.BuildList([]int{1, 2, 3, 4, 5})
	got := ReverseKGroup(head, 2)
	want := common.BuildList([]int{2, 1, 4, 3, 5})
	if !common.EqualList(got, want) {
		t.Errorf("ReverseKGroup failed: got %v, want %v", got, want)
	}
}

func TestGetIntersectionNode(t *testing.T) {
	commonPart := common.BuildList([]int{8, 9, 10})

	a1 := &common.ListNode{Val: 1}
	a2 := &common.ListNode{Val: 2}
	a1.Next = a2
	a2.Next = commonPart

	b1 := &common.ListNode{Val: 3}
	b1.Next = commonPart

	got := GetIntersectionNode(a1, b1)
	if got != commonPart {
		t.Errorf("GetIntersectionNode failed: got %v, want %v", got, commonPart)
	}
}
