package main

import (
	"fmt"
	"math/bits"
)

func onevar() (a uint64) {
	fmt.Println("Enter number: ")
	fmt.Scan(&a)
	return a
}
func twovar() (x uint64, n int) {
	fmt.Println("Enter number: ")
	fmt.Scan(&x)
	fmt.Println("Enter n: ")
	fmt.Scan(&n)
	return x, n
}
func UI() {
	for {
		fmt.Println("Choose program: 1,2,3,4,5,6,7,8,9 or 10* else - exit ")
		var action int
		fmt.Scan(&action)
		if action == 1 {
			fmt.Println(Even(onevar()))
		} else if action == 2 {
			fmt.Println(Odd(onevar()))
		} else if action == 3 {
			fmt.Println(SetBit(twovar()))
		} else if action == 4 {
			fmt.Println(ClearBit(twovar()))
		} else if action == 5 {
			fmt.Println(FlipBit(twovar()))
		} else if action == 6 {
			fmt.Println(Ones(onevar()))
		} else if action == 7 {
			fmt.Println(Reverse(onevar()))
		} else if action == 8 {
			fmt.Println(LeadingZeroes(onevar()))
		} else if action == 9 {
			fmt.Println(TrailingZeroes(onevar()))
		} else if action == 10 {
			Powerof2()
		} else {
			break
		}
	}
}
func Even(a uint64) bool {
	return a&1 == 0
}
func Odd(a uint64) bool {
	return a&1 != 0
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
func Ones(a uint64) (count int) {
	for a != 0 {
		if (a & 1) == 1 {
			count++
		}
		a >>= 1
	}
	return count
	//bits.OnesCount64(a)
}
func Reverse(a uint64) uint64 {
	return bits.Reverse64(a)
}
func LeadingZeroes(a uint64) int { // i dont really understood how the output supposed to be...
	return bits.LeadingZeros64(a)
}
func TrailingZeroes(a uint64) int {
	return bits.TrailingZeros64(a)
}
func Powerof2() {
	fmt.Println("The program determines if a number is a power of 2.")
	var n, i int
	fmt.Print("Enter the number: ")
	fmt.Scan(&n)
	if n > 0 && (n&(n-1)) == 0 {
		for z := 1; z != n; z *= 2 {
			i++
		}
		fmt.Printf("%v = 2^%v\n", n, i)
	} else {
		fmt.Println(n, "is not a power of 2.")
	}
}
func main() {
	UI()
}
