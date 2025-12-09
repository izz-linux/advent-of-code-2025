package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Open the input file
	file, err := os.Open("./input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Initialize variables
	dial := 50
	countZero := 0
	const maxDial = 100

	// Read the file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 2 {
			continue
		}

		// Parse the direction and distance
		direction := line[0]
		distance, err := strconv.Atoi(line[1:])
		if err != nil {
			fmt.Println("Error parsing distance:", err)
			return
		}

		// Simulate each click and count zero crossings
		for i := 0; i < distance; i++ {
			if direction == 'L' {
				dial = (dial - 1 + maxDial) % maxDial
			} else if direction == 'R' {
				dial = (dial + 1) % maxDial
			}
			if dial == 0 {
				countZero++
			}
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Output the result
	fmt.Println("Password:", countZero)
}
