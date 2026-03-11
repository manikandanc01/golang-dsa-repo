package subarraysandsubsequences

/*
Problem Description:

Given an integer array, A of size N.
You have to find all possible non-empty subsequences of the array of numbers and then,
for each subsequence, find the difference between the largest and smallest number in that subsequence.
Then add up all the differences to get the number.

As the number may be large, output the number modulo 1e9 + 7 (1000000007).

NOTE: Subsequence can be non-contiguous.

Problem Constraints
1 <= N <= 10000
1<= A[i] <=1000

Input Format
First argument is an integer array A.

Output Format
Return an integer denoting the output.

Example Input
Input 1:
A = [1, 2]

Input 2:
A = [3, 5, 10]

Example Output
Output 1:
 1

Output 2:
 21

Example Explanation
Explanation 1:

All possible non-empty subsets are:
[1]     largest-smallest = 1 - 1 = 0
[2]     largest-smallest = 2 - 2 = 0
[1, 2]  largest-smallest = 2 - 1 = 1
Sum of the differences = 0 + 0 + 1 = 1
So, the resultant number is 1
Explanation 2:

All possible non-empty subsets are:
[3]         largest-smallest = 3 - 3 = 0
[5]         largest-smallest = 5 - 5 = 0
[10]        largest-smallest = 10 - 10 = 0
[3, 5]      largest-smallest = 5 - 3 = 2
[3, 10]     largest-smallest = 10 - 3 = 7
[5, 10]     largest-smallest = 10 - 5 = 5
[3, 5, 10]  largest-smallest = 10 - 3 = 7
Sum of the differences = 0 + 0 + 0 + 2 + 7 + 5 + 7 = 21
So, the resultant number is 21

*/

/**
 * @input A : Integer array
 *
 * @Output Integer
 */

import "sort"

func solve(A []int) int {
	// Brute Force:
	// Sort the given array: O(n logn)
	// Generate all subsequences. Using Backtracking either select or not select. O(n.2^n), n for copying and storing the final slices
	// Iterate over all the subsequences, and take the first and last element, because already sorted.
	// Take the difference and add that to the final answer. O(2^n)
	// TC will be O(n log n + n.2^n + 2^n)
	// Sc will be O(n) -> height of the recursive tree, O(n) for snapshot slice, O(2^n) for storing all the subsequences.
	// Sc will be O(n + n.2^n). 2^n subsets, each subsets on average n/2 elements. so n.2^n

	/*
	   Below solution is using contribution technique.
	   Finally we need the difference (sum of the largest number from all subsequences - sum of the smallest numbers from all subsequences)
	   For every i-th element, in (A[i] * number of subsequences A[i] is largest) - (A[i] * number of subsequences A[i] is largest)
	   We sort the array before starting the process.
	   So All subsequences starting with the i-th element the i-th element will be the smallest
	   And all subsequences ending with the i-th element the i-th element will be the largest

	   [3,5,10] -> in sorted order
	   all possible subsequences non-empty subsequences:
	   [3], [3,5], [3,5,10], [3,10], [5], [5,10], [10]

	   i=0, there are 2 elements after i-th element (n-i-1)
	   number of possible subsequences start with the i-th element= 2^(n-i-1) = 2^2 = 4 , in 4 subsequences the element A[i] => 3 is the smallest element
	   similarly there is 0 element before i-th element
	   number of possible subsequences ends with the i-th element = 2^i = 2^0 = 1, in one subsequence the element A[i] => 3 is the largest element

	   We can find the difference and add that to the final result.

	   Repeat this for all other elements.

	   Tc will be: O(n log n)
	   Sc will be: O(log n) => recursive stack for finding the power value

	*/

	sort.Ints(A)
	mod := 1000000007
	var sum int = 0
	n := len(A)

	for i := 0; i < len(A); i++ {
		cntSmallest := findPow(2, n-i-1, mod)
		cntLargest := findPow(2, i, mod)

		sum = (sum%mod + (cntLargest%mod*A[i]%mod)%mod - (cntSmallest%mod*A[i]%mod)%mod) % mod
	}

	return sum
}

func findPow(base int, exp int, mod int) int {
	if exp == 0 {
		return 1
	}

	powerRes := findPow(base, exp/2, mod)
	if exp%2 == 0 {
		return (powerRes % mod * powerRes % mod) % mod
	} else {
		return (powerRes % mod * powerRes % mod * base % mod) % mod
	}
}

// To avoid extra recursive stack space, we can use loop to build the power
/*
func power(x, y int) int {
    res := 1
    for y > 0 {
        if y % 2 != 0 {
            res = (res * x) % M
        }
        y = y >> 1 // => y/2
        x = (x * x) % M
    }
    return res % M
}

*/
