package main

import "testing"

func TestAdd(t *testing.T) {
	got := add(10, 5)
	want := 15
	if got != want {
		t.Errorf("Add expectation mismtached. \n\t\t got: %d \n\t\t want: %d ", got, want)
	}

	got = add(-10, -5)
	want = -15
	if got != want {
		t.Errorf("Add expectation mismtached. \n\t\t got: %d \n\t\t want: %d ", got, want)
	}

	got = add(0, 0)
	want = 0
	if got != want {
		t.Errorf("Add expectation mismtached. \n\t\t got: %d \n\t\t want: %d ", got, want)
	}

	got = add(-10, 5)
	want = -5
	if got != want {
		t.Errorf("Add expectation mismtached. \n\t\t got: %d \n\t\t want: %d ", got, want)
	}
}

func TestSub(t *testing.T) {
	got := sub(20, 10)
	want := 10
	if got != want {
		t.Errorf("Sub expectation mismtached. \n\t\t got: %d \n\t\t want: %d ", got, want)
	}
	got = sub(-20, -10)
	want = -10
	if got != want {
		t.Errorf("sub expectation mismtached. \n\t\t got: %d \n\t\t want: %d ", got, want)
	}
	got = sub(0, 0)
	want = 0
	if got != want {
		t.Errorf("sub expectation mismtached. \n\t\t got: %d \n\t\t want: %d ", got, want)
	}
	got = sub(-12, 10)
	want = -22
	if got != want {
		t.Errorf("sub expectation mismtached. \n\t\t got: %d \n\t\t want: %d ", got, want)
	}
}
