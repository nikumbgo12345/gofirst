package main

import "fmt"

func main() {
	fmt.Println(addBinary("1010", "1011"))
}

func addBinary(a string, b string) string {

	if len(a) < len(b) {
		a, b = b, a
	}
	var c string
	remainder := false
	index2 := len(b) - 1
	for index1 := len(a) - 1; index1 >= 0; index1-- {
		if index2 < 0 {
			if remainder == false {
				c = a[index1:index1+1] + c
			} else {
				if a[index1:index1+1] == "0" {
					c = "1" + c
					remainder = false
				} else {
					c = "0" + c

				}
			}
		} else {

			if a[index1:index1+1] == b[index2:index2+1] {
				if a[index1:index1+1] == "1" {
					if remainder == false {
						remainder = true
						c = "0" + c
					} else {
						c = "1" + c
					}
				} else {
					if remainder == true {
						c = "1" + c
						remainder = false

					} else {
						c = "0" + c
					}
				}

			} else {
				if remainder {
					c = "0" + c
				} else {
					c = "1" + c
				}
			}
			index2--
		}
	}
	if remainder {
		c = "1" + c
	}

	return c
}

/*
Given two binary strings a and b, return their sum as a binary string.



Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"


Constraints:

1 <= a.length, b.length <= 104
a and b consist only of '0' or '1' characters.
Each string does not contain leading zeros except for the zero itself.*/
