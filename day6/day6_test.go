package day6

import "testing"

func TestAddTwoNumbersTogether(t *testing.T) {
	want := 3
	got := CalculateMath([][]interface{}{
		{1},
		{2},
		{"+"},
	})

	if want != got {
		t.Errorf("got = %d, want = %d", got, want)
	}
}

func TestTimesTwoNumbersTogether(t *testing.T) {
	want := 2
	got := CalculateMath([][]interface{}{
		{1},
		{2},
		{"*"},
	})

	if want != got {
		t.Errorf("got = %d, want = %d", got, want)
	}
}

func TestTimesThreeNumbersTogether(t *testing.T) {
	want := 4
	got := CalculateMath([][]interface{}{
		{1},
		{2},
		{2},
		{"*"},
	})

	if want != got {
		t.Errorf("got = %d, want = %d", got, want)
	}
}

func TestOneOfTheExamples(t *testing.T) {
	want := 33210
	got := CalculateMath([][]interface{}{
		{123},
		{45},
		{6},
		{"*"},
	})

	if want != got {
		t.Errorf("got = %d, want = %d", got, want)
	}
}
