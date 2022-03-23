package main

import (
	"math"
	"testing"
)

func TestEven(t *testing.T) {
	for _, tc := range []struct {
		number uint64
		want   bool
	}{
		{96, true},
		{69, false},
	} {
		got := Even(tc.number)
		if got != tc.want {
			t.Errorf("ERR: Even(%v): got = %v, want = %v", tc.number, got, tc.want)
		}
	}
}

func TestOdd(t *testing.T) {
	for _, tc := range []struct {
		number uint64
		want   bool
	}{
		{96, false},
		{69, true},
	} {
		got := Odd(tc.number)
		if got != tc.want {
			t.Errorf("ERR: Odd(%v): got = %v, want = %v", tc.number, got, tc.want)
		}
	}
}

func TestSetBit(t *testing.T) {
	for _, tc := range []struct {
		x    uint64
		n    int
		want uint64
	}{
		{0, 0, 1},
		{1, 8, 257},
		{1024, 1, 1026},
		{1, 0, 1},
		{256, 8, 256},
	} {
		got := SetBit(tc.x, tc.n)
		if got != tc.want {
			t.Errorf("ERR: SetBit(%v, %v): got = %v, want = %v", tc.x, tc.n, got, tc.want)
		}
	}
}

func TestClearBit(t *testing.T) {
	for _, tc := range []struct {
		x    uint64
		n    int
		want uint64
	}{
		{1, 0, 0},
		{257, 8, 1},
		{1026, 1, 1024},
		{1, 1, 1},
		{256, 0, 256},
	} {
		got := ClearBit(tc.x, tc.n)
		if got != tc.want {
			t.Errorf("ERR: ClearBit(%v, %v): got = %v, want = %v", tc.x, tc.n, got, tc.want)
		}
	}
}

func TestFlipBit(t *testing.T) {
	for _, tc := range []struct {
		x    uint64
		n    int
		want uint64
	}{
		{1, 0, 0},
		{0, 0, 1},
		{258, 1, 256},
		{256, 1, 258},
	} {
		got := FlipBit(tc.x, tc.n)
		if got != tc.want {
			t.Errorf("ERR: FlipBit(%v, %v): got = %v, want = %v", tc.x, tc.n, got, tc.want)
		}
	}
}

func TestOnes(t *testing.T) {
	for _, tc := range []struct {
		number uint64
		want   int
	}{
		{0, 0},
		{15, 4},
		{240, 4},
		{0xf0, 4},
		{0xffffffffffffffff, 64},
		{math.MaxUint64, 64},
	} {
		got := Ones(tc.number)
		if got != tc.want {
			t.Errorf("ERR: Ones(%v): got = %v, want = %v", tc.number, got, tc.want)
		}
	}
}

func TestReverse(t *testing.T) {
	for _, tc := range []struct {
		number uint64
		want   uint64
	}{
		{0, 0},
		{0xff00000000000000, 0xff},
		{0xff, 0xff00000000000000},
		{13, 12682136550675316736},
	} {
		got := Reverse(tc.number)
		if got != tc.want {
			t.Errorf("ERR: Reverse(%v): got = %v, want = %v", tc.number, got, tc.want)
		}
	}
}

func TestLeadingZeroes(t *testing.T) {
	for _, tc := range []struct {
		number uint64
		want   int
	}{
		{0, 64},
		{1, 63},
		{0x7fffffffffffffff, 1},
		{0x7000000000000000, 1},
		{0xffffffffffffffff, 0},
	} {
		got := LeadingZeroes(tc.number)
		if got != tc.want {
			t.Errorf("ERR: LeadingZeroes(%v): got = %v, want = %v", tc.number, got, tc.want)
		}
	}
}

func TestTrailingZeroes(t *testing.T) {
	for _, tc := range []struct {
		number uint64
		want   int
	}{
		{0, 64},
		{1, 0},
		{2, 1},
		{0x8000000000000000, 63},
	} {
		got := TrailingZeroes(tc.number)
		if got != tc.want {
			t.Errorf("ERR: TrailingZeroes(%v): got = %v, want = %v", tc.number, got, tc.want)
		}
	}
}
