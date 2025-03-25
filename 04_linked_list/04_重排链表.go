package linked_list

import "AlgorithmPractise/common"

// ReorderList 重排链表
// 将 L(0) → L(1) → … → L(n-1) → L(n)
// 重排为：L(0) → L(n) → L(1)→ L(n - 1) → L(2) → L(n - 2) → …
// 1. 快慢指针，切割链表 2. 反转后面的链表 3. 顺序嵌入生成新链表
func ReorderList(head *common.ListNode) {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	first, second := head, reverse(slow.Next)
	slow.Next = nil
	for second != nil {
		tmp1, tmp2 := first.Next, second.Next
		first.Next = second
		second.Next = tmp1
		first, second = tmp1, tmp2
	}
}

func reverse(head *common.ListNode) *common.ListNode {
	var prev *common.ListNode
	curr := head
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
