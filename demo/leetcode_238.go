package main

// productExceptSelf 采用了空间优化的双向遍历方案（最终推荐版）
// 核心直觉：左右夹击。结果 = 左侧积 * 右侧积。
// 时间复杂度：O(n), 空间复杂度：O(1) (除输出数组外)
func productExceptSelf(nums []int) []int {
	n := len(nums)
	res := make([]int, n)

	// 1. 第一遍：计算每个元素左侧的乘积
	res[0] = 1
	for i := 1; i < n; i++ {
		res[i] = res[i-1] * nums[i-1]
	}

	// 2. 第二遍：利用动态变量 right 维护右侧乘积并实时计算结果
	right := 1
	for i := n - 1; i >= 0; i-- {
		res[i] = res[i] * right
		right = right * nums[i]
	}

	return res
}

// productExceptSelfBasic 展示了使用两个额外数组的方案（基础版）
// 核心：明确划分前缀积数组和后缀积数组。
// 时间复杂度：O(n), 空间复杂度：O(n)
func productExceptSelfBasic(nums []int) []int {
	n := len(nums)
	L, R := make([]int, n), make([]int, n)
	res := make([]int, n)

	L[0] = 1
	for i := 1; i < n; i++ {
		L[i] = L[i-1] * nums[i-1]
	}

	R[n-1] = 1
	for i := n - 2; i >= 0; i-- {
		R[i] = R[i+1] * nums[i+1]
	}

	for i := 0; i < n; i++ {
		res[i] = L[i] * R[i]
	}
	return res
}
