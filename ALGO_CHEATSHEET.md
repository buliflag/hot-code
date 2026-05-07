# Go 语言算法面试通关 Cheatsheet (深度增强版) 🚀

这份文档旨在帮助你建立从“看到题”到“写出代码”的直觉链路。

---

## 1. 动态规划 (DP) - 坐标/网格类
**直觉模型**：**“木桶效应”** —— 当前格子的最优解受限于它最弱的邻居。
**代表题**：[221] 最大正方形

### 💡 解题思路
1. **定义状态**：`dp[i][j]` 表示以 `(i, j)` 为**右下角**能达到的目标属性（如最大边长）。
2. **转移逻辑**：观察当前点与左、上、左上邻居的几何关系。
3. **初始边界**：处理第一行和第一列，它们没有邻居，通常值就是原数组的值。

### ⚠️ 核心陷阱
- **类型转换**：Go 的 `math.Min` 只收 `float64`，处理 `int` 时建议手写 `min(a, b int) int` 以保证性能。
- **空间优化 (核心)**：对于只依赖前 1-2 个状态的线性 DP，务必使用**变量滚动**。
    - **技巧**：利用初始化为 0 的变量作为“虚拟边界”，简化边界判断逻辑。
    - **Go 语法**：利用 `a, b = b, new` 多重赋值避免使用临时变量。

### 🛠 代码模板 (线性 DP 空间压缩)
```go
func linearDP(nums []int) int {
    // prev2(i-2), prev1(i-1) 初始化为 0 作为虚拟边界
    prev2, prev1 := 0, 0
    for _, v := range nums {
        // 更新逻辑：计算当前状态并滚动
        prev2, prev1 = prev1, max(prev1, prev2+v)
    }
    return prev1
}
```

---
### 🛠 代码模板
```go
func maximalSquare(matrix [][]byte) int {
    m, n := len(matrix), len(matrix[0])
    dp := make([][]int, m)
    for i := range dp { dp[i] = make([]int, n) }
    
    maxSide := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == '1' {
                if i == 0 || j == 0 {
                    dp[i][j] = 1
                } else {
                    dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
                }
                if dp[i][j] > maxSide { maxSide = dp[i][j] }
            }
        }
    }
    return maxSide * maxSide
}
```

---

## 2. 快速选择 (Quick Select) - 寻找第 K 大
**直觉模型**：**“选秀淘汰赛”** —— 通过分区（Partition）不断缩小搜索范围。
**代表题**：[215] 数组中第 K 大元素

### 💡 解题思路
1. **索引转换**：第 `K` 大对应升序排序后的索引为 `target = len - K`。
2. **随机化**：必须随机选基准点并交换到末尾，防止处理有序数组时退化为 $O(N^2)$。
3. **局部搜索**：判断 `pivot` 位置与 `target` 的关系，只递归目标所在的半区。

### ⚠️ 核心陷阱
- **递归出口**：当 `pivot == target` 时必须立刻返回。
- **原地修改**：理解切片是引用传递，`Partition` 会直接修改原数组顺序。

### 🛠 代码模板
```go
func findKthLargest(nums []int, k int) int {
    rand.Seed(time.Now().UnixNano())
    target := len(nums) - k
    return quickSelect(nums, 0, len(nums)-1, target)
}

func quickSelect(nums []int, l, r, target int) int {
    p := partition(nums, l, r)
    if p == target { return nums[p] }
    if p < target { return quickSelect(nums, p+1, r, target) }
    return quickSelect(nums, l, p-1, target)
}

func partition(nums []int, l, r int) int {
    randIdx := l + rand.Intn(r-l+1)
    nums[randIdx], nums[r] = nums[r], nums[randIdx] // 临时泊车
    i := l
    for j := l; j < r; j++ {
        if nums[j] < nums[r] {
            nums[i], nums[j] = nums[j], nums[i]; i++
        }
    }
    nums[i], nums[r] = nums[r], nums[i]
    return i
}
```

---

## 3. 图论：拓扑排序 (Topological Sort)
**直觉模型**：**“拆锁游戏”** —— 先找没锁的门，进门后再拆掉后续门的锁。
**代表题**：[207/210] 课程表, [802] 安全状态

### 💡 解题思路
1. **建模**：构建邻接表 `adjacency`（记录我是谁的前置）和入度表 `indegrees`（记录我有几个前置）。
2. **启动**：将所有入度为 0 的节点放入队列。
3. **消耗**：从队列取出节点，将其指向的所有邻居的入度减 1。若邻居入度清零，则入队。

### ⚠️ 核心陷阱
- **有向环检测**：如果最后处理的节点数小于总节点数，说明图中存在环。
- **逆向思维**：若题目求“安全点”或“出度”，尝试**反转图**，把出度当入度处理。

