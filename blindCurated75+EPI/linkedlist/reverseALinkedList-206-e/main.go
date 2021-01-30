/**
 * Definition for singly-linked list.
 *
 */

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		oldNext := curr.Next
		curr.Next = prev
		prev = curr
		curr = oldNext
	}
	return prev
}

func main() {

}
