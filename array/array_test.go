package main

import (
	"reflect"
	"testing"
)

func TestSquareArray(t *testing.T) {
	got := squareArray([]int{3, 6, 2, 4})
	want := []int{9, 36, 4, 16}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Square expectation mismathced\n got=%v\nwant=%v", got, want)
	}
	got = squareArray([]int{})
	want = []int{}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Square expectation mismathced\n got=%v\nwant=%v", got, want)
	}

}
func TestMissingInArray(t *testing.T) {
	got := missingInArray([]int{1, 3, 4, 5}, 5)
	want := 2
	if got != want {
		t.Errorf("Missing number in array expectation mismathced\n got=%d\nwant=%d", got, want)

	}
	got = missingInArray([]int{1, 2, 3, 4, 5}, 5)
	want = 0
	if got != want {
		t.Errorf("Missing number in array expectation mismathced\n got=%d\nwant=%d", got, want)
	}

}
func TestModifiedArray(t *testing.T) {
	gotOdd, gotEven := modifiedArray([]int{2, 3, 4, 5, 6, 7})
	wantOdd := []int{3, 5, 7}
	wantEven := []int{3, 5, 7}
	if !reflect.DeepEqual(gotOdd, wantOdd) {
		t.Errorf("Odd array expectation mismathced\n got=%v\nwant=%v", gotOdd, wantOdd)
	}
	if !reflect.DeepEqual(gotEven, wantEven) {
		t.Errorf("Even array expectation mismathced\n got=%v\nwant=%v", gotEven, wantEven)
	}
}
