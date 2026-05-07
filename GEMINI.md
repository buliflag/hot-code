# Gemini CLI 项目上下文: hot-code

这是一个名为 **hot-code** 的个性化算法训练项目，旨在通过 Gemini CLI 的 `leetcode-coach` 技能，深入学习和记录 LeetCode 算法题的解题思路、个性化卡点及心得。

## 项目概览
- **核心目标**：不仅是刷题，而是通过“形象类比”和“深度互动”建立直觉，并形成个性化的算法模型。
- **主要技术栈**：Go 语言。
- **特色功能**：集成 `leetcode-coach` 技能，自动化维护学习日志 `LEETCODE_JOURNEY.md`。

## 目录结构
- `demo/`: 存放 LeetCode 题目的 Go 语言实现代码。
- `skills/leetcode-coach/`: `leetcode-coach` 技能的定义文件、模板和参考资料。
- `LEETCODE_JOURNEY.md`: 核心文档，记录每道题的解题思路、我的特定卡点和关键转折点。
- `ALGO_CHEATSHEET.md`: 算法套路速查表，包含常用策略和 Go 语言特性。
- `leetcode-coach.skill`: 技能快捷方式或定义。

## 关键流程与约定

### 1. 刷题与记录流程
使用 `leetcode-coach` 技能进行互动：
1. **输入题目**：向 Gemini 提供 LeetCode 题目或编号。
2. **深度讨论**：针对实现细节、边界条件或不理解的地方进行追问。
3. **模型更新**：当理解透彻后，要求 Gemini 更新 `LEETCODE_JOURNEY.md`。

### 2. 代码约定
- 存放在 `demo/` 目录下，文件命名格式为 `leetcode_{编号}.go`。
- 使用 `package main`，但通常作为算法片段，不包含 `main` 函数，除非用于本地测试。
- 追求 idiomatic Go (地道的 Go 代码)，例如使用切片原地修改、匿名 DFS 函数等。

### 3. 文档维护
- **LEETCODE_JOURNEY.md**：应包含“核心思想”、“解法套路”、“我的特定卡点”和“关键转折点”。这是最具有个人价值的部分。
- **ALGO_CHEATSHEET.md**：用于沉淀通用的模板和 Go 语言踩坑指南。

## 常用命令
- **运行测试**：本项目暂无统一测试框架，如需本地测试单个文件，可手动添加 `main` 函数后运行：
  ```bash
  go run demo/leetcode_xxx.go
  ```
- **查看技能详情**：
  ```bash
  gemini skill leetcode-coach --show
  ```

---
*此文件由 Gemini CLI 自动生成，用于提供项目上下文。*
