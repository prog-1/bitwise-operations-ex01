package main

import "testing"

func TestEven(t *testing.T) {
	for _, tc := range []struct {
		n    uint64
		want bool
	}{
		{96, true},
		{69, false},
		{1, false},
		{10000, true},
	} {
		if got := Even(tc.n); got != tc.want {
			t.Errorf("Even: got %v want %v", got, tc.want)
		}
	}
}

func TestOdd(t *testing.T) {
	for _, tc := range []struct {
		n    uint64
		want bool
	}{
		{96, false},
		{69, true},
		{1, true},
		{10000, false},
	} {
		if got := Odd(tc.n); got != tc.want {
			t.Errorf("Odd: got %v want %v", got, tc.want)
		}
	}
}
func TestSetBit(t *testing.T) {
	for _, tc := range []struct {
		nm   uint64
		n    int
		want uint64
	}{
		{0, 0, 1},
		{1, 8, 257},
		{1024, 1, 1026},
		{1, 0, 1},
		{256, 8, 256},
		{0, 0, 0},
	} {
		if got := SetBit(tc.nm, tc.n); got != tc.want {
			t.Errorf("SetBit: got %v want %v", got, tc.want)
		}
	}
}
func TestClearBit(t *testing.T) {
	for _, tc := range []struct {
		nm   uint64
		n    int
		want uint64
	}{
		{1, 0, 0},
		{257, 8, 1},
		{1026, 1, 1024},
		{1, 1, 1},
		{256, 0, 256},
	} {
		if got := ClearBit(tc.nm, tc.n); got != tc.want {
			t.Errorf("ClearBit: got %v want %v", got, tc.want)
		}
	}
}
func TestFlipBit(t *testing.T) {
	for _, tc := range []struct {
		nm   uint64
		n    int
		want uint64
	}{
		{1, 0, 0},
		{258, 1, 256},
		{0, 0, 1},
		{256, 1, 258},
	} {
		if got := FlipBit(tc.nm, tc.n); got != tc.want {
			t.Errorf("TestFlipBit: got %v want %v", got, tc.want)
		}
	}
}
func TestOnes(t *testing.T) {
	for _, tc := range []struct {
		n    uint64
		want int
	}{
		{15, 4},
		{240, 4},
		{0, 0},
		{0xf0, 4},
		{0xffffffffffffffff, 64},
		{0, 0},
	} {
		if got := Ones(tc.n); got != tc.want {
			t.Errorf("TestOnes: got %v want %v", got, tc.want)
		}
	}
}
func TestReverse(t *testing.T) {
	for _, tc := range []struct {
		n    uint64
		want uint64
	}{
		{0, 0},
		{0xff00000000000000, 0xff},
		{0xff, 0xff00000000000000},
		{13, 12682136550675316736},
		{12682136550675316736, 13},
	} {
		if got := Reverse(tc.n); got != tc.want {
			t.Errorf("Reverse(%v): got = %v, want = %v", tc.n, got, tc.want)
		}
	}
}
func TestLeadingZeros(t *testing.T) {
	for _, tc := range []struct {
		n    uint64
		want int
	}{
		{0, 64},
		{1, 63},
		{0x7fffffffffffffff, 1},
		{0x7000000000000000, 1},
		{0xffffffffffffffff, 0},
	} {
		if got := LeadingZeroes(tc.n); got != tc.want {
			t.Errorf("LeadingZeros: got %v want %v", got, tc.want)
		}
	}
}
func TestTrailingZeros(t *testing.T) {
	for _, tc := range []struct {
		n    uint64
		want int
	}{
		{0, 64},
		{1, 0},
		{2, 1},
		{0x8000000000000000, 63},
	} {
		if got := TrailingZeroes(tc.n); got != tc.want {
			t.Errorf("TrailingZeros: got %v want %v", got, tc.want)
		}
	}
}
