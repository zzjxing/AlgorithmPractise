package linked_list

import "AlgorithmPractise/common"
import "container/heap"

// MergeTwoOrderedLists 合并两个有序链表
//  1. 双指针指向每个链表头节点，将较小值的节点链接入新链表中
//  2. 直到某个链表元素全部被嵌入，再将另一个链表剩余节点链入新链表尾部
func MergeTwoOrderedLists(headA, headB *common.ListNode) *common.ListNode {
	dummy := &common.ListNode{}
	curr := dummy
	for headA != nil && headB != nil {
		if headA.Val < headB.Val {
			curr.Next = headA
			headA = headA.Next
		} else {
			curr.Next = headB
			headB = headB.Next
		}
		curr = curr.Next
	}
	if headA != nil {
		curr.Next = headA
	}
	if headB != nil {
		curr.Next = headB
	}
	return dummy.Next
}

// MergeKLists 合并 K 个有序链表
// 思路：1. 两两合并 2. 小根堆（以下为小根堆实现）
func MergeKLists(lists []*common.ListNode) *common.ListNode {
	h := &MinHeap{}
	heap.Init(h)
	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(h, lists[i])
		}
	}

	dummy := &common.ListNode{}
	curr := dummy
	for h.Len() > 0 {
		node := heap.Pop(h).(*common.ListNode)
		curr.Next = node
		curr = curr.Next
		if node.Next != nil {
			heap.Push(h, node.Next)
		}
	}
	return dummy.Next
}

type MinHeap []*common.ListNode

func (h *MinHeap) Len() int {
	return len(*h)
}
func (h *MinHeap) Less(i, j int) bool {
	return (*h)[i].Val < (*h)[j].Val
}
func (h *MinHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *MinHeap) Pop() interface{} {
	top := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return top
}
func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*common.ListNode))
}
