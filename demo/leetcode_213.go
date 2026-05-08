package main

func rob_213(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	if n == 2 {
		return max(nums[0], nums[1])
	}
	res1 := robRange(nums, 0, n-2)
	res2 := robRange(nums, 1, n-1)
	return max(res1, res2)

}
func robRange(nums []int, start, end int) int {
	prev1, prev2 := 0, 0
	for i := start; i <= end; i++ {
		prev2, prev1 = prev1, max(prev1, prev2+nums[i])
	}
	return prev1
}
