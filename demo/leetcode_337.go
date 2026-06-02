package main

// TreeNode 是二叉树节点的定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// rob 采用了树形 DP 方案（最终推荐版）
// 核心：每个节点向父节点汇报两个状态 [不偷当前节点的最大值, 偷当前节点的最大值]
func rob_337(root *TreeNode) int {
	res := dfs(root)
	return max(res[0], res[1])
}

func dfs(node *TreeNode) []int {
	if node == nil {
		// 虚拟边界：空节点无论偷不偷都是 0 元
		return []int{0, 0}
	}

	// 演进式思考：从子节点获取“状态报告”
	left := dfs(node.Left)
	right := dfs(node.Right)

	// 状态 0：不偷当前节点
	// 此时子节点偷不偷随它，我们取它两种方案中的最大值进行累加
	doNotRob := max(left[0], left[1]) + max(right[0], right[1])

	// 状态 1：偷当前节点
	// 核心约束：子节点绝对不能偷，必须强制取子节点的“不偷状态” (索引 0)
	doRob := node.Val + left[0] + right[0]

	return []int{doNotRob, doRob}
}

// robBasic 展示了带记忆化搜索的暴力递归方案（基础版）
// 适合理解“爷爷-爸爸-孙子”之间的逻辑依赖
func robBasic_337(root *TreeNode) int {
	memo := make(map[*TreeNode]int)
	var helper func(node *TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		if v, ok := memo[node]; ok {
			return v
		}

		// 1. 偷当前节点：则只能加孙子辈的钱
		robIt := node.Val
		if node.Left != nil {
			robIt += helper(node.Left.Left) + helper(node.Left.Right)
		}
		if node.Right != nil {
			robIt += helper(node.Right.Left) + helper(node.Right.Right)
		}

		// 2. 不偷当前节点：直接加两个儿子的钱（儿子偷不偷由儿子自己决策）
		notRobIt := helper(node.Left) + helper(node.Right)

		res := max(robIt, notRobIt)
		memo[node] = res
		return res
	}
	return helper(root)
}
