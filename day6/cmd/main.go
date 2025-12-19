package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gethinjones1/advent-of-code-2025/day6"
)

func parseInput(input string) [][]interface{} {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	result := make([][]interface{}, len(lines))

	for i, line := range lines {
		parts := strings.Fields(line) // Split by whitespace
		row := make([]interface{}, len(parts))

		for j, part := range parts {
			// Try to convert to int, if it fails, keep as string
			if num, err := strconv.Atoi(part); err == nil {
				row[j] = num
			} else {
				row[j] = part // Keep as string (operators like "+", "*")
			}
		}
		result[i] = row
	}

	return result
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	parsed := parseInput(string(data))

	fmt.Println(parsed)

	got := day6.CalculateMath(parsed)
	fmt.Println(got)
}
