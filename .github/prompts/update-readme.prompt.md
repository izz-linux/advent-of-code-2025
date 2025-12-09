---
name: Update README
description: Instructions for automatically updating README files for each Advent of Code day.
---

## Update README for Advent of Code Solutions

Automatically generate and update `README.md` files for each day's solution with the following structure:

### Main README.md (Root Level)

1. **Project Title:** "Advent of Code 2025 Solutions"
2. **Description:** Overview of the project structure and purpose
3. **Directory Structure:** List all completed days
4. **Quick Start:** Instructions to run all solutions
5. **Daily Progress:** Table showing completion status for each day

### Per-Day README.md (day##/README.md)

For each day, include:

1. **Day Title:** "Day ##: [Problem Name]"
2. **Problem Description:** Extracted from problem statement
3. **Solution Overview:** Brief explanation of the approach
4. **Code:** Display the main.go file
5. **Input Format:** Show sample input (first 10 lines)
6. **Running the Solution:** Command to execute
7. **Result:** Final answer/password
8. **Time Complexity:** Performance notes (if applicable)

### Automation

- Generate README files automatically when new day solutions are added
- Update progress table in main README
- Include timestamps of last update
- Link between days for easy navigation

### Formatting Standards

- Use Markdown headings (H1-H3)
- Include code blocks with language specification
- Use tables for progress tracking
- Add horizontal rules between sections
- Include relative links between day folders
