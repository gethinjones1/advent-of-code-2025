package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gethinjones1/advent-of-code-2025/day1"
)

func main() {
	data, err := os.ReadFile("input-real.txt")

	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	game := day1.NewPasswordBreakerGame()

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		game.PasswordBreaker(line)
	}
	fmt.Println(game.ZeroCount)
}
