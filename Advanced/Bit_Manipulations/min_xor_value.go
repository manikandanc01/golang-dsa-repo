package bitmanipulations

/**
 * @input A : Integer array
 *
 * @Output Integer

Problem Description
Given an integer array A of N integers, find the pair of integers in the array which have minimum XOR value. Report the minimum XOR value.

Problem Constraints
2 <= length of the array <= 100000
0 <= A[i] <= 109

Input Format
First and only argument of input contains an integer array A.

Output Format
Return a single integer denoting minimum xor value.

Example Input
Input 1:
 A = [0, 2, 5, 7]
Input 2:
 A = [0, 4, 7, 9]

Example Output
Output 1:
 2
Output 2:
 3

Example Explanation
Explanation 1:
 0 xor 2 = 2
*/

import (
	"math"
	// "sort"
)

func minXorRecursion(A []int, msb int) int {
	if msb == -1 {
		return 0
	}
	if len(A) == 0 {
		return math.MaxInt
	}
	if len(A) == 1 {
		return math.MaxInt
	}
	if len(A) == 2 {
		return A[0] ^ A[1]
	}

	zeroBucket := make([]int, 0)
	oneBucket := make([]int, 0)

	for i := 0; i < len(A); i++ {
		if (A[i] & (1 << msb)) > 0 {
			oneBucket = append(oneBucket, A[i])
		} else {
			zeroBucket = append(zeroBucket, A[i])
		}
	}

	xorValue1 := minXorRecursion(zeroBucket, msb-1)
	xorValue2 := minXorRecursion(oneBucket, msb-1)

	if xorValue1 < xorValue2 {
		return xorValue1
	}

	return xorValue2
}
func findMinXor(A []int) int {
	// n := len(A)

	// Brute Force consider all pairs.
	// Take xor and return the minimum xor value.
	// Tc will be O(n*n), Sc will be O(1)
	// Time Limit will be exceeded

	// Optimal approach:
	// To get the minimum xor, the msb should be 0. If msb should be zero, either we have to select all the numbers with 0 or all the numbers with 1.
	// Tc will be -> O(n), for each recursive calls the loop will run for approx n times),
	// Sc will be O(n). To consider extra space for zero bucket and one bucket.
	// Each element is processed once per bit.
	// Total operations per element is ~32.
	// So Total work = n * 32 = n

	// We can also solve this problem using sorting.
	// First sort the entire array.
	// And find the xor of adjacent elements and return the minimum value.
	// In order to get the minimum xor elements, msb should be zero.
	// So when we are sorting we are keeping the elements with common prefix next to each other.
	// So those xor will always ebe minimal. They only differ in the suffix one or two bits.
	// Tc will be O(n log n), Sc will be O(1)

	/* Brute Force code

	   n := len(A)
	   minXorValue := math.MaxInt

	   for i:=0; i<n; i++{
	       for j:= i+1; j<n; j++{
	           currentXorValue := (A[i] ^ A[j])
	           if currentXorValue < minXorValue{
	               minXorValue = currentXorValue
	           }
	       }
	   }

	   return minXorValue
	*/

	/* Sorting Technique code

	   minXorValue := math.MaxInt
	   sort.Ints(A)
	   n := len(A)

	   for i:=0; i<n-1; i++{
	       currentXorValue := A[i] ^ A[i+1]
	       if currentXorValue < minXorValue{
	           minXorValue = currentXorValue
	       }
	   }

	   return minXorValue
	*/

	return minXorRecursion(A, 31)
}
