package main

import (
	"testing"
)

func TestEven(t *testing.T) {
	for _, tc := range []struct {
		num  uint64
		want bool
	}{
		{num: 69, want: false},
		{num: 96, want: true},
		{num: 70, want: true},
		{num: 71, want: false},
	} {
		if got := Even(tc.num); got != tc.want {
			t.Errorf("Even(%v) = (%v), want = (%v)", tc.num, got, tc.want)
		}
	}
}
func TestOdd(t *testing.T) {
	for _, tc := range []struct {
		num  uint64
		want bool
	}{
		{num: 69, want: true},
		{num: 96, want: false},
		{num: 70, want: false},
		{num: 71, want: true},
	} {
		if got := Odd(tc.num); got != tc.want {
			t.Errorf("Odd(%v) = (%v), want = (%v)", tc.num, got, tc.want)
		}
	}
}
func TestSetBit(t *testing.T) {
	for _, tc := range []struct {
		num  uint64
		n    int
		want uint64
	}{
		{num: 0, n: 0, want: 1},
		{num: 1, n: 8, want: 257},
		{num: 1024, n: 1, want: 1026},
		{num: 1, n: 0, want: 1},
	} {
		if got := SetBit(tc.num, tc.n); got != tc.want {
			t.Errorf("SetBit(%v) = (%v), want = (%v)", tc.num, got, tc.want)
		}
	}
}
func TestClearBit(t *testing.T) {
	for _, tc := range []struct {
		num  uint64
		n    int
		want uint64
	}{
		{num: 1, n: 0, want: 0},
		{num: 257, n: 8, want: 1},
		{num: 1026, n: 1, want: 1024},
		{num: 1, n: 1, want: 1},
	} {
		if got := ClearBit(tc.num, tc.n); got != tc.want {
			t.Errorf("SetBit(%v) = (%v), want = (%v)", tc.num, got, tc.want)
		}
	}
}
func TestFlipBit(t *testing.T) {
	for _, tc := range []struct {
		num  uint64
		n    int
		want uint64
	}{
		{num: 1, n: 0, want: 0},
		{num: 0, n: 0, want: 1},
		{num: 258, n: 1, want: 256},
		{num: 256, n: 1, want: 258},
	} {
		if got := FlipBit(tc.num, tc.n); got != tc.want {
			t.Errorf("SetBit(%v) = (%v), want = (%v)", tc.num, got, tc.want)
		}
	}
}
