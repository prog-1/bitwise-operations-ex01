package main

func Even(n uint64) bool {
	return n&1 == 0
}

func Odd(n uint64) bool {
	return n&1 == 1
}

func SetBit(x uint64, n int) uint64 {
	return (1 << n) | x
}

func ClearBit(x uint64, n int) uint64 {
	return x &^ (1 << n)
}

func FlipBit(x uint64, n int) uint64 {
	return x ^ (1 << n)
}

func Ones(n uint64) int {
	var count int
	for ; n != 0; n >>= 1 {
		if n&1 == 1 {
			count++
		}
	}
	return count
}

func Reverse(n uint64) uint64 {
	var result uint64
	for i := 0; i < 64; i, n = i+1, n>>1 {
		result <<= 1
		result += n & 1
	}
	return result
}

func LeadingZeroes(n uint64) int {
	var count int
	for i := 0; i < 64 && (n>>63)&1 == 0; i, n = i+1, n<<1 {
		count++
	}
	return count
}

func TrailingZeroes(n uint64) int {
	var count int
	for i := 0; i < 64 && n&1 == 0; i, n = i+1, n>>1 {
		count++
	}
	return count
}
