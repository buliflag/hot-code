# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Hot-Code is a Go-based LeetCode algorithm practice repository paired with a Gemini CLI custom skill (`leetcode-coach`) that provides interactive, analogy-driven algorithm coaching. The project emphasizes understanding the **evolution** from naive to optimized solutions, not just memorizing final answers.

## Build & Run

There is no build system or test framework. Code lives as algorithm fragments in `demo/`:

```bash
# Run a single file (only works if it has a main() function)
go run demo/leetcode_215.go

# Some files lack main() — add a temporary one for local testing
```

All files use `package main`. Go version is whatever `go` is on PATH.

## Code Architecture

### Directory Structure
- `demo/` — LeetCode solutions, one file per problem: `leetcode_{number}.go`
- `skills/leetcode-coach/` — Gemini CLI skill definition (SKILL.md), templates, and reference material
- `leetcode-coach.skill` — packaged skill archive (zip format)

### Code Conventions

**Dual-version pattern**: Many solutions include both a naive and optimized implementation side by side:
- `xxxBasic()` — the intuitive, often O(n)-space version
- `xxx()` — the optimized final version (O(1) space, in-place, etc.)

This lets readers compare the evolution path directly in one file.

**Recurring Go idioms across the codebase**:
- **Anonymous recursive functions**: Declare `var dfs func(...)` then assign the closure — required because Go doesn't allow recursive lambdas inline
- **Hand-written `min`/`max`**: Go's `math.Min`/`math.Max` only accept `float64`, so all files define local `min`/`max` helpers for `int`
- **Multiple assignment for state rolling**: `a, b = b, new` pattern in DP space compression (see leetcode_198.go)
- **Slice deep copy in backtracking**: Always `copy(temp, path)` before `append(res, temp)` — otherwise all result entries share the same backing array
- **In-place mutation as visited tracking**: DFS grid problems mutate `grid[r][c] = '0'` instead of allocating a `visited` matrix
- **Virtual boundary initialization**: Initialize DP variables to `0` to represent "empty house before the array starts," eliminating edge-case branches

### Key Documentation Files
- **LEETCODE_JOURNEY.md** — Personal learning log. Each entry captures: core idea, solution pattern, **specific sticking points** (the user's exact confusions), **key breakthrough moments** (the analogy that made it click), and Go implementation notes. This is the most personally valuable file.
- **ALGO_CHEATSHEET.md** — Dense reference of algorithm templates and Go gotchas organized by problem category (DP, Quick Select, Topological Sort, DFS/Islands, Backtracking, Trie). Updated when new reusable patterns emerge.
- **PROGRESS.md** — Visual progress dashboard tracking Hot 100 completion status. Contains a categorized checklist of all 100 problems with checkboxes, category-level progress bars, and an extended practice section for non-Hot-100 problems. Must be synced after each completed problem.
- **GEMINI.md** — Project context file for Gemini CLI, describing the repo structure and conventions.

### The leetcode-coach Skill

Defined in `skills/leetcode-coach/SKILL.md`. Its workflow has three phases:
1. **Evolutionary explanation** — present the naive approach first, then show optimizations with clear reasoning
2. **Interactive Q&A** — use analogies and manual tracing to resolve specific confusions before moving on
3. **Personalized recording** — only after the user says they understand, update `LEETCODE_JOURNEY.md` with their specific sticking points and breakthroughs, and update `ALGO_CHEATSHEET.md` if a new reusable pattern emerged
4. **Progress sync** — after recording, mark the problem as `[x]` in `PROGRESS.md`, update the stats and progress bars

The skill is strict about: never summarizing while the user still has doubts, and never omitting personal details from the record.