### 🛠 代码模板 (BFS/Kahn 算法)
```go
func canFinish(numCourses int, prerequisites [][]int) bool {
    indegrees := make([]int, numCourses)
    adj := make([][]int, numCourses)
    for _, p := range prerequisites {
        indegrees[p[0]]++; adj[p[1]] = append(adj[p[1]], p[0])
    }
    queue := []int{}
    for i, in := range indegrees {
        if in == 0 { queue = append(queue, i) }
    }
    count := 0
    for len(queue) > 0 {
        u := queue[0]; queue = queue[1:]; count++
        for _, v := range adj[u] {
            indegrees[v]--
            if indegrees[v] == 0 { queue = append(queue, v) }
        }
    }
    return count == numCourses
}
```

---

## 4. 网格搜索 (DFS) - 岛屿专题
**直觉模型**：**“沉没大法”** —— 发现一点，沉没整座岛（抹除证据）。
**代表题**：[200/695/1254/827] 岛屿系列

### 💡 解题思路
1. **遍历**：双层 `for` 循环寻找起始陆地（'1'）。
2. **递归（DFS/BFS）**：从起始点出发，向四个方向扩散，将遇到的陆地全部置为水（'0'）。
3. **统计**：
    - 数岛屿：计数器 `+1`。
    - 算面积：DFS 返回 `1 + 四周面积`。
    - 封闭岛：先沉没边缘陆地，再数剩下的岛。

### ⚠️ 核心陷阱
- **边界检查**：递归的第一步必须检查 `r, c` 是否越界。
- **类型识别**：注意题目给的是 `byte` ('1') 还是 `int` (1)。

### 🛠 代码模板 (沉没大法)
```go
func numIslands(grid [][]byte) int {
    count := 0
    var dfs func(r, c int)
    dfs = func(r, c int) {
        if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] == '0' {
            return
        }
        grid[r][c] = '0' // 原地修改，省去 visited 空间
        dfs(r-1, c); dfs(r+1, c); dfs(r, c-1); dfs(r, c+1)
    }
    for i := range grid {
        for j := range grid[i] {
            if grid[i][j] == '1' {
                count++; dfs(i, j)
            }
        }
    }
    return count
}
```

---

## 5. 回溯算法 (Backtracking)
**直觉模型**：**“多层梦境与橡皮擦”** —— 深入探索，醒来后必须恢复现场。
**代表题**：[46] 全排列, [78] 子集

### 💡 解题思路
1. **决策树**：`for` 循环横向展开选择，`backtrack()` 纵向深入探索。
2. **选择与撤销**：在递归前后，必须对称地进行“添加路径”和“移除路径”的操作。
3. **去重策略**：
    - 全排列（顺序敏感）：使用 `used` 数组。
    - 组合/子集（顺序无关）：使用 `start` 索引限制搜索范围。

### ⚠️ 核心陷阱
- **切片深拷贝 (Go 必考)**：`res = append(res, path)` 时，必须先 `copy` 一份 `path` 的副本，否则后续修改会破坏结果集。

### 🛠 代码模板 (通用回溯)
```go
func permute(nums []int) [][]int {
    var res [][]int
    path := []int{}
    used := make([]bool, len(nums))
    var backtrack func()
    backtrack = func() {
        if len(path) == len(nums) {
            temp := make([]int, len(path)); copy(temp, path)
            res = append(res, temp); return
        }
        for i := 0; i < len(nums); i++ {
            if used[i] { continue }
            path = append(path, nums[i]); used[i] = true // 做选择
            backtrack()
            path = path[:len(path)-1]; used[i] = false    // 撤销选择（恢复现场）
        }
    }
    backtrack(); return res
}
```

---

## 6. 前缀树 (Trie)
**直觉模型**：**“查字典”** —— 共享前缀的词走同一条路。
**代表题**：[208] 实现 Trie

### 💡 解题思路
1. **节点结构**：`children` 指向下一级，`isEnd` 标记此处是否构成完整单词。
2. **空间权衡**：
    - 字符集小（a-z）：用 `[26]*Trie` 数组，速度最快。
    - 字符集大（中文/Emoji）：用 `map[rune]*Trie`，更灵活。

### 🛠 代码模板
```go
type Trie struct {
    children [26]*Trie
    isEnd    bool
}

func (this *Trie) Insert(word string) {
    node := this
    for _, ch := range word {
        idx := ch - 'a'
        if node.children[idx] == nil { node.children[idx] = &Trie{} }
        node = node.children[idx]
    }
    node.isEnd = true
}

func (this *Trie) Search(word string) bool {
    node := this
    for _, ch := range word {
        idx := ch - 'a'
        if node.children[idx] == nil { return false }
        node = node.children[idx]
    }
    return node.isEnd
}
```

---
*保持更新，这是你的面试必杀技！*
