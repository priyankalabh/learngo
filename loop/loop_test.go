package main

import "testing"

func TestIsPrime(t *testing.T) {
	got := isPrime(7)
	want := true
	if got != want {
		t.Errorf("Prime expectations mismatched.\n got: %t\t\twant: %t", got, want)
	}

	got = isPrime(4)
	want = false
	if got != want {
		t.Errorf("Prime expectations mismatched.\n got: %t\t\twant: %t", got, want)

	}
}

func TestSumOddNumbers(t *testing.T) {
	got := sumOddNumbers(5)
	want := 9
	if got != want {
		t.Errorf("odd expectations mismatched.\n got: %d\t\twant: %d", got, want)

	}

	got = sumOddNumbers(0)
	want = 0
	if got != want {
		t.Errorf("odd expectations mismatched.\n got: %d\t\twant: %d", got, want)

	}

	got = sumOddNumbers(1)
	want = 1
	if got != want {
		t.Errorf("odd expectations mismatched.\n got: %d\t\twant: %d", got, want)

	}
}

func TestCalculateDigits(t *testing.T) {
	got := calculateDigits(1234)
	want := 4
	if got != want {
		t.Errorf("Digits Expextations mismatched.\n got:%d\t\twant: %d", got, want)
	}

	got = calculateDigits(0)
	want = 1
	if got != want {
		t.Errorf("Digits Expextations mismatched.\n got:%d\t\twant: %d", got, want)
	}
}
