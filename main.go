package main

import "fmt"

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
func Ones(n uint64) uint64 {
	var p, cnt uint64
	for ; n >= (1 << p); p++ {
		if (1 << p) == n&(1<<p) {
			cnt++
		}
	}

	return cnt
}
func Reverse(n uint64) (res uint64) {
	for i := 0; i < 64; i++ {
		res = res<<1 | n&1
		n = n >> 1
	}
	return res
}
func LeadingZeroes(n uint64) int {
	return TrailingZeroes(Reverse(n))
}
func TrailingZeroes(n uint64) int {
	var p, cnt uint64
	for ; n >= (1 << p); p++ {
		if (1 << p) == n&(1<<p) {
			return int(cnt)
		}
		cnt++
	}

	return 0
}

func main() {
	fmt.Println(Even(58))
	fmt.Println(Odd(125))
	fmt.Println(SetBit(1024, 1))
	fmt.Println(ClearBit(1023, 0))
	fmt.Println(FlipBit(657, 0))
	fmt.Println(Ones(69))
	fmt.Println(Reverse(13))
	fmt.Println(LeadingZeroes(1))
	fmt.Println(TrailingZeroes(100000003200))
}
