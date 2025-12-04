package day3

import "testing"

func TestReturnMaxBatteryOf11(t *testing.T) {
	want := 11
	got := CalculateMaxJoltage("111111111111111")
	if got != want {
		t.Errorf(`CalculateMaxJoltage("111111111111111") = %d, want %d`, got, want)
	}
}

func TestReturnMaxBatteryOf21(t *testing.T) {
	want := 11
	got := CalculateMaxJoltage("111111121111111")
	if got != want {
		t.Errorf(`CalculateMaxJoltage("111111111111111") = %d, want %d`, got, want)
	}
}
