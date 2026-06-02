package main

// rob 采用了拆解方案（最终推荐版）
// 核心直觉：破环成链。
func rob(nums []int) int {
	n := len(nums)
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}

	// 拆解为两个线性问题：1. 不偷第一间房； 2. 不偷最后一间房
	return max(robRange(nums, 1, n-1), robRange(nums, 0, n-2))
}

func robRange(nums []int, start, end int) int {
	prev2, prev1 := 0, 0
	for i := start; i <= end; i++ {
		prev2, prev1 = prev1, max(prev1, prev2+nums[i])
	}
	return prev1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
