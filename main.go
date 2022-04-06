package main

import (
	"fmt"
)

func Even(n uint64) bool {
	// return n%2 == 0
	return n&1 == 0
}

func Odd(n uint64) bool {
	// return n%2 != 0
	return n&1 != 0 // or == 1
}

func SetBit(x uint64, n int) uint64 {
	return (1 << n) | x
}

func ClearBit(x uint64, n int) uint64 {
	return x &^ (1 << n)
}

func FlipBit(x uint64, n int) uint64 {
	return (1 << n) ^ x
}

func Ones(n uint64) int {
	var count int
	for i := 0; i < 64; i++ {
		if n&1 == 1 {
			count++
		}
		n = n >> 1
	}
	return count
}

func Reverse(n uint64) uint64 {
	var reversed uint64
	for i := 0; i < 64; i, n = i+1, n>>1 {
		reversed = reversed << 1
		reversed = reversed + n&1
	}
	return reversed
}

func LeadingZeroes(n uint64) int {
	var count int
	for i := 0; i < 64; i, n = i+1, n<<1 {
		if (n>>63)&1 == 1 {
			break
		}
		count++
	}
	return count
}

func TrailingZeroes(n uint64) int {
	var count int
	for i := 0; i < 64; i, n = i+1, n>>1 {
		if n&1 == 1 {
			break
		}
		count++
	}
	return count
}

func main() {
	fmt.Printf("Even:\n %v (%08b) - %v\n %v (%08b) - %v\n", 96, 96, Even(96), 69, 69, Even(69))
	fmt.Printf("Odd:\n %v (%08b) - %v\n %v (%08b) - %v\n", 96, 96, Odd(96), 69, 69, Odd(69))
	fmt.Printf("SetBit:\n x = %v (%08b), n = %v  --  %v\n x = %v (%08b), n = %v  --  %v\n x = %v (%08b), n = %v  --  %v\n x = %v (%08b), n = %v  --  %v\n x = %v (%08b), n = %v  --  %v\n", 0, 0, 0, SetBit(0, 0), 1, 1, 8, SetBit(1, 8), 1024, 1024, 1, SetBit(1024, 1), 1, 1, 0, SetBit(1, 0), 256, 256, 8, SetBit(256, 8))
	fmt.Printf("ClearBit:\n x = %v (%08b), n = %v  --  %v\n x = %v (%08b), n = %v  --  %v\n x = %v (%08b), n = %v  --  %v\n x = %v (%08b), n = %v  --  %v\n x = %v (%08b), n = %v  --  %v\n", 1, 1, 0, ClearBit(1, 0), 257, 257, 8, ClearBit(257, 8), 1026, 1026, 1, ClearBit(1026, 1), 1, 1, 1, ClearBit(1, 1), 256, 256, 0, ClearBit(256, 0))
	fmt.Printf("FlipBit:\n x = %v (%08b), n = %v  --  %v\n x = %v (%08b), n = %v  --  %v\n x = %v (%08b), n = %v  --  %v\n x = %v (%08b), n = %v  --  %v\n", 1, 1, 0, FlipBit(1, 0), 0, 0, 0, FlipBit(0, 0), 258, 258, 1, FlipBit(258, 1), 256, 256, 1, FlipBit(256, 1))
	fmt.Printf("Ones:\n %v (%08b) - %v\n %v (%08b) - %v\n", 0, 0, Ones(0), 15, 15, Ones(15))
	fmt.Printf("Reverse:\n %v (%08b) - %v\n %v (%08b) - %v\n", 0, 0, Reverse(0), 13, 13, Reverse(13))
	fmt.Printf("LeadingZeroes:\n %v (%08b) - %v\n %v (%08b) - %v\n", 0, 0, LeadingZeroes(0), 1, 1, LeadingZeroes(1))
	fmt.Printf("TrailingZeroes:\n %v (%08b) - %v\n %v (%08b) - %v\n", 0, 0, TrailingZeroes(0), 1, 1, TrailingZeroes(1))
}
