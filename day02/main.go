package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read input file
	data, err := os.ReadFile("./input")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse ranges
	input := strings.TrimSpace(string(data))
	ranges := strings.Split(input, ",")

	totalSum := 0

	for _, rangeStr := range ranges {
		rangeStr = strings.TrimSpace(rangeStr)
		parts := strings.Split(rangeStr, "-")
		if len(parts) != 2 {
			continue
		}

		start, err1 := strconv.Atoi(parts[0])
		end, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			continue
		}

		// Check each number in range
		for num := start; num <= end; num++ {
			if isInvalidID(num) {
				totalSum += num
			}
		}
	}

	fmt.Println("Sum of invalid IDs:", totalSum)
}

func isInvalidID(num int) bool {
	str := strconv.Itoa(num)
	length := len(str)

	// Check for leading zeros (invalid)
	if str[0] == '0' {
		return false
	}

	// Try all possible pattern lengths (from 1 to length/2)
	for patternLen := 1; patternLen <= length/2; patternLen++ {
		// Check if the string length is divisible by pattern length
		if length%patternLen != 0 {
			continue
		}

		// Extract the pattern
		pattern := str[:patternLen]

		// Check if leading zeros in pattern
		if pattern[0] == '0' {
			continue
		}

		// Check if entire string is this pattern repeated
		repeats := length / patternLen
		if repeats >= 2 && isRepeatedPattern(str, pattern, repeats) {
			return true
		}
	}

	return false
}

func isRepeatedPattern(str, pattern string, repeats int) bool {
	for i := 0; i < repeats; i++ {
		start := i * len(pattern)
		end := start + len(pattern)
		if str[start:end] != pattern {
			return false
		}
	}
	return true
}
