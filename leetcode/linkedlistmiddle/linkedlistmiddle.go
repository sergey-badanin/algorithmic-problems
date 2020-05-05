package linkedlistmiddle

//Solution for the problem:
//Given a non-empty, singly linked list, return a middle node of linked list.
//If there are two middle nodes, return the second middle node.
//
//The solution idea is inspired by the Floyd’s Cycle-Finding Algorithm
func middleNode(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

//ListNode is type provided by the Problem spec
type ListNode struct {
	Val  int
	Next *ListNode
}
