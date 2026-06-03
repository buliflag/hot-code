package main

// ListNode 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// hasCycle 快慢指针法（最终推荐版）
// 核心直觉：操场跑步套圈。🐢一次1步，🐰一次2步，有环必相遇。
// 时间 O(n)，空间 O(1)
func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// hasCycleBasic 哈希表法（基础版）
// 核心：走过就标记，遇到重复说明有环
// 时间 O(n)，空间 O(n)
func hasCycleBasic(head *ListNode) bool {
	visited := make(map[*ListNode]bool)
	cur := head
	for cur != nil {
		if visited[cur] {
			return true
		}
		visited[cur] = true
		cur = cur.Next
	}
	return false
}
