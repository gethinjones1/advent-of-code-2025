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
	changeInt := ConvertStringToInt(change)

	p.ZeroCount += zerosOnRotation(direction, p.currentPosition, changeInt)
	p.currentPosition = findNextPosition(direction, p.currentPosition, changeInt)

	return p.currentPosition
}

func zerosOnRotation(direction string, start, delta int) int {
	if delta < 0 {
		delta = -delta
	}
	if delta == 0 {
		return 0
	}

	var first int

	if direction == "R" {
		if start == 0 {
			first = lockSize
		} else {
			first = lockSize - start
		}
	} else {
		if start == 0 {
			first = lockSize
		} else {
			first = start
		}
	}

	if first > delta {
		return 0
	}

	return 1 + (delta-first)/lockSize
}

func ConvertStringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func resolveToLockRange(i int) int {
	if i < 0 {
		return lockSize + i
	}
	if i >= lockSize {
		return i - lockSize
	}
	return i
}

func normaliseDelta(delta int) int {
	return int(math.Abs(float64(delta))) % lockSize
}

func findNextPosition(direction string, initial int, change int) int {
	change = normaliseDelta(change)
	if direction == "L" {
		return resolveToLockRange(initial - change)
	}
	return resolveToLockRange(initial + change)
}
