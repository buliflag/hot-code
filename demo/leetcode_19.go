package main

// ListNode 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// removeNthFromEnd 双指针 + 哨兵法（最终推荐版）
// 核心直觉：拉开 n 格尺子，平移到尾部，尺尾指向的节点就是要删的。
// 时间 O(n)，空间 O(1)
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, dummy

	// 1. 拉开尺子：fast 先走 n 步
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// 2. 平移尺子：两人同步走，fast 停在最后一个节点（fast.Next == nil）
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	// 3. 动刀：跳过待删节点
	slow.Next = slow.Next.Next

	return dummy.Next
}

// removeNthFromEndBasic 两次遍历法（基础版）
// 核心：第一遍数总数，第二遍定位删除
// 时间 O(n)，空间 O(1)
func removeNthFromEndBasic(head *ListNode, n int) *ListNode {
	// 第一遍：计数
	length := 0
	for cur := head; cur != nil; cur = cur.Next {
		length++
	}

	// 要删的是第 length-n 个节点（从 0 开始）
	dummy := &ListNode{Next: head}
	cur := dummy
	for i := 0; i < length-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}
