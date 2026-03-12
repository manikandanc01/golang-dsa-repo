package arraysonedimensional

/*
Problem Description
Given an array A of length N, your task is to find the maximum possible sum of any non-empty contiguous subarray.
In other words, among all possible subarrays of A, determine the one that yields the highest sum and return that sum.

Problem Constraints
1 <= N <= 106
-1000 <= A[i] <= 1000

Input Format
The first and the only argument contains an integer array, A.

Output Format
Return an integer representing the maximum possible sum of the contiguous subarray.

Example Input
Input 1:
 A = [1, 2, 3, 4, -10]

Input 2:
 A = [-2, 1, -3, 4, -1, 2, 1, -5, 4]


Example Output
Output 1:
 10
Output 2:
 6

Example Explanation:
Explanation 1:
 The subarray [1, 2, 3, 4] has the maximum possible sum of 10.
Explanation 2:
 The subarray [4,-1,2,1] has the maximum possible sum of 6.
*/

/**
 * @input A : Integer array
 *
 * @Output Integer
 */
func maxSubArray(A []int) int {
	/*
	   Brute force is considering all subarrays. outer loop start index
	   inner loop end index, third loop iterate from start index to end index and calculate sum. Finally if sum is greater
	   than the ans, store. TC: O(n3) , SC(O(1)).
	   Better approach: Using prefix sum or carry forward to calculate the sum.
	   Efficient approach is Kadane's algorithm. Based on the fact that negative sum will not contribute overall maximum value.

	   Optimized is Kadane's algorithm. Negative value won't contribute to maximum sum at any cost.
	   [-2,4], 4 is greater than -2+4 = 2. [-2,1] 1is greater than -2+1 = -1. So whenever our sum is less than 0, we reset the sum to 0.
	   Time complexity is O(n), Sc is O(1)
	*/

	n := len(A)
	sum := 0
	maxSum := A[0]

	for i := 0; i < n; i++ {
		sum += A[i]
		if sum > maxSum {
			maxSum = sum
		}

		if sum < 0 {
			sum = 0
		}
	}

	return maxSum
}
