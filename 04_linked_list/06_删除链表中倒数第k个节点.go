package linked_list

import "AlgorithmPractise/common"

// RemoveKthFromEnd 删除链表的倒数第 K 个节点
// 思路，快慢指针，快指针先向前走 k 步，然后快慢同步向前，快指针走到结尾的时候，慢指针就在第 k 个节点了
func RemoveKthFromEnd(head *common.ListNode, k int) *common.ListNode {
	dummy := &common.ListNode{Next: head}
	slow, fast := dummy, dummy
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}
