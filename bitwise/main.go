package main

import "fmt"

func Even(a uint64) bool {
	c := a & 3
	if c == 1 || c == 3 {
		return false
	}
	return true
}
func Odd(a uint64) bool {
	c := a & 3
	if c == 0 || c == 2 {
		return false
	}
	return true
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
func main() {
	fmt.Println(Even(69))
	fmt.Println(Odd(69))
	fmt.Println(SetBit(1, 8))
	fmt.Println(ClearBit(1026, 1))
	fmt.Println(FlipBit(258, 1))
}
