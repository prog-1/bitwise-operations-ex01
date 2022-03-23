package main

import "fmt"

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
	for i := 0; i < 64; i, x = i+1, x>>1 {
		if x&1 == 1 {
			cnt++
		}
	}
	return cnt
}

func Reverse(x uint64) (rev uint64) {
	for i := 0; i < 64; i++ {
		tmp := x & 1
		rev = rev << 1
		rev = rev + tmp
		x = x >> 1
	}
	return rev
}

func LeadingZeroes(x uint64) (cnt int) {
	var rev uint64
	for i := 0; i < 64; i++ {
		tmp := x & 1
		rev = rev << 1
		rev = rev + tmp
		x = x >> 1
	}
	for rev = ^rev; rev&1 == 1; rev = rev >> 1 {
		cnt++
	}
	return cnt
}

func TrailingZeroes(x uint64) (cnt int) {
	for x = ^x; x&1 == 1; x = x >> 1 {
		cnt++
	}
	return cnt
}

func main() {
	fmt.Println(Even(8))
	fmt.Println(Odd(10))
}
