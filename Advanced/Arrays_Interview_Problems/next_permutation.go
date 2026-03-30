package arraysinterviewproblems

/*
Problem Description:
Implement the next permutation, which rearranges numbers into the numerically
next greater permutation of numbers for a given array A of size N.
If such arrangement is not possible, it must be rearranged as the lowest possible order,
i.e., sorted in ascending order.

NOTE:
The replacement must be in-place, do not allocate extra memory.
DO NOT USE LIBRARY FUNCTION FOR NEXT PERMUTATION. Use of Library functions will disqualify your submission retroactively and will give you penalty points.

Problem Constraints
1 <= N <= 5 * 105
1 <= A[i] <= 109

Input Format
The first and the only argument of input has an array of integers, A.

Output Format
Return an array of integers, representing the next permutation of the given array.

Example Input
Input 1:
 A = [1, 2, 3]

Input 2:
 A = [3, 2, 1]

Example Output
Output 1:
  [1, 3, 2]

Output 2:
 [1, 2, 3]


Example Explanation
Explanation 1:
 Next permutaion of [1, 2, 3] will be [1, 3, 2].

Explanation 2:
 No arrangement is possible such that the number are arranged into the numerically next greater permutation of numbers.
 So will rearranges it in the lowest possible order.
*/

/**
 * @input A : Integer array
 *
 * @Output Integer array.
 */
func nextPermutation(A []int) []int {
	// Brute force : Generate all permutations, sort to get the lexicographical order
	// and find the next permutation.
	// Total: n! permutation will be there. and the time complexity will be n!*n

	// Optimal Approach:
	// See the notes.txt of this section to know about the algorithm

	n := len(A)
	i := n - 2

	// Find the first decreasing element from the last
	for i >= 0 && A[i] > A[i+1] {
		i--
	}

	// If entire array is in descending order, reverse the entire array and then return
	if i == -1 {
		return reverseArray(A, 0, n-1)
	}

	// Find the first greater element than the element at i from the last
	j := n - 1
	for j >= i && A[j] < A[i] {
		j--
	}

	// swap those elements
	A[i], A[j] = A[j], A[i]

	// reverse the suffix part and then return
	return reverseArray(A, i+1, n-1)
}

func reverseArray(A []int, i, j int) []int {
	for i <= j {
		A[i], A[j] = A[j], A[i]
		i++
		j--
	}

	return A
}

/*
Brute Force Code

package main
import "fmt"

// Generate all permutations using backtracking
func generate(nums []int, used []bool, current []int, result *[][]int) {
	if len(current) == len(nums) {
		temp := make([]int, len(current))
		copy(temp, current)
		*result = append(*result, temp)
		return
	}

	for i := 0; i < len(nums); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		current = append(current, nums[i])

		generate(nums, used, current, result)

		current = current[:len(current)-1]
		used[i] = false
	}
}

// Compare two permutations lexicographically
func isLess(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] < b[i] {
			return true
		} else if a[i] > b[i] {
			return false
		}
	}
	return false
}

// Simple bubble sort for permutations
func sortPermutations(perms [][]int) {
	n := len(perms)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if !isLess(perms[j], perms[j+1]) {
				perms[j], perms[j+1] = perms[j+1], perms[j]
			}
		}
	}
}

// Check if two slices are equal
func isEqual(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Brute force next permutation
func nextPermutation(nums []int) []int {
	var perms [][]int
	used := make([]bool, len(nums))

	// Step 1: generate all permutations
	generate(nums, used, []int{}, &perms)

	// Step 2: sort them
	sortPermutations(perms)

	// Step 3: find current permutation
	for i := 0; i < len(perms); i++ {
		if isEqual(perms[i], nums) {
			// Step 4: return next (or first if last)
			if i == len(perms)-1 {
				return perms[0]
			}
			return perms[i+1]
		}
	}

	return nums
}

func main() {
	nums := []int{1, 2, 3}
	fmt.Println(nextPermutation(nums))
}

Time Complexity: O(n! × n) (generation + comparisons)
Space Complexity: O(n! × n)

*/
