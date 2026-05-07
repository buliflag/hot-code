package main

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	var ans [][]int
	var path []int
	//给候选集排序
	sort.Ints(candidates)

	var backtrack func(start, sum int)
	backtrack = func(start, sum int) {
		if sum == target {
			temp := make([]int, len(path))
			copy(temp, path)
			ans = append(ans, temp)
		}
		for i := start; i < len(candidates); i++ {
			if sum+candidates[i] > target {
				break
			}
			path = append(path, candidates[i])
			backtrack(i, sum+candidates[i])
			path = path[:len(path)-1]
		}
	}

	backtrack(0, 0)
	return ans
}
