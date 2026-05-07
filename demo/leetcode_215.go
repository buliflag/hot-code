package main

import "math/rand"

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	return quickselect(nums, 0, n-1, n-k)
}
func quickselect(nums []int, l, r, k int) int {
	//进行分区
	p := partition(nums, l, r)
	if p == k {
		return nums[p]
	} else if p > k {
		return quickselect(nums, l, p-1, k)
	} else {
		return quickselect(nums, p+1, r, k)
	}

}
func partition(nums []int, left, right int) int {
	//随机选一个，作为标准
	randomIndex := left + rand.Intn(right-left+1)
	//把这个值固定到最右侧
	nums[right], nums[randomIndex] = nums[randomIndex], nums[right]
	pivot := nums[right]
	i := left
	//进行比较
	for j := left; j < right; j++ {
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[right] = nums[right], nums[i]
	return i
}

func main() {
	nums := []int{2, 9, 1, 2, 3, 4, 5, 6, 7}
	k := 3
	print(findKthLargest(nums, k))
}
