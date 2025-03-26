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

func TestReorderList(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 4, 2, 3}},
		{[]int{1, 2, 3, 4, 5}, []int{1, 5, 2, 4, 3}},
		{[]int{1}, []int{1}},
		{[]int{1, 2}, []int{1, 2}},
	}

	for _, tt := range tests {
		head := common.BuildList(tt.input)
		ReorderList(head)
		want := common.BuildList(tt.expected)

		if !common.EqualList(head, want) {
			t.Errorf("ReorderList(%v) failed: got %v, want %v", tt.input, common.ListToSlice(head), tt.expected)
		}
	}
}

func TestMergeTwoOrderedLists(t *testing.T) {
	tests := []struct {
		a, b     []int
		expected []int
	}{
		{[]int{1, 3, 5}, []int{2, 4, 6}, []int{1, 2, 3, 4, 5, 6}},
		{[]int{1, 2}, []int{}, []int{1, 2}},
		{[]int{}, []int{3, 4}, []int{3, 4}},
		{[]int{}, []int{}, []int{}},
	}

	for _, tt := range tests {
		l1 := common.BuildList(tt.a)
		l2 := common.BuildList(tt.b)
		got := MergeTwoOrderedLists(l1, l2)
		want := common.BuildList(tt.expected)
		if !common.EqualList(got, want) {
			t.Errorf("MergeTwoOrderedLists(%v, %v) got %v, want %v", tt.a, tt.b, common.ListToSlice(got), tt.expected)
		}
	}
}

func TestMergeKLists(t *testing.T) {
	tests := []struct {
		input    [][]int
		expected []int
	}{
		{
			input:    [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}},
			expected: []int{1, 1, 2, 3, 4, 4, 5, 6},
		},
		{
			input:    [][]int{},
			expected: []int{},
		},
		{
			input:    [][]int{{}},
			expected: []int{},
		},
	}

	for _, tt := range tests {
		var listNodes []*common.ListNode
		for _, arr := range tt.input {
			listNodes = append(listNodes, common.BuildList(arr))
		}
		got := MergeKLists(listNodes)
		want := common.BuildList(tt.expected)
		if !common.EqualList(got, want) {
			t.Errorf("MergeKLists(%v) got %v, want %v", tt.input, common.ListToSlice(got), tt.expected)
		}
	}
}

func TestRemoveKthFromEnd(t *testing.T) {
	tests := []struct {
		input    []int
		k        int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, 2, []int{1, 2, 3, 5}}, // 删除 4
		{[]int{1}, 1, []int{}},                       // 删除 1
		{[]int{1, 2}, 1, []int{1}},                   // 删除 2
		{[]int{1, 2}, 2, []int{2}},                   // 删除 1
		{[]int{1, 2, 3}, 3, []int{2, 3}},             // 删除头部
	}

	for _, tt := range tests {
		head := common.BuildList(tt.input)
		got := RemoveKthFromEnd(head, tt.k)
		want := common.BuildList(tt.expected)

		if !common.EqualList(got, want) {
			t.Errorf("RemoveKthFromEnd(%v, %d) = %v, want %v", tt.input, tt.k, common.ListToSlice(got), tt.expected)
		}
	}
}

func TestDeleteDuplicates(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 2, 3, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
		{[]int{1, 1, 1, 1}, []int{1}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{}, []int{}},
	}

	for _, tt := range tests {
		head := common.BuildList(tt.input)
		got := DeleteDuplicates(head)
		want := common.BuildList(tt.expected)

		if !common.EqualList(got, want) {
			t.Errorf("DeleteDuplicates(%v) = %v, want %v", tt.input, common.ListToSlice(got), tt.expected)
		}
	}
}

func TestDeleteDuplicates2(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 2, 3, 3, 4, 5}, []int{1, 4, 5}},
		{[]int{1, 1, 2, 3, 3}, []int{2}},
		{[]int{1, 1, 1, 1}, []int{}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{}, []int{}},
	}

	for _, tt := range tests {
		head := common.BuildList(tt.input)
		got := DeleteDuplicates2(head)
		want := common.BuildList(tt.expected)

		if !common.EqualList(got, want) {
			t.Errorf("DeleteDuplicates2(%v) = %v, want %v", tt.input, common.ListToSlice(got), tt.expected)
		}
	}
}

// 测试链表排序
func TestSortList(t *testing.T) {
	input := common.BuildList([]int{4, 2, 1, 3})
	got := SortList(input)
	want := common.BuildList([]int{1, 2, 3, 4})

	if !common.EqualList(got, want) {
		t.Errorf("SortList failed: got %v, want %v", common.ListToSlice(got), common.ListToSlice(want))
	}
}

// 测试链表加法
func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1, l2   []int
		expected []int
	}{
		{[]int{2, 4, 3}, []int{5, 6, 4}, []int{7, 0, 8}}, // 342 + 465 = 807
		{[]int{0}, []int{0}, []int{0}},                   // 0 + 0 = 0
		{[]int{9, 9, 9}, []int{1}, []int{0, 0, 0, 1}},    // 999 + 1 = 1000
	}

	for _, tt := range tests {
		l1 := common.BuildList(tt.l1)
		l2 := common.BuildList(tt.l2)
		got := AddTwoNumbers(l1, l2)
		want := common.BuildList(tt.expected)

		if !common.EqualList(got, want) {
			t.Errorf("AddTwoNumbers(%v, %v) = %v, want %v",
				tt.l1, tt.l2, common.ListToSlice(got), tt.expected)
		}
	}
}

// 测试链表是否为回文
func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		input    []int
		expected bool
	}{
		{[]int{1, 2, 2, 1}, true},
		{[]int{1, 2}, false},
		{[]int{1, 2, 1}, true},
		{[]int{1}, true},
		{[]int{}, true},
	}

	for _, tt := range tests {
		head := common.BuildList(tt.input)
		got := IsPalindrome(head)
		if got != tt.expected {
			t.Errorf("IsPalindrome(%v) = %v, want %v", tt.input, got, tt.expected)
		}
	}
}
