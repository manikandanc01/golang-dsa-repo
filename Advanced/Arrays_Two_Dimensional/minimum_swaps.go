package arraystwodimensional

/*
Problem Description
Given an array of integers A and an integer B, find and return the minimum number of swaps required to bring all the numbers less than or equal to B together.

Note: It is possible to swap any two elements, not necessarily consecutive.

Problem Constraints
1 <= length of the array <= 100000
-10^9 <= A[i], B <= 10^9

Input Format
The first argument given is the integer array A.
The second argument given is the integer B.

Output Format
Return the minimum number of swaps.

Example Input
Input 1:
 A = [1, 12, 10, 3, 14, 10, 5]
 B = 8

Input 2:
 A = [5, 17, 100, 11]
 B = 20

Example Output
Output 1:
 2

Output 2:
 1


Example Explanation
Explanation 1:
 A = [1, 12, 10, 3, 14, 10, 5]
 After swapping  12 and 3, A => [1, 3, 10, 12, 14, 10, 5].
 After swapping  the first occurence of 10 and 5, A => [1, 3, 5, 12, 14, 10, 10].
 Now, all elements less than or equal to 8 are together.

Explanation 2:
 A = [5, 17, 100, 11]
 After swapping 100 and 11, A => [5, 17, 11, 100].
 Now, all elements less than or equal to 20 are together.
*/

/**
 * @input A : Integer array
 * @input B : Integer
 *
 * @Output Integer
 */
import "math"

func solveMinimumSwaps(A []int, B int) int {
	/*
	   Approach 1:
	   - First count number of good element (<=B)
	   - This is the size of the subarray (k)
	   - Consider all possible subarrays of size k
	   - Iterate and figure out how many bad elements are there, we need that many swaps.
	   - Worst case tc will be O(n^2), Sc will be O(1). TLE

	   Approach 2:
	    - instead of iterate and figure out the count of bad elements, we can figure out using sliding window.
	    - Here window size will be count of good elements
	    - Tc will be O(n), Sc will be O(1)
	*/

	// Approach 2 : Code
	n := len(A)
	countOfGoodElements := 0
	minimumSwaps := math.MaxInt

	for i := 0; i < n; i++ {
		if A[i] <= B {
			countOfGoodElements += 1
		}
	}

	badElementsCount := 0
	for i := 0; i < countOfGoodElements; i++ {
		if A[i] > B {
			badElementsCount += 1
		}
	}

	if badElementsCount < minimumSwaps {
		minimumSwaps = badElementsCount
	}

	endIndex := countOfGoodElements
	startIndex := 0

	for endIndex < n {
		if A[startIndex] > B {
			badElementsCount--
		}

		if A[endIndex] > B {
			badElementsCount++
		}

		if badElementsCount < minimumSwaps {
			minimumSwaps = badElementsCount
		}

		endIndex++
		startIndex++
	}

	return minimumSwaps

	/*
	   // Approach 1: Code

	   	n := len(A)
	   	countOfGoodElements := 0
	   	minimumSwaps := math.MaxInt

	   	for i:=0; i<n; i++{
	   	    if A[i] <= B{
	   	        countOfGoodElements += 1
	   	    }
	   	}

	   	for start:=0; start < n-countOfGoodElements+1; start++{
	   	    badElementsCount := 0
	   	    end := start+countOfGoodElements-1

	   	    for i:=start; i<=end; i++{
	   	         if A[i] > B{
	   	             badElementsCount++
	   	         }
	   	    }

	   	    if badElementsCount < minimumSwaps{
	   	        minimumSwaps = badElementsCount
	   	    }
	   	}

	   	return minimumSwaps
	*/
}
