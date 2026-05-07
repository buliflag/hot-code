# 算法套路库 (Go 语言实现)

## 1. 滑动窗口 (Sliding Window)
**适用场景**：连续子数组/子串问题。
**Go 模板**：
```go
func slidingWindow(s string) {
    window := make(map[byte]int)
    left, right := 0, 0
    for right < len(s) {
        c := s[right]
        right++
        // 进行窗口内数据的一系列更新
        ...
        // 判断左侧窗口是否要收缩
        for windowNeedsShrink {
            d := s[left]
            left++
            // 进行窗口内数据的一系列更新
            ...
        }
    }
}
```

## 2. 二分查找 (Binary Search)
**核心要点**：确定搜索区间（左闭右闭 `[left, right]` 或左闭右开 `[left, right)`）。
**Go 实现**：
```go
func search(nums []int, target int) int {
    left, right := 0, len(nums)-1
    for left <= right {
        mid := left + (right-left)/2
        if nums[mid] == target {
            return mid
        } else if nums[mid] < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return -1
}
```

## 3. 回溯算法 (Backtracking)
**适用场景**：全排列、组合、子集等 DFS 问题。
**Go 实现技巧**：注意 slice 的深拷贝问题。
```go
func backtrack(path []int, choices []int) {
    if isTarget(path) {
        // 关键：必须拷贝 path，因为 path 是共用的
        temp := make([]int, len(path))
        copy(temp, path)
        res = append(res, temp)
        return
    }
    for i := range choices {
        path = append(path, choices[i])
        backtrack(path, choices)
        path = path[:len(path)-1] // 回溯
    }
}
```
