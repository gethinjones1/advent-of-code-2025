package day3

import (
	"testing"
)

func TestTwoDigitBankSupplied(t *testing.T) {
	want := 12
	got := MaxBankJoltage("12")

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestThreeDigitsBankSuppliedAlreadyInOrder(t *testing.T) {
	want := 23
	got := MaxBankJoltage("123")

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestThreeDigitBankSupplied(t *testing.T) {
	want := 32
	got := MaxBankJoltage("132")

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestFullDigitBankSupplied(t *testing.T) {
	want := 98
	got := MaxBankJoltage("987654321111111")

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestFullDigitBankSuppliedFirstAndLast(t *testing.T) {
	want := 89
	got := MaxBankJoltage("811111111111119")

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestLargeBankExample_234234234234278(t *testing.T) {
	want := 78
	got := MaxBankJoltage("234234234234278")

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestLargeBankExample_818181911112111(t *testing.T) {
	want := 92
	got := MaxBankJoltage("818181911112111")

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
