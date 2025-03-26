package linked_list

import (
	"AlgorithmPractise/common"
	"container/heap"
)

func SortList(head *common.ListNode) *common.ListNode {
	h := &MinListHeap{}
	heap.Init(h)
	curr := head
	for curr != nil {
		heap.Push(h, curr)
		curr = curr.Next
	}
	dummy := &common.ListNode{}
	curr = dummy
	for h.Len() > 0 {
		top := heap.Pop(h).(*common.ListNode)
		curr.Next = top
		curr = curr.Next
	}
	curr.Next = nil // 尾节点 Next 指针置空
	return dummy.Next
}

type MinListHeap []*common.ListNode

func (h *MinListHeap) Len() int {
	return len(*h)
}
func (h *MinListHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
func (h *MinListHeap) Less(i, j int) bool {
	return (*h)[i].Val < (*h)[j].Val
}
func (h *MinListHeap) Push(x interface{}) {
	*h = append(*h, x.(*common.ListNode))
}
func (h *MinListHeap) Pop() interface{} {
	l := len(*h)
	x := (*h)[l-1]
	*h = (*h)[:l-1]
	return x
}
