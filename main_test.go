package main

import (
	"testing"
)

func TestEven(t *testing.T) {
	for _, tc := range []struct {
		input uint64
		want  bool
	}{
		{312, true},
		{321, false},
	} {
		t.Run("", func(t *testing.T) {
			if got := Even(tc.input); got != tc.want {
				t.Errorf("Even(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestOdd(t *testing.T) {
	for _, tc := range []struct {
		input uint64
		want  bool
	}{
		{64, false},
		{11, true},
	} {
		t.Run("", func(t *testing.T) {
			if got := Odd(tc.input); got != tc.want {
				t.Errorf("Odd(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestSetBit(t *testing.T) {
	for _, tc := range []struct {
		input  uint64
		input2 int
		want   uint64
	}{
		{0, 0, 1},
		{1, 8, 257},
		{1024, 1, 1026},
		{1, 0, 1},
		{256, 8, 256},
	} {
		t.Run("", func(t *testing.T) {
			if got := SetBit(tc.input, tc.input2); got != tc.want {
				t.Errorf("SetBit(%v, %v) = %v, want %v", tc.input, tc.input2, got, tc.want)
			}
		})
	}
}

func TestClearBit(t *testing.T) {
	for _, tc := range []struct {
		input  uint64
		input2 int
		want   uint64
	}{
		{1, 0, 0},
		{257, 8, 1},
		{1026, 1, 1024},
		{1, 1, 1},
		{256, 0, 256},
	} {
		t.Run("", func(t *testing.T) {
			if got := ClearBit(tc.input, tc.input2); got != tc.want {
				t.Errorf("ClearBit(%v, %v) = %v, want %v", tc.input, tc.input2, got, tc.want)
			}
		})
	}
}

func TestFlipBit(t *testing.T) {
	for _, tc := range []struct {
		input  uint64
		input2 int
		want   uint64
	}{
		{1, 0, 0},
		{0, 0, 1},
		{258, 1, 256},
		{256, 1, 258},
	} {
		t.Run("", func(t *testing.T) {
			if got := FlipBit(tc.input, tc.input2); got != tc.want {
				t.Errorf("FlipBit(%v, %v) = %v, want %v", tc.input, tc.input2, got, tc.want)
			}
		})
	}
}

func TestOnes(t *testing.T) {
	for _, tc := range []struct {
		input uint64
		want  int
	}{
		{0, 0},
		{15, 4},
		{240, 4},
		{0xf0, 4},
		{0xffffffffffffffff, 64},
		{18446744073709551615, 64},
	} {
		t.Run("", func(t *testing.T) {
			if got := Ones(tc.input); got != tc.want {
				t.Errorf("Ones(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	for _, tc := range []struct {
		input uint64
		want  uint64
	}{
		{0, 0},
		{0xff00000000000000, 0xff},
		{0xff, 0xff00000000000000},
		{0b1101, 0b1011000000000000000000000000000000000000000000000000000000000000},
		{13, 12682136550675316736},
	} {
		t.Run("", func(t *testing.T) {
			if got := Reverse(tc.input); got != tc.want {
				t.Errorf("Reverse(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestLeadingZeroes(t *testing.T) {
	for _, tc := range []struct {
		input uint64
		want  int
	}{
		{0, 64},
		{1, 63},
		{0x7fffffffffffffff, 1},
		{0x7000000000000000, 1},
		{0xffffffffffffffff, 0},
	} {
		t.Run("", func(t *testing.T) {
			if got := LeadingZeroes(tc.input); got != tc.want {
				t.Errorf("LeadingZeroes(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestTrailingZeroes(t *testing.T) {
	for _, tc := range []struct {
		input uint64
		want  int
	}{
		{0, 64},
		{1, 0},
		{0x8000000000000000, 63},
	} {
		t.Run("", func(t *testing.T) {
			if got := TrailingZeroes(tc.input); got != tc.want {
				t.Errorf("TrailingZeroes(%v) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}
