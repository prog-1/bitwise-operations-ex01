# Bitwise Operations

Implement the following functions using bitwise operations (and sometimes counters):

1. `func Even(uint64) bool` returns is a number is even.
       
    ```
    Even(96) -> true
    Even(69) -> false
    ```
    
2. `func Odd(uint64) bool` returns if a number is odd.
    
    ```
    Odd(96) -> false
    Odd(69) -> true
    ```

3. `func SetBit(x uint64, n int) uint64` sets the `n`th bit in the number `x`. The bits are enumerated from `0` from the right.

    ```
    SetBit(0, 0) -> 1
    SetBit(1, 8) -> 257
    SetBit(1024, 1) -> 1026
    SetBit(1, 0) -> 1
    SetBit(256, 8) -> 256
    ```
    
4. `func ClearBit(x uint64, n int) uint64` clears the `n`th bit in the number `x`. The bits are enumerated from `0` from the right.

   ```
   ClearBit(1, 0) -> 0
   ClearBit(257, 8) -> 1
   ClearBit(1026, 1) -> 1024
   ClearBit(1, 1) -> 1
   ClearBit(256, 0) -> 256
   ```
   
5. `func FlipBit(x uint64, n int) uint64` flips the `n`th bit in the number `x`.  The bits are enumerated from `0` from the right.
 
   ```
   FlipBit(1, 0) -> 0
   FlipBit(0, 0) -> 1
   FlipBit(258, 1) -> 256
   FlipBit(256, 1) -> 258
   ```
   
6. `func Ones(uint64) int` returns the count of `1` bits in a number.

   ```
   Ones(0) -> 0
   Ones(15) -> 4
   Ones(240) -> 4
   Ones(0xf0) -> 4
   Ones(0xffffffffffffffff) -> 64
   Ones(math.MaxUint64) -> 64
   ```
   
7. `func Reverse(uint64) uint64` returns the number in a reversed bit order.

   ```
   Reverse(0) -> 0
   Reverse(0xff00000000000000) -> 0xff
   Reverse(0xff) -> 0xff00000000000000
   Reverse(0b1101) -> 0b1011000...00000
   Reverse(13) -> 936748722493063168
   ```
   
8. `func LeadingZeroes(uint64) int` returns the count of `0` preceding `1` in a number from the left.

   ```
   LeadingZeroes(0) -> 64
   LeadingZeroes(1) -> 63
   LeadingZeroes(0x7fffffffffffffff) -> 1
   LeadingZeroes(0x7000000000000000) -> 1
   LeadingZeroes(0xffffffffffffffff) -> 0
   ```
  
9. `func TrailingZeroes(uint64) int` retunrns the count of `0` preceding `1` in a number from the right.

   ```
   TrailingZeroes(0) -> 64
   TrailingZeroes(1) -> 0
   TrailingZeroes(2) -> 1
   TrailingZeroes(0x8000000000000000) -> 63
   ```

Implement tests for all functions.
