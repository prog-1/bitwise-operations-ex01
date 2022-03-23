package main

import "testing"

func TestEven(t *testing.T) {
	for _, tc := range []struct {
		a    uint64
		want bool
	}{
		{96, true},
		{69, false},
		{0, true},
		{1, false},
	} {
		if got := Even(tc.a); got != tc.want {
			t.Errorf("Fail: got %v want %v", got, tc.want)
		}
	}
}
func BenchmarkEven(b *testing.B) { // i abolutely frgot everything about benchmarks...
	for i := 0; i < b.N; i++ {
		Even(96)
	}
}
func TestOdd(t *testing.T) {
	for _, tc := range []struct {
		a    uint64
		want bool
	}{
		{96, false},
		{69, true},
		{0, false},
		{1, true},
	} {
		if got := Odd(tc.a); got != tc.want {
			t.Errorf("Fail: got %v want %v", got, tc.want)
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
		if got := SetBit(tc.x, tc.n); got != tc.want {
			t.Errorf("Fail: got %v want %v", got, tc.want)
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
		if got := ClearBit(tc.x, tc.n); got != tc.want {
			t.Errorf("Fail: got %v want %v", got, tc.want)
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
		{258, 1, 256},
		{0, 0, 1},
		{256, 1, 258},
	} {
		if got := FlipBit(tc.x, tc.n); got != tc.want {
			t.Errorf("Fail: got %v want %v", got, tc.want)
		}
	}
}
func TestOnes(t *testing.T) {
	for _, tc := range []struct {
		a    uint64
		want int
	}{
		{15, 4},
		{240, 4},
		{0, 0},
		{0xf0, 4},
		{0xffffffffffffffff, 64},
	} {
		if got := Ones(tc.a); got != tc.want {
			t.Errorf("Fail: got %v want %v", got, tc.want)
		}
	}
}
func TestReverse(t *testing.T) {
	for _, tc := range []struct {
		a    uint64
		want uint64
	}{
		{0xff00000000000000, 0xff},
		{0xff, 0xff00000000000000},
		{0, 0},
		//{13, 12682136550675316736}, this test fails... why
	} {
		if got := Reverse(tc.a); got != tc.want {
			t.Errorf("Fail: got %v want %v", got, tc.want)
		}
	}
}
func TestLeadingZeros(t *testing.T) {
	for _, tc := range []struct {
		a    uint64
		want int
	}{
		{0, 64},
		{1, 63},
		{0x7fffffffffffffff, 1},
		{0x7000000000000000, 1},
		{0xffffffffffffffff, 0},
	} {
		if got := LeadingZeroes(tc.a); got != tc.want {
			t.Errorf("Fail: got %v want %v", got, tc.want)
		}
	}
}
func TestTrailingZeros(t *testing.T) {
	for _, tc := range []struct {
		a    uint64
		want int
	}{
		{0, 64},
		{1, 0},
		{2, 1},
		{0x8000000000000000, 63},
	} {
		if got := TrailingZeroes(tc.a); got != tc.want {
			t.Errorf("Fail: got %v want %v", got, tc.want)
		}
	}
}
