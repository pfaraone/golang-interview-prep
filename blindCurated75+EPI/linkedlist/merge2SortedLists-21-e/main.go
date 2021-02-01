/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	var dummyHead *ListNode = &ListNode{}
	dummyHead.Next = nil
	var curr *ListNode = dummyHead
	var p1 *ListNode = l1
	var p2 *ListNode = l2
	for p1 != nil && p2 != nil {
		if p1.Val <= p2.Val {
			curr.Next = p1
			p1 = p1.Next
		} else {
			curr.Next = p2
			p2 = p2.Next
		}
		curr = curr.Next
	}
	if p1 != nil {
		curr.Next = p1
	} else if p2 != nil {
		curr.Next = p2
	}
	return dummyHead.Next
}