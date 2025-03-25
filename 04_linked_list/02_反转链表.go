package linked_list

import (
	"AlgorithmPractise/common"
)

// ReverseList 整体反转链表 循环实现
func ReverseList(head *common.ListNode) *common.ListNode {
	curr := head
	var pre *common.ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	return pre
}

// ReverseBetween 指定位置旋转链表
// 反转链表的第 left 到 right 的节点
func ReverseBetween(head *common.ListNode, left, right int) *common.ListNode {
	if head == nil || right == left {
		return head
	}
	dummy := &common.ListNode{Next: head}
	prev := dummy
	for i := 1; i < left; i++ {
		prev = prev.Next
	}
	curr := prev.Next
	// 每次循环将curr插入到 prev 之后，然后 curr 后移
	for i := left; i < right; i++ {
		next := curr.Next
		curr.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}
	return dummy.Next
}

// ReverseKGroup 每 k 个节点为一组，进行链表反转
// 思路：递归反转
//
//	先从 head 反转 k 个节点，然后再以第 k+1 个节点为 head，反转 k 个节点
//	如此循环，直到最后的节点数量小于k
func ReverseKGroup(head *common.ListNode, k int) *common.ListNode {
	count := common.Length(head)
	if count < k { // 递归终止
		return head
	}
	curr := head
	var prev, next *common.ListNode
	for i := 0; i < k; i++ {
		next = curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	if curr != nil {
		head.Next = ReverseKGroup(curr, k)
	}
	return prev
}
