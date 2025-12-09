package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

func main() {
	root := "."
	days := findDays(root)
	updateMainReadme(root, days)
	for _, day := range days {
		updateDayReadme(root, day)
	}
	fmt.Println("README files updated.")
}

func findDays(root string) []string {
	entries, _ := os.ReadDir(root)
	var days []string
	for _, e := range entries {
		if e.IsDir() && strings.HasPrefix(e.Name(), "day") {
			days = append(days, e.Name())
		}
	}
	return days
}

func getDayNumber(day string) string {
	if day < "day10" {
		return strings.TrimPrefix(day, "day0")
	}
	return strings.TrimPrefix(day, "day")
}

func fetchProblemDescription(dayNum string) string {
	url := fmt.Sprintf("https://adventofcode.com/2025/day/%s", dayNum)
	resp, err := http.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	content := string(body)
	start := strings.Index(content, "<article class=\"day-desc\">")
	end := strings.Index(content, "</article>")

	if start == -1 || end == -1 {
		return ""
	}

	return content[start+26 : end]
}

func stripHTMLTags(html string) string {
	// Remove HTML tags but keep the text content
	re := regexp.MustCompile(`<[^>]*>`)
	text := re.ReplaceAllString(html, " ")

	// Clean up entities
	text = strings.ReplaceAll(text, "&quot;", "\"")
	text = strings.ReplaceAll(text, "&amp;", "&")
	text = strings.ReplaceAll(text, "&lt;", "<")
	text = strings.ReplaceAll(text, "&gt;", ">")
	text = strings.ReplaceAll(text, "&nbsp;", " ")

	// Clean up whitespace
	text = regexp.MustCompile(`\s+`).ReplaceAllString(text, " ")

	return strings.TrimSpace(text)
}

func summarizeProblem(html string) string {
	if html == "" {
		return "Problem description not available. See https://adventofcode.com/2025 for details."
	}

	// Extract just the text content
	text := stripHTMLTags(html)

	// Split into sentences
	sentences := regexp.MustCompile(`[.!?]+\s+`).Split(text, -1)

	// Take first 3-5 sentences that are substantial (more than 20 chars)
	var summary []string
	for _, sentence := range sentences {
		sentence = strings.TrimSpace(sentence)
		if len(sentence) > 20 && len(summary) < 4 {
			summary = append(summary, sentence)
		}
		if len(summary) >= 3 && len(strings.Join(summary, " ")) > 150 {
			break
		}
	}

	if len(summary) == 0 {
		return "Solve the daily Advent of Code puzzle."
	}

	result := strings.Join(summary, ". ")
	if !strings.HasSuffix(result, ".") && !strings.HasSuffix(result, "!") && !strings.HasSuffix(result, "?") {
		result += "."
	}

	return result
}

