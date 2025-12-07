package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/gethinjones1/advent-of-code-2025/day5"
)

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	ranges, ids := SplitDatabase(string(data))
	var result []int
	for _, v := range ids {
		id, _ := strconv.Atoi(v)
		result = append(result, day5.CheckFoodIsFresh(ranges, id))
	}

	fmt.Println(SumInts(result))
}

func SplitDatabase(input string) ([]string, []string) {
	lines := strings.Split(input, "\n")

	var ranges []string
	var ids []string

	isRangeSection := true

	for _, line := range lines {
		if line == "" {
			isRangeSection = false
			continue
		}

		if isRangeSection {
			ranges = append(ranges, line)
		} else {
			ids = append(ids, line)
		}
	}

	return ranges, ids
}

func SumInts(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}
