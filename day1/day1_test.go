package day1

import (
	"os"
	"strings"
	"testing"
)

func TestRotateLeftBy5Points(t *testing.T) {
	game := NewPasswordBreakerGame()
	result := game.PasswordBreaker("L5")
	expected := 45
	if result != expected {
		t.Errorf(`PasswordBreaker("L5") = %d, want %d`, result, expected)
	}
}

func TestRotateLeftBy10Points(t *testing.T) {
	game := NewPasswordBreakerGame()
	result := game.PasswordBreaker("L10")
	expected := 40
	if result != expected {
		t.Errorf(`PasswordBreaker("L10") = %d, want %d`, result, expected)
	}
}

func TestRotateLeftPast0Points(t *testing.T) {
	game := NewPasswordBreakerGame()
	result := game.PasswordBreaker("L55")
	expected := 95
	if result != expected {
		t.Errorf(`PasswordBreaker("L55") = %d, want %d`, result, expected)
	}
}

func TestRotateRightPast99Points(t *testing.T) {
	game := NewPasswordBreakerGame()
	result := game.PasswordBreaker("R55")
	expected := 5
	if result != expected {
		t.Errorf(`PasswordBreaker("L55") = %d, want %d`, result, expected)
	}
}

func TestRotateRightPast99PointsAndThenLeft2(t *testing.T) {
	game := NewPasswordBreakerGame()
	game.PasswordBreaker("R55")
	result := game.PasswordBreaker("L2")
	expected := 3
	if result != expected {
		t.Errorf(`PasswordBreaker("L55") = %d, want %d`, result, expected)
	}
}
func TestStartAt5RotateL10AndThenR5(t *testing.T) {
	game := NewPasswordBreakerGame()
	game.PasswordBreaker("L45")
	game.PasswordBreaker("L10")
	result := game.PasswordBreaker("R5")
	expected := 0
	if result != expected {
		t.Errorf(`got %d, want %d`, result, expected)
	}
}

func TestZeroCount(t *testing.T) {
	game := NewPasswordBreakerGame()
	game.PasswordBreaker("L50")  // hits 0
	game.PasswordBreaker("L3")   // hits 97
	game.PasswordBreaker("L97")  // hits 0
	game.PasswordBreaker("R100") // hits 0
	result := game.ZeroCount
	expected := 3
	if result != expected {
		t.Errorf(`game.zeroCount should be 3 = %d, want %d`, result, expected)
	}
}

// Example from the puzzle description
func TestExampleSequenceFromSpec(t *testing.T) {
	game := NewPasswordBreakerGame()

	rotations := []string{
		"L68",
		"L30",
		"R48",
		"L5",
		"R60",
		"L55",
		"L1",
		"L99",
		"R14",
		"L82",
	}

	var pos int
	for _, r := range rotations {
		pos = game.PasswordBreaker(r)
	}

	if pos != 32 {
		t.Errorf("final dial position = %d, want %d", pos, 32)
	}

	if game.ZeroCount != 3 {
		t.Errorf("ZeroCount = %d, want %d", game.ZeroCount, 3)
	}
}

// First 10 rotations from your real data
// Starting at 50, after these the dial should be 40 and we never hit 0.
func TestRealDataFirst10(t *testing.T) {
	game := NewPasswordBreakerGame()

	rotations := []string{
		"R19", // 69
		"R5",  // 74
		"R29", // 103 +1
		"R30", // 133
		"L24", // 109
		"L12", // 97  +1
		"R24", // 121 +1
		"R2",  // 123
		"R35", // 158
		"L18",
	}

	var pos int
	for _, r := range rotations {
		pos = game.PasswordBreaker(r)
	}

	if pos != 40 {
		t.Errorf("final dial position = %d, want %d", pos, 40)
	}

	if game.ZeroCount != 3 {
		t.Errorf("ZeroCount = %d, want %d", game.ZeroCount, 3)
	}
}

// First 20 rotations from your real data
// After these, starting at 50, dial should be 64 and still no zero hits.
func TestRealDataFirst20(t *testing.T) {
	game := NewPasswordBreakerGame()

	rotations := []string{
		"R19",
		"R5",
		"R29",
		"R30",
		"L24",
		"L12",
		"R24",
		"R2",
		"R35",
		"L18",
		"R22",
		"L20",
		"L1",
		"R12",
		"R46",
		"L31",
		"L38",
		"L27",
		"L33",
		"L6",
	}

	var pos int
	for _, r := range rotations {
		pos = game.PasswordBreaker(r)
	}

	if pos != 64 {
		t.Errorf("final dial position = %d, want %d", pos, 64)
	}

	if game.ZeroCount != 0 {
		t.Errorf("ZeroCount = %d, want %d", game.ZeroCount, 0)
	}
}

