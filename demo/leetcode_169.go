package main

// majorityElement 采用了摩尔投票法（最终推荐版）
// 核心直觉：对拼消耗。由于多数元素个数 > n/2，它在消耗战后一定能活下来。
// 时间复杂度：O(n), 空间复杂度：O(1)
func majorityElement(nums []int) int {
	// candidate: 当前霸主, count: 霸主的血量
	candidate, count := 0, 0

	for _, num := range nums {
		if count == 0 {
			// 旧霸主血量耗尽，新霸主登基
			candidate = num
			count = 1
		} else if num == candidate {
			// 遇到同党，血量增加
			count++
		} else {
			// 遇到异己，同归于尽（消耗）
			count--
		}
	}

	return candidate
}

// majorityElementBasic 展示了哈希表计数的通用方案（基础版）
// 核心：利用空间换时间。
// 时间复杂度：O(n), 空间复杂度：O(n)
func majorityElementBasic(nums []int) int {
	counts := make(map[int]int)
	threshold := len(nums) / 2
	for _, num := range nums {
		counts[num]++
		if counts[num] > threshold {
			return num
		}
	}
	return -1
}
