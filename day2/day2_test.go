package day2

import (
	"reflect"
	"testing"
)

func TestInvalidIdNoneReturned(t *testing.T) {
	var want []int
	got := FindInvalidIds("9-10")
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestInvalidIdOneReturned(t *testing.T) {
	want := []int{11}
	got := FindInvalidIds("11-20")
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestInvalidIdOneReturnedThreeConsecutive(t *testing.T) {
	want := []int{123123}
	got := FindInvalidIds("123120-123125") // 123123
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestInvalidIdLongSymmetricalNumber(t *testing.T) {
	want := []int{1188511885}
	got := FindInvalidIds("1188511880-1188511890")
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestInvalidId7Ones(t *testing.T) {
	want := []int{1111111}
	got := FindInvalidIds("1111110-1111115")
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestInvalidIdSequenceOfNone(t *testing.T) {
	var want []int
	got := FindInvalidIds("565654-565655")
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %d, want %d", got, want)
	}
}

func TestInvalidIdsTwoReturned11And22(t *testing.T) {
	want := []int{11, 22}
	got := FindInvalidIds("11-22")
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestInvalidIdsTwoReturned99And111(t *testing.T) {
	want := []int{99, 111}
	got := FindInvalidIds("95-115")
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestInvalidIdsTwoReturned999And1010(t *testing.T) {
	want := []int{999, 1010}
	got := FindInvalidIds("998-1012")
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v, want %v", got, want)
	}
}

// 1188511880-1188511890 still has one invalid ID, 1188511885
// (already covered by TestInvalidIdLongSymmetricalNumber)

func TestInvalidIdAllTwos222222(t *testing.T) {
	want := []int{222222}
	got := FindInvalidIds("222220-222224")
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestInvalidIdNoneIn1698522To1698528(t *testing.T) {
	var want []int
	got := FindInvalidIds("1698522-1698528")
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestInvalidId446446(t *testing.T) {
	want := []int{446446}
	got := FindInvalidIds("446443-446449")
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestInvalidId38593859(t *testing.T) {
	want := []int{38593859}
	got := FindInvalidIds("38593856-38593862")
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestInvalidId565656(t *testing.T) {
	want := []int{565656}
	got := FindInvalidIds("565653-565659")
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestInvalidId824824824(t *testing.T) {
	want := []int{824824824}
	got := FindInvalidIds("824824823-824824825")
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestInvalidId2121212121(t *testing.T) {
	want := []int{2121212121}
	got := FindInvalidIds("2121212118-2121212124")
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestPatternRepeatFinder(t *testing.T) {
	want := true
	got := CheckPatternRepeated(123123123)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("got %T, want %T", got, want)
	}
}
