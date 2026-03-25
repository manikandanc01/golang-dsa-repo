package arraystwodimensional

/*
Problem Description
Given a binary sorted matrix A of size N x N. Find the row with the maximum number of 1.

NOTE:
If two rows have the maximum number of 1 then return the row which has a lower index.
Rows are numbered from top to bottom and columns are numbered from left to right.
Assume 0-based indexing.
Assume each row to be sorted by values.
Expected time complexity is O(rows + columns).

Problem Constraints
1 <= N<= 1000
0 <= A[i] <= 1

Input Format
The only argument given is the integer matrix A.

Output Format
Return the row with the maximum number of 1.

Example Input
Input 1:
 A = [   [0, 1, 1]
         [0, 0, 1]
         [0, 1, 1]   ]

Input 2:
 A =    [
         [0, 0, 0, 0]
         [0, 0, 0, 1]
         [0, 0, 1, 1]
         [0, 1, 1, 1]
		]


Example Output
Output 1:
 0

Output 2:
 3

Example Explanation
Explanation 1:
 Row 0 has maximum number of 1s.

Explanation 2:
 Row 3 has maximum number of 1s.
*/

/**
 * @input A : 2D integer array
 *
 * @Output Integer
 */
func solveRowWithMaximumOnes(A [][]int) int {
	// Brute Force : Iterate over each row and find the row with maximum no of ones.
	// Tc is O(n * n), Sc will be O(1)

	// Approach 2:
	// The matrix is sorted, we need to find the row with maximum one's.
	// One's will be at the end, 0's will be at the beginning of the row.
	// start from the top right corner.
	// 1 -> move left, 0 -> move downwards
	// Tc will be O(m+n), Sc is O(1)

	n := len(A)
	m := len(A[0])

	row := 0
	column := m - 1
	maxOnesRowIdx := -1

	for row < n && column >= 0 {
		if A[row][column] == 1 {
			maxOnesRowIdx = row
			column--
		} else {
			row++
		}
	}

	return maxOnesRowIdx
	/*
	   // Brute force approach code.

	   	n := len(A)
	   	m := len(A[0])

	   	maxOnesCount := -1
	   	maxOnesRow := -1

	   	for i:=0; i<n; i++{
	   	    onesCount := 0
	   	    for j:=0; j<m; j++{
	   	        if A[i][j] == 1{
	   	            onesCount++
	   	        }
	   	    }

	   	    if onesCount > maxOnesCount{
	   	        maxOnesCount = onesCount
	   	        maxOnesRow = i
	   	    }
	   	}

	   	return maxOnesRow
	*/
}
