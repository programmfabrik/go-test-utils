package go_test_utils

import "testing"

func TestAssertStringEquals(t *testing.T) {
	iT := &testing.T{}
	AssertStringEquals(iT, "hallo", "hall")
	if !iT.Failed() {
		t.Error("AssertStringEquals did not fail but should")
	}
}

func TestAssertStringArraysEqualNoOrder(t *testing.T) {
	want := []string{
		"1",
		"2",
		"3",
	}
	got := []string{
		"2",
		"1",
		"3",
	}

	iT := &testing.T{}
	AssertStringArraysEqualNoOrder(iT, got, want)
	if iT.Failed() {
		t.Error("(Equal) AssertStringArraysEqualNoOrder did fail but should not")
	}

	want = []string{
		"1",
		"2",
		"3",
	}
	got = []string{
		"2",
		"1",
		"3",
		"4",
	}

	iT = &testing.T{}
	AssertStringArraysEqualNoOrder(iT, got, want)
	if !iT.Failed() {
		t.Error("(Too much in got) AssertStringArraysEqualNoOrder did not fail but should")
	}

	want = []string{
		"1",
		"2",
		"3",
		"4",
	}
	got = []string{
		"2",
		"1",
		"3",
	}

	iT = &testing.T{}
	AssertStringArraysEqualNoOrder(iT, got, want)
	if !iT.Failed() {
		t.Error("(Too much in have) AssertStringArraysEqualNoOrder did not fail but should")
	}
}
