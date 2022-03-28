package main

import "testing"

func TestEven(t *testing.T) {
	for _, tc := range []struct {
		name string
		n    uint64
		want bool
	}{
		{"from README", 96, true},
		{"from README", 69, false},
		{"simple", 1, false},
		{"simple", 2, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Even(tc.n); got != tc.want {
				t.Errorf("Even(%v) = %v, want %v", tc.n, got, tc.want)
			}
		})
	}
}

func TestOdd(t *testing.T) {
	for _, tc := range []struct {
		name string
		n    uint64
		want bool
	}{
		{"from README", 96, false},
		{"from README", 69, true},
		{"simple", 1, true},
		{"simple", 2, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Odd(tc.n); got != tc.want {
				t.Errorf("Odd(%v) = %v, want %v", tc.n, got, tc.want)
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
		{"from README", 0, 0, 1},
		{"from README", 1, 8, 257},
		{"from README", 1024, 1, 1026},
		{"from README", 1, 0, 1},
		{"from README", 256, 8, 256},
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
		{"from README", 1, 0, 0},
		{"from README", 257, 8, 1},
		{"from README", 1026, 1, 1024},
		{"from README", 1, 1, 1},
		{"from README", 256, 0, 256},
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
		{"from README", 1, 0, 0},
		{"from README", 0, 0, 1},
		{"from README", 258, 1, 256},
		{"from README", 256, 1, 258},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := FlipBit(tc.x, tc.n); got != tc.want {
				t.Errorf("FlipBit(%v, %v) = %v, want %v", tc.x, tc.n, got, tc.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	for _, tc := range []struct {
		name string
		n    uint64
		want uint64
	}{
		{"from README", 0, 0},
		{"from README", 0xff00000000000000, 0xff},
		{"from README", 0xff, 0xff00000000000000},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := Reverse(tc.n); got != tc.want {
				t.Errorf("Reverse(%v) = %v, want %v", tc.n, got, tc.want)
			}
		})
	}
}

func TestTrailingZeroes(t *testing.T) {
	for _, tc := range []struct {
		name string
		n    uint64
		want int
	}{
		{"from README", 0, 64},
		{"from README", 1, 0},
		{"from README", 2, 1},
		{"from README", 0x8000000000000000, 63},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := TrailingZeroes(tc.n); got != tc.want {
				t.Errorf("TrailingZeroes(%v) = %v, want %v", tc.n, got, tc.want)
			}
		})
	}
}

func TestLeadingingZeroes(t *testing.T) {
	for _, tc := range []struct {
		name string
		n    uint64
		want int
	}{
		{"from README", 0, 64},
		{"from README", 1, 63},
		{"from README", 0x7fffffffffffffff, 1},
		{"from README", 0x7000000000000000, 1},
		{"from README", 0xffffffffffffffff, 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := LeadingZeroes(tc.n); got != tc.want {
				t.Errorf("LeadingZeroes(%v) = %v, want %v", tc.n, got, tc.want)
			}
		})
	}
}
