package day1

import (
	"math"
	"strconv"
)

const (
	startPosition = 50
	lockSize      = 100
)

type PasswordBreakerGame struct {
	currentPosition int
	ZeroCount       int
}

func NewPasswordBreakerGame() *PasswordBreakerGame {
	return &PasswordBreakerGame{startPosition, 0}
}

func (p *PasswordBreakerGame) PasswordBreaker(position string) int {
	direction := position[0:1]
	change := position[1:]
	changeInt := convertStringToInt(change)
	result := mathsWorker(direction, p.currentPosition, changeInt)
	p.currentPosition = result
	if result == 0 {
		p.ZeroCount++
	}

	return result
}

func convertStringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}

	return i
}

func resolveToLockRange(i int) int {
	if i < 0 {
		return lockSize + i // IF int -5, 100 + (-5) = 95
	}

	if i >= lockSize {
		return i - lockSize
	}

	return i
}

func normaliseDelta(delta int) int {
	return int(math.Abs(float64(delta))) % lockSize
}

func mathsWorker(direction string, initial int, change int) int {
	change = normaliseDelta(change)
	if direction == "L" {
		return resolveToLockRange(initial - change)
	}
	return resolveToLockRange(initial + change)
}
