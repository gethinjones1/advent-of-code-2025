package main

import (
	"fmt"
	"os"

	"github.com/gethinjones1/advent-of-code-2025/day4"
)

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		panic(err)
	}

	fmt.Println(day4.CountTotalRemoved(string(data)))
}
