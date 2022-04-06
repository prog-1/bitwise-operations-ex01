package main

import (
	"math"
	"testing"
)

func TestEven(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input uint64
		want  bool
	}{
		{"1", 96, true},
		{"2", 69, false},
		{"3", 1, false},
		{"4", 2, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Even(tc.input); got != tc.want {
				t.Errorf("Even(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestOdd(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input uint64
		want  bool
	}{
		{"1", 96, false},
		{"2", 69, true},
		{"3", 1, true},
		{"4", 2, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Odd(tc.input); got != tc.want {
				t.Errorf("Odd(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestSetBit(t *testing.T) {
	for _, tc := range []struct {
		name string
		x    uint64
		n    int
		want uint64
	}{
		{"1", 0, 0, 1},
		{"2", 1, 8, 257},
		{"3", 1024, 1, 1026},
		{"4", 1, 0, 1},
		{"5", 256, 8, 256},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := SetBit(tc.x, tc.n); got != tc.want {
				t.Errorf("SetBit(%v, %v) = %v, want %v", tc.x, tc.n, got, tc.want)
			}
		})
	}
}

func TestClearBit(t *testing.T) {
	for _, tc := range []struct {
		name string
		x    uint64
		n    int
		want uint64
	}{
		{"1", 1, 0, 0},
		{"2", 257, 8, 1},
		{"3", 1026, 1, 1024},
		{"4", 1, 1, 1},
		{"5", 256, 0, 256},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := ClearBit(tc.x, tc.n); got != tc.want {
				t.Errorf("ClearBit(%v, %v) = %v, want %v", tc.x, tc.n, got, tc.want)
			}
		})
	}
}

func TestFlipBit(t *testing.T) {
	for _, tc := range []struct {
		name string
		x    uint64
		n    int
		want uint64
	}{
		{"1", 1, 0, 0},
		{"2", 0, 0, 1},
		{"3", 258, 1, 256},
		{"4", 256, 1, 258},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := FlipBit(tc.x, tc.n); got != tc.want {
				t.Errorf("FlipBit(%v, %v) = %v, want %v", tc.x, tc.n, got, tc.want)
			}
		})
	}
}

func TestOnes(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input uint64
		want  int
	}{
		{"1", 0, 0},
		{"2", 15, 4},
		{"3", 240, 4},
		{"4", 0xf0, 4},
		{"5", 0xffffffffffffffff, 64},
		{"6", math.MaxUint64, 64},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Ones(tc.input); got != tc.want {
				t.Errorf("Ones(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input uint64
		want  uint64
	}{
		{"1", 0, 0},
		{"2", 0xff00000000000000, 0xff},
		{"3", 0xff, 0xff00000000000000},
		{"4", 0b1101, 0b1011000000000000000000000000000000000000000000000000000000000000},
		{"5", 13, 12682136550675316736},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Reverse(tc.input); got != tc.want {
				t.Errorf("Reverse(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestLeadingZeroes(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input uint64
		want  int
	}{
		{"1", 0, 64},
		{"2", 1, 63},
		{"3", 0x7fffffffffffffff, 1},
		{"4", 0x7000000000000000, 1},
		{"5", 0xffffffffffffffff, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := LeadingZeroes(tc.input); got != tc.want {
				t.Errorf("LeadingZeroes(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestTrailingZeroes(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input uint64
		want  int
	}{
		{"1", 0, 64},
		{"2", 1, 0},
		{"3", 2, 1},
		{"4", 0x8000000000000000, 63},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := TrailingZeroes(tc.input); got != tc.want {
				t.Errorf("TrailingZeroes(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
