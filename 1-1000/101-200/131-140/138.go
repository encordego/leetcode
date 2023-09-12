package main

// 138. Copy List with Random Pointer
//
// https://leetcode.com/problems/copy-list-with-random-pointer
func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	storage, current := make(map[*Node]*Node), head
	for current != nil {
		storage[current] = &Node{Val: current.Val}
		current = current.Next
	}
	current = head
	for current != nil {
		storage[current].Next = storage[current.Next]
		storage[current].Random = storage[current.Random]
		current = current.Next
	}

	return storage[head]
}

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}
