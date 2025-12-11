package main

import (
	"bufio"
	"fmt"
	"os"
)

func countAdjacentRolls(grid [][]rune, row, col int) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	// Check all 8 adjacent positions
	directions := [][]int{
		{-1, -1}, {-1, 0}, {-1, 1}, // top-left, top, top-right
		{0, -1}, {0, 1}, // left, right
		{1, -1}, {1, 0}, {1, 1}, // bottom-left, bottom, bottom-right
	}

	for _, dir := range directions {
		newRow := row + dir[0]
		newCol := col + dir[1]

		// Check bounds
		if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
			if grid[newRow][newCol] == '@' {
				count++
			}
		}
	}

	return count
}

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

func part1(grid [][]rune) int {
	accessibleCount := 0

	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == '@' {
				adjacentCount := countAdjacentRolls(grid, row, col)
				if adjacentCount < 4 {
					accessibleCount++
				}
			}
		}
	}

	return accessibleCount
}

func part2(grid [][]rune) int {
	// Create a copy of the grid to modify
	workGrid := make([][]rune, len(grid))
	for i := range grid {
		workGrid[i] = make([]rune, len(grid[i]))
		copy(workGrid[i], grid[i])
	}

	totalRemoved := 0

	// Keep removing accessible rolls until none are left
	for {
		accessible := findAccessibleRolls(workGrid)

		if len(accessible) == 0 {
			break
		}

		// Remove all accessible rolls in this iteration
		for _, pos := range accessible {
			row, col := pos[0], pos[1]
			workGrid[row][col] = '.'
			totalRemoved++
		}
	}

	return totalRemoved
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Read the grid
	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			grid = append(grid, []rune(line))
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Printf("Part 1: %d rolls of paper can be accessed by a forklift\n", part1(grid))
	fmt.Printf("Part 2: %d total rolls of paper can be removed\n", part2(grid))
}
