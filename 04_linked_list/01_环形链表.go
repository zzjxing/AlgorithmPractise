package linked_list

import common "AlgorithmPractise/common"

// HasCycle 检测环：快慢指针
func HasCycle(head *common.ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// DetectCycle 寻找入环节点
// 方法：快慢指针，Floyd判圈算法
// 流程：
//  1. 使用快慢指针，检测环的存在
//  2. 指针相遇时，令其中一个指针指向头，双指针同步向前走，再次相遇即为入口
//
// 推理：
//
//		设：a=头节点到入环点的距离，b=入换点到相遇的距离，c=环的剩余长度  b+c就是环的长度
//		快慢指针相遇的时候，快指针是慢指针路程的 2 倍
//			快指针路程：a+b+k(b+c)  慢指针路程：a+b， ps：快指针可能走了 k 圈才和慢指针相遇
//			2(a+b)=a+b+k(b+c) --> a=(k-1)(b+c)+c
//		因此：a 等于 环的整数倍 加上 环的剩余时长，
//		此时让一个指针从头开始，一个指针从相遇点开始向前行
//	 当一个指针走了 a 步的时候，会到达环入口，另一个指针走了 c步，到环入口，又走了k-1圈，还是环入口
//	 因此双指针再次相遇就是环入口
func DetectCycle(head *common.ListNode) *common.ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			break
		}
	}
	if fast == nil || fast.Next == nil {
		return nil
	}
	fast = head
	for fast != slow {
		fast = fast.Next
		slow = slow.Next
	}
	return fast
}
