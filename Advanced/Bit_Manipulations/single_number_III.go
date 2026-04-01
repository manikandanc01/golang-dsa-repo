package bitmanipulations

/*
Problem Description
Given an array of positive integers A, two integers appear only once, and all the other integers appear twice.
Find the two integers that appear only once.

Note: Return the two numbers in ascending order.

Problem Constraints
2 <= |A| <= 100000
1 <= A[i] <= 109

Input Format
The first argument is an array of integers of size N.

Output Format
Return an array of two integers that appear only once.

Example Input
Input 1:
A = [1, 2, 3, 1, 2, 4]

Input 2:
A = [1, 2]

Example Output
Output 1:
[3, 4]

Output 2:
[1, 2]

Example Explanation
Explanation 1:
3 and 4 appear only once.

Explanation 2:
1 and 2 appear only once.
*/

/**
 * @input A : Integer array
 *
 * @Output Integer array.
 */

// import "sort"
func solve(A []int) []int {
	// Brute Force:
	// Given: 2 numbers appears only once
	// Iterate over the array, check each element is appears only once.
	// If no add that element to the answer.
	// Tc will be O(n^2), Sc will be O(1) -> not considering the output space., TLE

	// Better Approach:
	// Use hashmap to store the frequency of each element.
	// Then iterate over the map, if the frequency is equal to 1, add that to answer.
	// Tc will be O(n), Sc will be O(n), final array size is always 2, so sorting doesn't matter for TC.

	// Best approach:
	// Use the xor behaviour.
	// XOR of same elements is 0.
	// So in the given elements, only 2 elements appears only once and other elements will be appearing twice.
	// So final xor will be the xor of that 2 unique elements.
	// But we need to find the exact elements.
	// So we use bucketing.
	// [1,2,3,1,2,4] => final answer will be the xor of [3,4] => 7 [0011, 0100] => [0111]
	// We have to find the first set bit of the xor result from right.
	// Let's say if the 0th bit of answer is set, which means the 0th bit of one element is 0 and 0th bit of another element is 1.
	// Then only we will get the 1 in the final answer.
	// So we will group the element which have the 0th bit set as one group, and another group with 0th bit not set.
	// 0th bit not set group => [2, 2, 4]  => xor will give => 4
	// 0th bit set group => [1 1 3] => xor will give => 3
	// This is how we can find those 2 unique elements.
	// Tc will be O(n), Sc will be O(1)

	n := len(A)
	xorValue := 0

	for i := 0; i < n; i++ {
		xorValue ^= A[i]
	}

	firstSetBit := 0
	for (xorValue & (1 << firstSetBit)) == 0 {
		firstSetBit++
	}

	setBitElement, notSetBitElement := 0, 0

	for i := 0; i < n; i++ {
		if (A[i] & (1 << firstSetBit)) > 0 {
			setBitElement ^= A[i]
		} else {
			notSetBitElement ^= A[i]
		}
	}

	if setBitElement < notSetBitElement {
		return []int{setBitElement, notSetBitElement}
	}

	return []int{notSetBitElement, setBitElement}

	/* Better Approach code

	   n := len(A)
	   mp := make(map[int]int)
	   sol := make([]int,0)

	   for i:=0; i<n; i++{
	       mp[A[i]]++
	   }

	   for key, freq := range mp{
	       if freq == 1{
	            sol = append(sol, key)
	       }
	   }

	   sort.Ints(sol)
	   return sol


	*/

	/* Brute Force code
	   ans := make([]int,0)
	   for i:=0; i<len(A); i++{
	       isAppearOnce := true
	       for j:=0; j<len(A); j++{
	           if j == i{
	               continue
	           }

	           if A[j] == A[i]{
	               isAppearOnce = false
	               break
	           }
	       }

	       if isAppearOnce{
	           ans = append(ans, A[i])
	       }
	   }

	   sort.Ints(ans)
	   return ans

	*/

}
