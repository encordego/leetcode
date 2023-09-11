package main

// 141. Linked List Cycle
//
// https://leetcode.com/problems/linked-list-cycle/
func hasCycle(head *ListNode) bool {
	visited := make(map[*ListNode]bool)
	current := head
	for current != nil {
		if visited[current] {
			return true
		}
		visited[current] = true
		current = current.Next
	}
	return false
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
