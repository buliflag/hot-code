package main

func permute(nums []int) [][]int {
	var res [][]int
	path := []int{}
	used := make([]bool, len(nums))
	var backtrack func()
	backtrack = func() {
		//达到条件了
		if len(path) == len(nums) {
			temp := make([]int, len(nums))
			copy(temp, path)
			res = append(res, temp)
			return
		}
		//遍历
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			//[做选择]
			path = append(path, nums[i])
			used[i] = true

			//递归到下一层
			backtrack()

			//撤销
			path = path[:len(path)-1]
			used[i] = false
		}
	}
	backtrack()
	return res
}
