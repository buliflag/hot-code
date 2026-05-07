package main

// rob 采用了 O(1) 空间复杂度的动态规划演进方案（最终推荐版）
func rob(nums []int) int {
	// 演进思路：由于每次决策只依赖前两步，可以用两个变量代替数组。
	// 初始化 prev2, prev1 为 0（虚拟边界），使得循环逻辑统一。
	prev2, prev1 := 0, 0
	for _, v := range nums {
		// Go 多重赋值：计算完毕后统一更新状态
		prev2, prev1 = prev1, max(prev1, prev2+v)
	}
	return prev1
}

// robBasic 展示了最直觉的 O(n) 空间复杂度方案（基础版）
// 适合理解动态规划的原始状态转移逻辑
func robBasic(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}

	// 定义 dp 数组：记录每一个状态
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	for i := 2; i < n; i++ {
		// 核心决策：不偷当前家 vs 偷当前家+前前家
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[n-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
