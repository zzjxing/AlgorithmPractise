package common

import (
	"fmt"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// BuildList 创建链表
func BuildList(nums []int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, val := range nums {
		curr.Next = &ListNode{Val: val}
		curr = curr.Next
	}
	return dummy.Next
}

// PrintList 打印链表
func PrintList(head *ListNode) {
	var sb strings.Builder
	for head != nil {
		sb.WriteString(fmt.Sprintf("%d -> ", head.Val))
		head = head.Next
	}
	sb.WriteString("nil")
	fmt.Println(sb.String())
}

// LenOfList 返回链表的长度
func LenOfList(head *ListNode) int {
	count := 0
	for head != nil {
		count++
		head = head.Next
	}
	return count
}

// EqualList 判断两个链表是否相同
func EqualList(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

// ListToSlice 将链表转为切片
func ListToSlice(head *ListNode) []int {
	var res []int
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	return res
}
