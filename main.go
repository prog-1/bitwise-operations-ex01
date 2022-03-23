package main

func Even(n uint64) bool {
	return n&1 == 0
}
func Odd(n uint64) bool {
	return n&1 != 0
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

func main() {
}
