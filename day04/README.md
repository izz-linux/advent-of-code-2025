# Day 4

## Problem Overview

--- Day 4: Printing Department --- You ride the escalator down to the printing department. They're clearly getting ready for Christmas; they have lots of large rolls of paper everywhere, and there's even a massive printer in the corner (to handle the really big print jobs). Decorating here will be easy: they can make their own decorations.

[View full problem on Advent of Code](https://adventofcode.com/2025/day/4)

## Solution Approach

The solution employs iteration over data structures to efficiently solve the problem.

## Key Code Snippets

```go
func findAccessibleRolls(grid [][]rune) [][]int {
	var accessible [][]int

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == '@' {
				adjacentCount := countAdjacentRolls(grid, row, col)
				if adjacentCount < 4 {
					accessible = append(accessible, []int{row, col})
				}
			}
		}
	}

	return accessible
}
```

## Input Format

Sample input (first few lines):

```
...@@...@..@...@@...@@.@@@@@.@.@..@@@.@@@@@@@@@..@.@@.@@@.@...@.@.@@@@@.@@@.@@@@.@.@@@@@@@@@@@@@@.@@.@@@@@@@@@@@@@.@.@.@@@@@@@..@.@..@@@@@
@@@@..@@@@@@@@@..@..@@@@@..@@@.@.@@@@@..@@..@.@...@.@@@...@@@@....@@@@@@@@@.@@@@@@.@.@@@.@@@@.@@@.@@@@.@@@@.@...@@.@..@@@@@@@..@...@.@@.@@
.@@..@@@@.....@@@@@@@.@@.@@@@.@.@@..@@..@@...@@.@.@@@@@.@.@@@@@@@.@.@@.@@@@@@@@@@..@...@@..@.@@@@...@@@.@@.@@@.@..@@@@@.@@@..@@@@@.@@..@@.
..@@@@@@.@.@@@@@@....@..@@@@.@@.@@@@.@.@.@.@@.@.@@@...@.@@@@@@@@@.@.@@@@@.@@@@.@@@@@.@@@..@.@@..@@@@@@@..@.@@@....@..@..@@@.@@@..@@@.@@@@.
@@@.@@@@@.@.@.@.@.@.@@@@@.@..@.@@@.@@@....@@@@@....@@@@.@@@@.@@..@@.@@..@..@@.@....@@..@@@..@.@@..@@@@@@@@@@.@.@@@..@@.@.@@@@...@.@@.@@@@.
.@@.@.@@@@.....@@@@@@@.@@@@@@@@@@..@@@@@@@@.@.@@....@@.@..@@@@@@@@...@@..@..@@..@@@@.@@@.@@@..@@.@@.@..@@.@.@@@@...@.@.@@.@.@@@@@.@@@.@..@
.@..@@@...@@@@@@@...@@@@@.@@@....@.@@@@@...@..@@@@.@@..@@@@@@@.@@@@@@@@@..@@.@.@@@@.@@.@.@..@.@...@@@.@@@@@....@@.@...@@@..@.@@.@..@@@@@@@
```

## Running the Solution

```bash
cd day04
go run main.go
```

## Result

```
Run the solution to see the answer
```

---

*Last updated: 2025-12-11 06:29:43*
