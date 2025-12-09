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
2. **Problem Overview:** 2-3 sentence summary of the challenge (what needs to be solved, not the full story/context)
3. **Solution Approach:** Brief explanation of the algorithmic approach and key insights
4. **Key Code Snippets:** Extract only the most interesting/clever parts:
   - Core algorithm or data structure
   - Novel problem-solving logic
   - Tricky parsing or transformation logic
   - Omit boilerplate (file reading, error handling, main function setup)
   - Keep snippets focused (5-20 lines each)
5. **Input Format:** Show sample input (first 5-10 lines) with brief description
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
- Include code blocks with language specification (`go`)
- Use tables for progress tracking
- Add horizontal rules between sections
- Include relative links between day folders
- Keep it concise - focus on clarity over completeness