// Rotations 21–30 from your real data (starting fresh at 50)
// This sequence hits 0 exactly once (on the first step) and ends at 89.
func TestRealDataTwentyOneToThirty(t *testing.T) {
	game := NewPasswordBreakerGame()

	rotations := []string{
		"R50",
		"L22",
		"L37",
		"L18",
		"R7",
		"L28",
		"R50",
		"L44",
		"R5",
		"L24",
	}

	var pos int
	for _, r := range rotations {
		pos = game.PasswordBreaker(r)
	}

	if pos != 89 {
		t.Errorf("final dial position = %d, want %d", pos, 89)
	}

	if game.ZeroCount != 1 {
		t.Errorf("ZeroCount = %d, want %d", game.ZeroCount, 1)
	}
}

// Big rotation test: R805 is equivalent to R5 on a 0–99 dial.
// From 50, this should land on 55.
func TestRotateRightMoreThanFullCircle(t *testing.T) {
	game := NewPasswordBreakerGame()
	pos := game.PasswordBreaker("R805")

	expected := 55
	if pos != expected {
		t.Errorf(`PasswordBreaker("R805") = %d, want %d`, pos, expected)
	}
}

// Big rotation test: L943 is equivalent to L43 (943 % 100).
// From 50, 50 - 43 = 7 (with wrap).
func TestRotateLeftMoreThanFullCircle(t *testing.T) {
	game := NewPasswordBreakerGame()
	pos := game.PasswordBreaker("L943")

	expected := 7
	if pos != expected {
		t.Errorf(`PasswordBreaker("L943") = %d, want %d`, pos, expected)
	}
}

func TestRealDataFullGame(t *testing.T) {
	game := NewPasswordBreakerGame()
	data, err := os.ReadFile("input-real.txt")

	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	for _, r := range lines {
		game.PasswordBreaker(r)
	}

	if game.ZeroCount != 1078 {
		t.Errorf("ZeroCount = %d, want %d", game.ZeroCount, 1078)
	}
}

func TestRight150WrapsToZero(t *testing.T) {
	game := NewPasswordBreakerGame()

	pos := game.PasswordBreaker("R150")

	if pos != 0 {
		t.Fatalf(`PasswordBreaker("R150") = %d, want %d`, pos, 0)
	}

	if game.ZeroCount != 1 {
		t.Fatalf("ZeroCount = %d, want %d", game.ZeroCount, 1)
	}
}

func TestCountZeroWhenPassesZero(t *testing.T) {
	game := NewPasswordBreakerGame()
	game.PasswordBreaker("L1000")

	if game.ZeroCount != 10 {
		t.Errorf("ZeroCount = %d, want %d", game.ZeroCount, 10)
	}
}

func TestPart2ExampleCountsZeroDuringRotations(t *testing.T) {
	game := NewPasswordBreakerGame()

	rotations := []string{
		"L68", // ends at 82 (+1: passes 0)
		"L30", // ends at 52
		"R48", // ends at 0 (+1: ends on 0)
		"L5",  // ends at 95
		"R60", // ends at 55 (+1: passes 0)
		"L55", // ends at 0 (+1: ends on 0)
		"L1",  // ends at 99
		"L99", // ends at 0 (+1: ends on 0)
		"R14", // ends at 14
		"L82", // ends at 32 (+1: passes 0)
	}

	for _, r := range rotations {
		game.PasswordBreaker(r)
	}

	if game.ZeroCount != 6 {
		t.Fatalf("ZeroCount = %d, want %d", game.ZeroCount, 6)
	}
}

func TestRight150From50HitsZeroTwice(t *testing.T) {
	game := NewPasswordBreakerGame()

	// From 50:
	// R150 = 150 clicks to the right.
	// - After 50 clicks, we hit 0 once.
	// - After another 100 clicks (total 150), we land on 0 again.
	//
	// So:
	// - Final position should be 0.
	// - ZeroCount should be 2 (hit 0 twice in this single rotation).
	pos := game.PasswordBreaker("R150")

	if pos != 0 {
		t.Fatalf(`PasswordBreaker("R150") position = %d, want %d`, pos, 0)
	}

	if game.ZeroCount != 2 {
		t.Fatalf(`PasswordBreaker("R150") ZeroCount = %d, want %d`, game.ZeroCount, 2)
	}
}

func TestLeftRotationCrossesZeroFromFive(t *testing.T) {
	game := NewPasswordBreakerGame()

	// Move from 50 to 5 without ever hitting 0:
	// 50 L45 -> 5 (path: 49,48,...,5; never 0)
	pos := game.PasswordBreaker("L45")
	if pos != 5 {
		t.Fatalf(`after "L45" position = %d, want %d`, pos, 5)
	}
	if game.ZeroCount != 0 {
		t.Fatalf(`after "L45" ZeroCount = %d, want %d`, game.ZeroCount, 0)
	}

	// Now from 5:
	// 5 L5 -> 0
	// Clicks: 4,3,2,1,0 -> we should hit 0 exactly once.
	pos = game.PasswordBreaker("L5")
	if pos != 0 {
		t.Fatalf(`after "L5" position = %d, want %d`, pos, 0)
	}
	if game.ZeroCount != 1 {
		t.Fatalf(`after "L45","L5" ZeroCount = %d, want %d`, game.ZeroCount, 1)
	}
}
