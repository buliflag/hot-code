package main

// ListNode 链表节点定义
type ListNode struct {
	Val  int
	Next *ListNode
}

// mergeTwoLists 哨兵节点版（最终推荐）
// 核心直觉：两串排好序的珠子，每次挑小的那颗穿进针里
// 时间 O(m+n)，空间 O(1)
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cur.Next = list1
			list1 = list1.Next
		} else {
			cur.Next = list2
			list2 = list2.Next
		}
		cur = cur.Next
	}
	if list1 != nil {
		cur.Next = list1
	} else if list2 != nil {
		cur.Next = list2
	}
	return dummy.Next
}

// mergeTwoListsRecursive 递归版
// 核心直觉：挑小的当头，剩下的交给递归去合并
// 时间 O(m+n)，空间 O(m+n)
func mergeTwoListsRecursive(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoListsRecursive(list1.Next, list2)
		return list1
	}
	list2.Next = mergeTwoListsRecursive(list1, list2.Next)
	return list2
}
