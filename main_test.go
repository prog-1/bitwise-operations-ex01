package main

import (
	"math"
	"testing"
)

func TestEven(t *testing.T) {
	for _, tc := range []struct {
		num  uint64
		want bool
	}{
		{0, true},
		{69, false},
		{96, true},
		{120, true},
		{11, false},
	} {
		if got := Even(tc.num); got != tc.want {
			t.Errorf("Even(%v): got = %v, want = %v", tc.num, got, tc.want)
		}
	}
}
func TestOdd(t *testing.T) {
	for _, tc := range []struct {
		num  uint64
		want bool
	}{
		{0, false},
		{96, false},
		{69, true},
		{120, false},
		{11, true},
	} {
		if got := Odd(tc.num); got != tc.want {
			t.Errorf("Odd(%v): got = %v, want = %v", tc.num, got, tc.want)
		}
	}
}
func TestSetBit(t *testing.T) {
	for _, tc := range []struct {
		num  uint64
		n    int
		want uint64
	}{
		{0, 0, 1},
		{1, 8, 257},
		{1024, 1, 1026},
		{1, 0, 1},
		{256, 8, 256},
	} {
		if got := SetBit(tc.num, tc.n); got != tc.want {
			t.Errorf("SetBit(%v): got = %v, want = %v", tc.num, got, tc.want)
		}
	}
}
func TestClearBit(t *testing.T) {
	for _, tc := range []struct {
		num  uint64
		n    int
		want uint64
	}{
		{0, 0, 0},
		{1, 0, 0},
		{257, 8, 1},
		{1026, 1, 1024},
		{1, 1, 1},
		{256, 0, 256},
	} {
		if got := ClearBit(tc.num, tc.n); got != tc.want {
			t.Errorf("ClearBit(%v, %v): got = %v, want = %v", tc.num, tc.n, got, tc.want)
		}
	}
}
func TestFlipBit(t *testing.T) {
	for _, tc := range []struct {
		num  uint64
		n    int
		want uint64
	}{
		{1, 0, 0},
		{0, 0, 1},
		{258, 1, 256},
		{256, 1, 258},
		{3, 1, 1},
	} {
		if got := FlipBit(tc.num, tc.n); got != tc.want {
			t.Errorf("FlipBit(%v, %v): got = %v, want = %v", tc.num, tc.n, got, tc.want)
		}
	}
}
func TestOnes(t *testing.T) {
	for _, tc := range []struct {
		num  uint64
		want int
	}{
		{0, 0},
		{15, 4},
		{240, 4},
		{0xf0, 4},
		{0xffffffffffffffff, 64},
		{math.MaxUint64, 64},
	} {
		if got := Ones(tc.num); got != tc.want {
			t.Errorf("Ones(%v): got = %v, want = %v", tc.num, got, tc.want)
		}
	}
}
func TestReverse(t *testing.T) {
	for _, tc := range []struct {
		num  uint64
		want uint64
	}{
		{0, 0},
		{0xff00000000000000, 0xff},
		{0xff, 0xff00000000000000},
		{13, 12682136550675316736},
		{12682136550675316736, 13},
	} {
		if got := Reverse(tc.num); got != tc.want {
			t.Errorf("Reverse(%v): got = %v, want = %v", tc.num, got, tc.want)
		}
	}
}
func TestLeadingZeroes(t *testing.T) {
	for _, tc := range []struct {
		num  uint64
		want int
	}{
		{0, 64},
		{1, 63},
		{0x7fffffffffffffff, 1},
		{0x7000000000000000, 1},
		{0xffffffffffffffff, 0},
	} {
		if got := LeadingZeroes(tc.num); got != tc.want {
			t.Errorf("LeadingZeroes(%v): got = %v, want = %v", tc.num, got, tc.want)
		}
	}
}
func TestTrailingZeroes(t *testing.T) {
	for _, tc := range []struct {
		num  uint64
		want int
	}{
		{0, 64},
		{1, 0},
		{2, 1},
		{0x8000000000000000, 63},
	} {
		if got := TrailingZeroes(tc.num); got != tc.want {
			t.Errorf("TrailingZeroes(%v): got = %v, want = %v", tc.num, got, tc.want)
		}
	}
}
