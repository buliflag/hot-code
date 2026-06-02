package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList 迭代版（最终推荐）：三指针平移，逐节点翻转箭头
// 核心直觉：翻多米诺骨牌，每个节点把 Next 箭头拧向 prev
// 时间 O(n)，空间 O(1)
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		next := curr.Next  // 1. 记住下家，别走丢
		curr.Next = prev   // 2. 掰箭头，指向反转
		prev = curr        // 3. 三人组整体右移
		curr = next
	}
	return prev
}

// reverseListRecursive 递归版：信任子问题，只手动处理当前节点
// 核心直觉：假设 head.Next 之后已经反转好了，我只负责把 head 挂到末尾
// 时间 O(n)，空间 O(n)（递归栈）
func reverseListRecursive(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	newHead := reverseListRecursive(head.Next)
	head.Next.Next = head // 原本 head → next，现在 next → head
	head.Next = nil       // 断开旧箭头，head 成为新末尾
	return newHead
}
