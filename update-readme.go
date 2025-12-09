package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
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
	return strings.TrimPrefix(day, "day")
}

func fetchProblemDescription(dayNum string) string {
	url := fmt.Sprintf("https://adventofcode.com/2025/day/%s", dayNum)
	resp, err := http.Get(url)
	if err != nil {
		return "Could not fetch problem description."
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Could not read problem description."
	}

	content := string(body)
	start := strings.Index(content, "<article class=\"day-desc\">")
	end := strings.Index(content, "</article>")

	if start == -1 || end == -1 {
		return "Problem description not found."
	}

	extracted := content[start+26 : end]
	extracted = strings.ReplaceAll(extracted, "<p>", "")
	extracted = strings.ReplaceAll(extracted, "</p>", "\n\n")
	extracted = strings.ReplaceAll(extracted, "<h2>", "## ")
	extracted = strings.ReplaceAll(extracted, "</h2>", "")
	extracted = strings.ReplaceAll(extracted, "<code>", "`")
	extracted = strings.ReplaceAll(extracted, "</code>", "`")
	extracted = strings.ReplaceAll(extracted, "<pre>", "```\n")
	extracted = strings.ReplaceAll(extracted, "</pre>", "\n```")
	extracted = strings.ReplaceAll(extracted, "<em>", "*")
	extracted = strings.ReplaceAll(extracted, "</em>", "*")
	extracted = strings.ReplaceAll(extracted, "<strong>", "**")
	extracted = strings.ReplaceAll(extracted, "</strong>", "**")
	extracted = strings.ReplaceAll(extracted, "&quot;", "\"")
	extracted = strings.ReplaceAll(extracted, "&amp;", "&")

	return strings.TrimSpace(extracted)
}

func updateMainReadme(root string, days []string) {
	var dirList, table string
	for _, day := range days {
		dirList += fmt.Sprintf("- [%s](./%s)\n", day, day)
		status := "✓ Complete"
		if _, err := os.Stat(filepath.Join(root, day, "main.go")); err != nil {
			status = "⏳ In Progress"
		}
		table += fmt.Sprintf("| %s | %s | [Link](./%s) |\n", day, status, day)
	}
	content := fmt.Sprintf(`# Advent of Code 2025 Solutions

Solutions for Advent of Code 2025 in Go.

## Directory Structure

%s

## Quick Start

make all

## Daily Progress

| Day | Status | Link |
|-----|--------|------|
%s

_Last updated: %s_
`, dirList, table, time.Now().Format("2006-01-02 15:04:05"))
	os.WriteFile(filepath.Join(root, "README.md"), []byte(content), 0644)
}

func updateDayReadme(root, day string) {
	mainPath := filepath.Join(root, day, "main.go")
	inputPath := filepath.Join(root, day, "input")
	mainCode := readFile(mainPath)
	inputSample := readLines(inputPath, 10)
	dayNum := getDayNumber(day)
	problemDesc := fetchProblemDescription(dayNum)

	content := day + "\n\n" +
		"## Problem Description\n\n" +
		problemDesc + "\n\n" +
		"## Solution Overview\n\n" +
		"Simulates dial rotations and counts zero crossings.\n\n" +
		"## Code\n\n" +
		"```go\n" +
		mainCode + "\n" +
		"```\n\n" +
		"## Input Sample\n\n" +
		"```\n" +
		inputSample + "\n" +
		"```\n\n" +
		"## Running\n\n" +
		"cd " + day + "\n" +
		"go run main.go\n\n" +
		"_Last updated: " + time.Now().Format("2006-01-02 15:04:05") + "_"

	os.WriteFile(filepath.Join(root, day, "README.md"), []byte(content), 0644)
}

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		return "// main.go not found"
	}
	return string(data)
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