func extractKeySnippets(mainPath string) string {
	data, err := os.ReadFile(mainPath)
	if err != nil {
		return "// Code not available"
	}

	content := string(data)
	lines := strings.Split(content, "\n")

	var snippets []string
	var currentSnippet []string
	inInterestingFunc := false
	braceCount := 0

	// Patterns to identify interesting functions (not boilerplate)
	boilerplatePatterns := []string{
		"func main()",
		"func readFile",
		"func readLines",
		"func readInput",
	}

	interestingKeywords := []string{
		"solve", "calculate", "process", "parse",
		"find", "count", "simulate", "compute",
		"search", "build", "generate",
	}

	for i, line := range lines {
		trimmed := strings.TrimSpace(line)

		// Check if we're entering an interesting function
		if strings.HasPrefix(trimmed, "func ") && !inInterestingFunc {
			isBoilerplate := false
			for _, pattern := range boilerplatePatterns {
				if strings.Contains(line, pattern) {
					isBoilerplate = true
					break
				}
			}

			isInteresting := false
			lowerLine := strings.ToLower(line)
			for _, keyword := range interestingKeywords {
				if strings.Contains(lowerLine, keyword) {
					isInteresting = true
					break
				}
			}

			// Start capturing if interesting and not boilerplate
			if !isBoilerplate && isInteresting {
				inInterestingFunc = true
				currentSnippet = []string{line}
				braceCount = strings.Count(line, "{") - strings.Count(line, "}")
				continue
			}
		}

		if inInterestingFunc {
			currentSnippet = append(currentSnippet, line)
			braceCount += strings.Count(line, "{") - strings.Count(line, "}")

			// End of function
			if braceCount == 0 && len(currentSnippet) > 1 {
				// Only include if it's between 5-25 lines
				if len(currentSnippet) >= 5 && len(currentSnippet) <= 25 {
					snippetText := strings.Join(currentSnippet, "\n")
					// Skip if it's mostly error handling or I/O
					errorCount := strings.Count(snippetText, "if err != nil")
					ioCount := strings.Count(snippetText, "fmt.") +
						strings.Count(snippetText, "os.") +
						strings.Count(snippetText, "bufio.")

					if errorCount <= 2 && ioCount < len(currentSnippet)/4 {
						snippets = append(snippets, snippetText)
					}
				}
				inInterestingFunc = false
				currentSnippet = nil
			}
		}

		// Look for interesting algorithm sections marked with comments
		if !inInterestingFunc {
			commentMarkers := []string{
				"// Core algorithm", "// Main logic", "// Key insight",
				"// Algorithm:", "// Solution:", "// Strategy:",
			}
			for _, marker := range commentMarkers {
				if strings.Contains(trimmed, marker) {
					end := i + 12
					if end > len(lines) {
						end = len(lines)
					}
					snippet := strings.Join(lines[i:end], "\n")
					snippets = append(snippets, snippet)
					break
				}
			}
		}
	}

	if len(snippets) == 0 {
		// Fallback: extract any non-main, non-read function
		return extractFallbackSnippet(lines)
	}

	// Limit to 2 most relevant snippets
	if len(snippets) > 2 {
		snippets = snippets[:2]
	}

	result := ""
	for i, snippet := range snippets {
		if i == 0 {
			result = "```go\n" + snippet + "\n```"
		} else {
			result += "\n\n```go\n" + snippet + "\n```"
		}
	}

	return result
}

func extractFallbackSnippet(lines []string) string {
	// Find the first substantial function that isn't main or read*
	var snippet []string
	inFunc := false
	braceCount := 0

	for _, line := range lines {
		trimmed := strings.TrimSpace(line)

		if strings.HasPrefix(trimmed, "func ") && !inFunc {
			if !strings.Contains(line, "main()") &&
				!strings.Contains(line, "read") &&
				!strings.Contains(line, "Read") {
				inFunc = true
				snippet = []string{line}
				braceCount = strings.Count(line, "{") - strings.Count(line, "}")
				continue
			}
		}

		if inFunc {
			snippet = append(snippet, line)
			braceCount += strings.Count(line, "{") - strings.Count(line, "}")

			if braceCount == 0 && len(snippet) > 3 {
				if len(snippet) <= 30 {
					return "```go\n" + strings.Join(snippet, "\n") + "\n```"
				}
				break
			}
		}
	}

	return "```go\n// See main.go for implementation\n```"
}

