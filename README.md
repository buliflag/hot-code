# Hot-Code: 个性化算法训练营 🚀

`Hot-Code` 是一个基于 **Go 语言** 的 LeetCode 算法练习项目，它不仅仅是代码的堆砌，更是一个通过 **Gemini CLI** 自定义技能深度驱动的个性化学习模型。

## 🌟 核心理念

本项目通过自定义技能 `leetcode-coach`，将算法学习分为三个阶段：
1.  **演进式讲解**：从最直觉的暴力解法或基础 DP 数组，逐步演进到空间压缩的最终态，掌握算法的进化逻辑。
2.  **深度互动**：针对特定的理解卡点（如“虚拟边界”、“多重赋值”）进行专项答疑。
3.  **个性化模型**：自动维护 `LEETCODE_JOURNEY.md`，记录属于你自己的“悟道”时刻。

## 🛠 技术栈
- **编程语言**: Go (Golang)
- **核心工具**: [Gemini CLI](https://github.com/google/gemini-cli)
- **自定义技能**: `leetcode-coach`

## 📂 目录结构
- `/demo`: LeetCode 题目的 Go 语言实现。每份代码均包含 **基础版 (`xxxBasic`)** 和 **优化版**，展示演进过程。
- `/skills`: 存放 `leetcode-coach` 技能的核心定义与工作流。
- `LEETCODE_JOURNEY.md`: **核心学习日志**，记录每道题的特定卡点与关键转折点。
- `ALGO_CHEATSHEET.md`: **高浓度速查手册**，沉淀通用模板与 Go 实战技巧。
- `GEMINI.md`: 项目的指令上下文文件。

## 🚀 快速开始

### 1. 安装与配置
确保你已安装 `gemini-cli`。克隆本项目后，将 `leetcode-coach` 技能激活。

### 2. 使用 leetcode-coach 训练
在终端中启动 Gemini 并调用技能：
```bash
gemini "讲讲力扣 198 题"
```

### 3. 查看演进代码
每道题的代码中，你可以看到类似以下的结构：
```go
// robBasic: 基础 O(n) 方案
func robBasic(nums []int) int { ... }

// rob: 优化 O(1) 方案
func rob(nums []int) int { ... }
```

## 📈 学习心法
- **不只是 Copy 代码**：重点关注 `LEETCODE_JOURNEY.md` 中的“特定卡点”。
- **掌握演进路径**：明白为什么我们要从 `Basic` 版本进化到当前版本。
- **虚拟边界思维**：在处理 DP 和数组问题时，优先思考是否可以通过初始化“虚拟节点”来简化边界判断。

---
*由 Gemini CLI 辅助构建，开启你的个性化算法进阶之路。*
