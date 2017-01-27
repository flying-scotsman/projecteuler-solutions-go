/*
This program will solve the following problem:

The palindromic number 595 is interesting because it can be written
as the sum of consecutive squares: 6^2 + 7^2 + 8^2 + 9^2 + 10^2 + 11^2 + 12^2.

There are exactly eleven palindromes below one-thousand that can be written
as consecutive square sums, and the sum of these palindromes is 4164.
Note that 1 = 0^2 + 1^2 has not been included as this problem is concerned with the squares of positive integers.

Find the sum of all the numbers less than 10^8 that are both palindromic and can be written as the sum of consecutive squares.
*/

/* IDEAS
- Palindrome detection - convert integer to a string and reverse it - there could be a better, numeric way
    - string(str[i]) gives you the ith element of a converted integer
- Procedure: sum squares until you reach the upper limit, then remove the bottom square and start again
    - Again there could be a more efficient way but this will work
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	// Compute the sum
	sum := palindromeSummer(100000000)
	fmt.Printf("The sum of all palindromic sums of consecutive squares up to 10^8 is %v. \n", sum)
}

// This function is a tedious way of finding all palindromic sums of distinct squares between 5 and the upper limit
// Since the palindromes are sums of squares we only have to sum up to approximately the square root of the upper limit divided by two
func palindromeSummer(n int) int {
	var sum int
	var listNum []int
	// Outer loop, every time the inner loop will have one fewer number to sum over
	for i := 1; float64(i) < math.Sqrt(float64(n)); i++ {
		var numToTest int
		numToTest += i * i
		for j := i + 1; float64(j) < math.Sqrt(float64(n)); j++ {
			numToTest += j * j
			// Exit loop if numToTest is larger than n
			if numToTest > n {
				break
			}
			// Convert the number to be tested into a string
			str := fmt.Sprint(numToTest)
			// Create a flag to test if str is a palindrome
			palindromeFlag := true
			if len(str) > 1 { // palindromeFlag is automatically true if numToTest is a single digit (1^2 + 2^2)
				for k := 0; k <= len(str)/2-1; k++ {
					if str[len(str)-k-1] != str[k] {
						palindromeFlag = false
						break
					}
				}
			}
			// If palindromeFlag is true and we don't already have that number in our list, we add numToTest to sum
			listFlag := true
			for _, listElement := range listNum {
				if numToTest == listElement {
					listFlag = false
					break
				}
			}
			if palindromeFlag && listFlag {
				sum += numToTest
				listNum = append(listNum, numToTest)
			}
		}
	}
	return sum
}
