package main

// 725. Split Linked List in Parts
//
// https://leetcode.com/problems/split-linked-list-in-parts
func splitListToParts(head *ListNode, k int) []*ListNode {
	var ans []*ListNode
	var prev, current *ListNode
	totalLen := length(head)
	subLen, add := totalLen/k, totalLen%k
	prev, current = nil, head
	for i := 0; i < k; i++ {
		ans = append(ans, current)
		size := subLen
		if add > 0 {
			size++
			add--
		}
		for j := 0; j < size; j++ {
			prev, current = current, current.Next
		}

		if prev != nil {
			prev.Next = nil
		}
	}
	return ans
}

func length(node *ListNode) int {
	var size int
	for node != nil {
		node, size = node.Next, size+1
	}
	return size
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}
