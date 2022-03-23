package main

import (
	"fmt"
)

func Even(a uint64) bool {
	a = a << 63
	return a == 0
}
func Odd(a uint64) bool {
	return !Even(a)
}

func SetBit(x uint64, n int) uint64 {
	return x | (1 << n)
}
func ClearBit(x uint64, n int) uint64 {
	return ^((^x) | (1 << n))
}
func FlipBit(x uint64, n int) uint64 {
	return x ^ (1 << n)
}
func Ones(x uint64) int {
	var res int
	for x != 0 {
		res += int(x & 1)
		x = x >> 1
	}
	return res
}

func Reverse(n uint64) (r uint64) {
	for i := 64; i != 0; n, i = n>>1, i-1 {
		r = r << 1
		r += n & 1
	}
	return
}

func LeadingZeroes(n uint64) (i int) {
	if n == 0 {
		return 64
	}
	for ; (n&(1<<63))>>63 != 1; n = n << 1 {
		i++
	}
	return
}

func TrailingZeroes(n uint64) (i int) {
	if n == 0 {
		return 64
	}
	for ; n&1 != 1; n = n >> 1 {
		i++
	}
	return
}

func main() {
	fmt.Printf("%b", Reverse(255))
}
