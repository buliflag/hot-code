package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {

	var prev *ListNode
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev

		//移动位置
		prev = curr
		curr = next
	}
	return prev

}
