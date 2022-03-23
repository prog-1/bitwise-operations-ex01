package main

import (
	"fmt"
)

func Even(x uint64) bool {
	return x&1 == 0
}

func Odd(x uint64) bool {
	return x&1 != 0
}

func SetBit(x uint64, n int) uint64 {
	return x | (1 << n)
}

func ClearBit(x uint64, n int) uint64 {
	return x &^ (1 << n)
}

func FlipBit(x uint64, n int) uint64 {
	return x ^ (1 << n)
}

func Ones(x uint64) (cnt int) {
	for n := x; n > 0; n = n >> 1 {
		if n&1 == 1 {
			cnt++
		}
	}
	return cnt
}

func Reverse(x uint64) (result uint64) {
	for i := 0; i < 64; i++ {
		result = result<<1 | x&1
		x = x >> 1
	}
	return result
}

func LeadingZeroes(x uint64) int {
	return TrailingZeroes(Reverse(x))
}

func TrailingZeroes(x uint64) (cnt int) {
	if x == 0 {
		return 64
	}
	for ; x&1 != 1; x = x >> 1 {
		cnt++
	}
	return cnt
}

func main() {
	fmt.Println(Even(42))
	fmt.Println(Odd(321))
	fmt.Println(SetBit(1024, 1))
	fmt.Println(ClearBit(1026, 1))
	fmt.Println(FlipBit(658, 1))
	fmt.Println(Ones(124))
	fmt.Println(Reverse(0b1101))
	fmt.Println(LeadingZeroes(1))
	fmt.Println(TrailingZeroes(2))
}
