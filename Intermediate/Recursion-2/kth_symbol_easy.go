package recursion2

/**
Problem Description

On the first row, we write a 0. Now in every subsequent row, we look at the previous row and replace each occurrence of 0 with 01, and each occurrence of 1 with 10.

Given row number A and index B, return the Bth indexed symbol in row A. (The values of B are 0-indexed.).



Problem Constraints

1 <= A <= 20

0 <= B < 2A - 1



Input Format

First argument is an integer A.

Second argument is an integer B.



Output Format

Return an integer denoting the Bth indexed symbol in row A.



Example Input

Input 1:

 A = 3
 B = 0
Input 2:

 A = 4
 B = 4


Example Output

Output 1:

 0
Output 2:

 1


Example Explanation

Explanation 1:

 Row 1: 0
 Row 2: 01
 Row 3: 0110
Explanation 2:

 Row 1: 0
 Row 2: 01
 Row 3: 0110
 Row 4: 01101001
*/

func buildSlice(A int) []int {
	if A == 1 {
		return []int{0}
	}

	slice := buildSlice(A - 1)
	var newSlice []int

	for _, elem := range slice {
		switch elem {
		case 0:
			newSlice = append(newSlice, 0, 1)
			// newSlice = append(newSlice, 1)
		case 1:
			newSlice = append(newSlice, 1, 0)
			// newSlice = append(newSlice, 0)
		}
	}

	return newSlice
}

func solvekthsymbol(A int, B int) int {
	slice := buildSlice(A)
	return slice[B]
}

// Brute Force: Tc is O(2^A), because everytime, 2^1+2^2+2^3+.....+2^A times the loop runs, Total A times
// ignoring smaller exponential thing we can say Tc is O(2^A)
// Space Complexity: We are building the newSlice: O(2^A), max size of the slice is 2^A

// In go, tagged switch is more clear than if-else. By default go execute only one case, no fallback to other cases. 
