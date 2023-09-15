package main

// 92. Reverse Linked List II
//
// https://leetcode.com/problems/reverse-linked-list-ii
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}
	dummy := &ListNode{Val: 0, Next: head}
	prev := dummy
	for i := 1; i < left; i++ {
		prev = prev.Next
	}
	current := prev.Next
	for i := 0; i < right-left; i++ {
		buf := prev.Next
		prev.Next = current.Next
		current.Next = current.Next.Next
		prev.Next.Next = buf
	}

	return dummy.Next
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
