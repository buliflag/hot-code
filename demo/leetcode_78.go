package main

func subsets(nums []int) [][]int {
	var res [][]int
	path := []int{}
	var backtrack func(start int)
	backtrack = func(start int) {
		//直接记录
		temp := make([]int, len(path))
		copy(temp, path)
		res = append(res, temp)

		//遍历
		for i := start; i < len(nums); i++ {
			//[做选择]
			path = append(path, nums[i])

			//递归到下一层
			backtrack(i + 1)

			//撤销
			path = path[:len(path)-1]
		}
	}
	backtrack(0)
	return res
}
