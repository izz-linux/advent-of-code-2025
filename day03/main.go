package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
)

func maxJoltageFromBank(bank string, numBatteries int) *big.Int {
	n := len(bank)
	// We need to skip (n - numBatteries) batteries
	toSkip := n - numBatteries

	// Greedy approach: skip the smallest digits from left to right
	// while considering their position value

	result := make([]byte, 0, numBatteries)
	skipped := 0
	i := 0

	for len(result) < numBatteries {
		// How many more digits do we need?
		remaining := numBatteries - len(result)
		// How many digits are left to choose from?
		available := n - i
		// How many more can we skip?
		canSkip := available - remaining

		if skipped < toSkip && canSkip > 0 {
			// Look ahead to see if we should skip this digit
			// Find the best digit we can take in the next 'canSkip+1' positions
			bestDigit := byte('0')
			bestPos := i

			for j := i; j <= i+canSkip && j < n; j++ {
				if bank[j] > bestDigit {
					bestDigit = bank[j]
					bestPos = j
				}
			}

			// Skip everything before the best digit
			skipsNeeded := bestPos - i
			skipped += skipsNeeded
			i = bestPos
		}

		// Take this digit
		result = append(result, bank[i])
		i++
	}

	// Convert result to big.Int
	resultStr := string(result)
	joltage := new(big.Int)
	joltage.SetString(resultStr, 10)

	return joltage
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalJoltage := big.NewInt(0)

	for scanner.Scan() {
		bank := scanner.Text()
		if bank != "" {
			maxJoltage := maxJoltageFromBank(bank, 12)
			fmt.Printf("Bank %s: max joltage = %s\n", bank, maxJoltage.String())
			totalJoltage.Add(totalJoltage, maxJoltage)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Printf("\nTotal output joltage: %s\n", totalJoltage.String())
}
