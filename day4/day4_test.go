package day4

import "testing"

func TestCheckIfForkliftCanMoveForSmallArea(t *testing.T) {
	want := 1
	got := CountAccessPaths(".@")
	if want != got {
		t.Errorf("got = %q want = %q", got, want)
	}
}

func TestCountAccessPathsInTwoRowGrid(t *testing.T) {
	input := "@.\n.@"
	want := 2 // both @ have < 4 neighbours, so both are accessible

	got := CountAccessPaths(input)
	if got != want {
		t.Errorf("got = %d want = %d", got, want)
	}
}

func TestSomeRollsAreInaccessibleInDenseArea(t *testing.T) {
	input := "" +
		"@@@\n" +
		"@@@"

	want := 4

	got := CountAccessPaths(input)
	if got != want {
		t.Errorf("got = %d want = %d", got, want)
	}
}

func TestDiagonalNeighboursCountToo(t *testing.T) {
	input := "" +
		"@.@\n" +
		".@.\n" +
		"@.@"

	// The center @ has 4 diagonal neighbours, so it is NOT accessible.
	// The four corner @ are accessible (each has only 1 neighbour).
	want := 4

	got := CountAccessPaths(input)
	if got != want {
		t.Errorf("got = %d want = %d", got, want)
	}
}

func TestHandlesRowsOfDifferentLengths(t *testing.T) {
	input := "" +
		"@@@@\n" + // length 4
		"@@\n" // length 2

	// All @ have < 4 neighbours here → all accessible.
	want := 6

	got := CountAccessPaths(input)
	if got != want {
		t.Errorf("got = %d want = %d", got, want)
	}
}
