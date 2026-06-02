package main

// maxProduct 采用了正负双轨 DP 的空间优化版（最终推荐版）
// 核心直觉：负负得正。维护两条轨道——当前结尾的最大值和最小值，
// 当遇到负数时，之前"最烂"的 minF 乘以负数可能直接翻身成为 maxF。
// 时间复杂度：O(n), 空间复杂度：O(1)
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 两条轨道同时跑
	maxF, minF, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		// 必须先保存上一轮的值，因为计算 minF 时也需要上一轮的 maxF/minF
		preMaxF, preMinF := maxF, minF

		// 三种选择：重新出发 | 接上之前的最大值 | 接上之前的最小值（可能翻身）
		maxF = max(nums[i], max(preMaxF*nums[i], preMinF*nums[i]))
		minF = min(nums[i], min(preMaxF*nums[i], preMinF*nums[i]))
		res = max(res, maxF)
	}
	return res
}

// maxProductBasic 展示了使用两个 DP 数组的基础方案
// 核心：maxDp[i] 和 minDp[i] 分别记录以 i 结尾的最大/最小乘积。
// 时间复杂度：O(n), 空间复杂度：O(n)
func maxProductBasic(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	maxDp := make([]int, n)
	minDp := make([]int, n)
	maxDp[0], minDp[0] = nums[0], nums[0]
	res := nums[0]

	for i := 1; i < n; i++ {
		// 演进逻辑与优化版完全一致，只是用数组代替了滚动变量
		maxDp[i] = max(nums[i], max(maxDp[i-1]*nums[i], minDp[i-1]*nums[i]))
		minDp[i] = min(nums[i], min(maxDp[i-1]*nums[i], minDp[i-1]*nums[i]))
		res = max(res, maxDp[i])
	}
	return res
}
