package day5

import "testing"

func TestCheckFoodIsFresh(t *testing.T) {
	want := 1
	got := CheckFoodIsFresh([]string{"1-1"}, 1)

	if want != got {
		t.Errorf("got = %q want = %q", got, want)
	}
}

func TestCheckFoodIsFreshWithSmallRange(t *testing.T) {
	want := 1
	got := CheckFoodIsFresh([]string{"1-2"}, 1)

	if want != got {
		t.Errorf("got = %q want = %q", got, want)
	}
}

func TestCheckFoodIsntFreshWithSmallRange(t *testing.T) {
	want := 0
	got := CheckFoodIsFresh([]string{"2-3"}, 1)

	if want != got {
		t.Errorf("got = %q want = %q", got, want)
	}
}

func TestCheckFoodIsFreshWithinSeconRange(t *testing.T) {
	want := 1
	got := CheckFoodIsFresh([]string{"1-2", "2-3"}, 1)

	if want != got {
		t.Errorf("got = %q want = %q", got, want)
	}
}

func TestCheckFoodIsFreshWithinTheRangeBounds(t *testing.T) {
	want := 1
	got := CheckFoodIsFresh([]string{"1-2", "30-35"}, 33)

	if want != got {
		t.Errorf("got = %q want = %q", got, want)
	}
}

func TestCheckFoodIsFreshWithinTheRangeBoundsPotentialDoubleCount(t *testing.T) {
	want := 1
	got := CheckFoodIsFresh([]string{"1-2", "30-35", "33-35"}, 33)

	if want != got {
		t.Errorf("got = %q want = %q", got, want)
	}
}