func getSolutionApproach(dayPath string) string {
	mainPath := filepath.Join(dayPath, "main.go")
	data, err := os.ReadFile(mainPath)
	if err != nil {
		return "Implementation uses standard algorithms and data structures."
	}

	content := string(data)

	// Look for approach indicators in the code
	var approaches []string

	if strings.Contains(content, "map[") {
		approaches = append(approaches, "hash maps for efficient lookups")
	}
	if strings.Contains(content, "sort.") {
		approaches = append(approaches, "sorting algorithms")
	}
	if strings.Contains(content, "strings.Split") && strings.Contains(content, "strconv") {
		approaches = append(approaches, "input parsing and conversion")
	}
	if regexp.MustCompile(`for.*range`).MatchString(content) {
		approaches = append(approaches, "iteration over data structures")
	}
	if strings.Contains(content, "recursive") ||
		(strings.Contains(content, "func ") && regexp.MustCompile(`func \w+\([^)]*\).*\{\s*return \w+\(`).MatchString(content)) {
		approaches = append(approaches, "recursive problem solving")
	}
	if strings.Contains(content, "queue") || strings.Contains(content, "stack") {
		approaches = append(approaches, "queue/stack-based traversal")
	}
	if strings.Contains(content, "visited") || strings.Contains(content, "seen") {
		approaches = append(approaches, "state tracking to avoid revisiting")
	}
	if strings.Contains(content, "dp") || strings.Contains(content, "memo") {
		approaches = append(approaches, "dynamic programming/memoization")
	}

	if len(approaches) > 0 {
		return "The solution employs " + strings.Join(approaches, ", ") + " to efficiently solve the problem."
	}

	return "Direct implementation solving the problem with standard Go idioms."
}

func updateMainReadme(root string, days []string) {
	var dirList, table string
	for _, day := range days {
		dirList += fmt.Sprintf("- [%s](./%s)\n", day, day)
		status := "✓ Complete"
		if _, err := os.Stat(filepath.Join(root, day, "main.go")); err != nil {
			status = "⏳ In Progress"
		}
		dayNum := getDayNumber(day)
		table += fmt.Sprintf("| %s | %s | [Link](./%s) |\n", dayNum, status, day)
	}

	content := fmt.Sprintf(`# Advent of Code 2025 Solutions

Solutions for Advent of Code 2025 implemented in Go.

## Project Structure

This repository contains daily solutions organized in separate directories:

%s

## Quick Start

Run all solutions:
`+"```bash"+`
make all
`+"```"+`

Run a specific day:
`+"```bash"+`
cd day01
go run main.go
`+"```"+`

## Daily Progress

| Day | Status | Link |
|-----|--------|------|
%s

---

*Last updated: %s*
`, dirList, table, time.Now().Format("2006-01-02 15:04:05"))

	os.WriteFile(filepath.Join(root, "README.md"), []byte(content), 0644)
}

func updateDayReadme(root, day string) {
	dayPath := filepath.Join(root, day)
	mainPath := filepath.Join(dayPath, "main.go")
	inputPath := filepath.Join(dayPath, "input")

	dayNum := getDayNumber(day)

	// Fetch and summarize problem
	problemHTML := fetchProblemDescription(dayNum)
	problemOverview := summarizeProblem(problemHTML)

	solutionApproach := getSolutionApproach(dayPath)
	keySnippets := extractKeySnippets(mainPath)
	inputSample := readLines(inputPath, 7)

	// Try to extract result if available
	result := "Run the solution to see the answer"
	if resultData, err := os.ReadFile(filepath.Join(dayPath, "output.txt")); err == nil {
		result = strings.TrimSpace(string(resultData))
	}

	content := fmt.Sprintf(`# Day %s

## Problem Overview

%s

[View full problem on Advent of Code](https://adventofcode.com/2025/day/%s)

## Solution Approach

%s

## Key Code Snippets

%s

## Input Format

Sample input (first few lines):

`+"```"+`
%s
`+"```"+`

## Running the Solution

`+"```bash"+`
cd %s
go run main.go
`+"```"+`

## Result

`+"```"+`
%s
`+"```"+`

---

*Last updated: %s*
`, dayNum, problemOverview, dayNum, solutionApproach, keySnippets, inputSample, day, result, time.Now().Format("2006-01-02 15:04:05"))

	os.WriteFile(filepath.Join(dayPath, "README.md"), []byte(content), 0644)
}

func readLines(path string, n int) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return "input not found"
	}
	lines := strings.Split(string(data), "\n")
	if len(lines) > n {
		lines = lines[:n]
	}
	return strings.Join(lines, "\n")
}
