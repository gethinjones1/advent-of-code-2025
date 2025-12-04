package day3

import "testing"

func TestCheckBankJoltage_987654321111111(t *testing.T) {
	want := 987654321111
	got := MaxBankJoltagePt2("987654321111111")

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestCheckBankJoltage_811111111111119(t *testing.T) {
	want := 811111111119
	got := MaxBankJoltagePt2("811111111111119")

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestCheckBankJoltage_234234234234278(t *testing.T) {
	want := 434234234278
	got := MaxBankJoltagePt2("234234234234278")

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestCheckBankJoltage_818181911112111(t *testing.T) {
	want := 888911112111
	got := MaxBankJoltagePt2("818181911112111")

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
