package linked_list

import "AlgorithmPractise/common"

// GetIntersectionNode 判断两个链表是否相交，如果相交，返回交点，否则返回 nil
// 双指针执行两个链表头节点，同步向后走，任意一个指针走到 nil 的时候，则指向另一个链表头节点，继续向右走，直到双指针值相同，即为交点(or nil)
// 首先是 nil 的情况，双指针每个都走了len1+len2步，同时走到空节点
// 如果相交，假设：x:链表 A 的长度，y:链表 B 的长度，z:链表的公共长度
// 指针 a 从链表 A 头节点，走到尾节点，再从链表 B 头节点出发，走到相交节点的时候，一共走了 x+y-z 步
// 指针 b 从链表 B 头节点，走到尾节点，再从链表 A 头节点出发，走到相交节点的时候，一共走了 y+x-z 步
// 因此在 x+y-z 步之后，双指针必定会在交点相遇
func GetIntersectionNode(headA, headB *common.ListNode) *common.ListNode {
	pA, pB := headA, headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}
	return pA
}
