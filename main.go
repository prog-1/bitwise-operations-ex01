package main

func Even(n uint64) bool {
	if n&1 == 0 {
		return true
	}
	return false

}

func Odd(n uint64) bool {
	if n&1 != 0 {
		return true
	}
	return false
}

func SetBit(x uint64, n int) (result uint64) {
	result = (1 << n) | x
	return result
}

func ClearBit(x uint64, n int) (result uint64) {
	result = x &^ (1 << n)
	return result
}

func FlipBit(x uint64, n int) (result uint64) {
	result = x ^ (1 << n)
	return result
}

func Ones(n uint64) (result int) {
	for ; n != 0; n >>= 1 {
		if n == 1 {
			result++
		}
	}
	return result
}

func Reverse(n uint64) (result uint64) {
	for i := 0; i < 64; i++ {
		result = result<<1 | n&1
		n = n >> 1
	}
	return result
}

func LeadingZeroes(n uint64) int {
	return TrailingZeroes(Reverse(n))
}

func TrailingZeroes(n uint64) (result int) {
	if n == 0 {
		return 64
	}
	for ; n&1 != 1; n = n >> 1 {
		result++
	}
	return result
}
