package linked_list

import "AlgorithmPractise/common"

func AddTwoNumbers(l1, l2 *common.ListNode) *common.ListNode {
	dummy := &common.ListNode{}
	carry := 0
	curr := dummy
	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		carry = sum / 10
		curr.Next = &common.ListNode{Val: sum % 10}
		curr = curr.Next
	}
	return dummy.Next
}
