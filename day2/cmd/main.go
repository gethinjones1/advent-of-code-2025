package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gethinjones1/advent-of-code-2025/day2"
)

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), ",")
	var result []int
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		result = append(result, day2.FindInvalidIds(line)...)
	}
	fmt.Println(sumInts(result))
}

func sumInts(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
