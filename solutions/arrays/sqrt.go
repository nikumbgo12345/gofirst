package main

import "fmt"

func main() {
	fmt.Println("test")
	fmt.Println(mySqrt(9))
}
func mySqrt(x int) int {
	if x == 1 {
		return 1
	}
	if x <= 4 {
		return x / 2
	}
	mid := x / 2

	for mid*mid > x {
		mid = mid / 2
	}
	value := mid
	for value*value <= x {
		value++
	}
	return (value - 1)
}

/**
Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.

You must not use any built-in exponent function or operator.

For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.


Example 1:

Input: x = 4
Output: 2
Explanation: The square root of 4 is 2, so we return 2.
Example 2:

Input: x = 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since we round it down to the nearest integer, 2 is returned.
*/
