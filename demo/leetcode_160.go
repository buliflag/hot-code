package main

// ListNode 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// getIntersectionNode 双指针浪漫相遇法（最终推荐版）
// 核心直觉：我走过你走的路。pA 走完 A 跳去 B 的起点，pB 走完 B 跳去 A 的起点，
// 两者走的总路程相同（a+b+c），必在交汇点相遇。无交汇则同时到达 nil。
// 时间 O(m+n)，空间 O(1)
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pA, pB := headA, headB

	for pA != pB {
		if pA != nil {
			pA = pA.Next
		} else {
			pA = headB // 走完 A，跳去 B 从头走
		}

		if pB != nil {
			pB = pB.Next
		} else {
			pB = headA // 走完 B，跳去 A 从头走
		}
	}

	return pA // 要么是交点，要么是 nil（无交点）
}

// getIntersectionNodeBasic 哈希表法（基础版）
// 核心：把 A 的所有节点记到花名册，B 第一个在册的节点就是交点
// 时间 O(m+n)，空间 O(m)
func getIntersectionNodeBasic(headA, headB *ListNode) *ListNode {
	visited := make(map[*ListNode]bool)
	for cur := headA; cur != nil; cur = cur.Next {
		visited[cur] = true
	}
	for cur := headB; cur != nil; cur = cur.Next {
		if visited[cur] {
			return cur
		}
	}
	return nil
}
