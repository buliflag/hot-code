package main

// ListNode 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// detectCycle 快慢指针两阶段法（最终推荐版）
// 核心直觉：第1阶段🐢🐰相遇；第2阶段🐢回起点，两人同速走，再次相遇即入口。
// 时间 O(n)，空间 O(1)
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head

	// 第1阶段：判断有环 + 找到相遇点
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			// 有环！进入第2阶段
			slow = head // 🐢 回到起点
			for slow != fast {
				slow = slow.Next // 两人都走 1 步
				fast = fast.Next
			}
			return slow // 再次相遇 = 环入口 🎯
		}
	}

	return nil // 无环
}

// detectCycleBasic 哈希表法（基础版）
// 核心：第一个被重复访问的节点就是环入口
// 时间 O(n)，空间 O(n)
func detectCycleBasic(head *ListNode) *ListNode {
	visited := make(map[*ListNode]bool)
	cur := head
	for cur != nil {
		if visited[cur] {
			return cur // 第一个重复访问的 = 环入口
		}
		visited[cur] = true
		cur = cur.Next
	}
	return nil
}
