package linked_list

import "AlgorithmPractise/common"

// DeleteDuplicates 删除链表中的重复元素，保留一个 e.g. 1-2-2-3-3-4-5. -> 1-2-3-4-5
func DeleteDuplicates(head *common.ListNode) *common.ListNode {
	if head == nil {
		return nil
	}
	curr := head
	for curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}

// DeleteDuplicates2 完全删除链表中的重复元素，不保留，e.g. 1-2-2-3-3-4-5. -> 1-4-5
func DeleteDuplicates2(head *common.ListNode) *common.ListNode {
	if head == nil {
		return nil
	}
	dummy := &common.ListNode{Next: head}
	curr := head
	prev := dummy
	for curr != nil {
		if curr.Next != nil && curr.Val == curr.Next.Val {
			for curr.Next != nil && curr.Next.Val == curr.Val {
				curr = curr.Next
			}
			prev.Next = curr.Next
			curr = curr.Next
		} else {
			prev = prev.Next
			curr = curr.Next
		}
	}
	return dummy.Next
}
