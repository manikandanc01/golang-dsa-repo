package arraystwodimensional

/*
Problem Description
Given a 2D Matrix A of dimensions N*N, we need to return the sum of all possible submatrices.

Problem Constraints
1 <= N <=30
0 <= A[i][j] <= 10

Input Format
Single argument representing a 2-D array A of size N x N.

Output Format
Return an integer denoting the sum of all possible submatrices in the given matrix.

Example Input
Input 1:
A = [ [1, 1]
      [1, 1] ]

Input 2:
A = [ [1, 2]
      [3, 4] ]


Example Output
Output 1:
16
Output 2:
40


Example Explanation
Example 1:
Number of submatrices with 1 elements = 4, so sum of all such submatrices = 4 * 1 = 4
Number of submatrices with 2 elements = 4, so sum of all such submatrices = 4 * 2 = 8
Number of submatrices with 3 elements = 0
Number of submatrices with 4 elements = 1, so sum of such submatrix = 4
Total Sum = 4+8+4 = 16

Example 2:
The submatrices are [1], [2], [3], [4], [1, 2], [3, 4], [1, 3], [2, 4] and [[1, 2], [3, 4]].
Total sum = 40
*/

/**
 * @input A : 2D integer array
 *
 * @Output Integer
 */
func solve(A [][]int) int {

	// Brute Force: Consider all submatrices and find sum and return answer.
	// Total submatrices count: (n*(n+1))/2 * (m*(m+1))/2
	// Tc will be: O(n^2 * m^2), Sc will be O(1)

	// Better approach is using contribution technique.
	// In how many submatrices the ith element is present.
	// To build the matrices we need (TL, BR) pair or (TR, BL) pair. Let's choose (TL, BR)
	// TL -> Top right, BR -> Bottom right, TR -> Top right, BL -> Bottom Left
	// So number submarices the element present = total number of valid TL's * total number of valid BR's
	// [[1,2,3],
	//  [4,5,6],
	//  [7,8,9]]

	// A[1][1] => for this element, total valid TL's = 4, valid BR's = 4, total submatrices = 4*4 = 16
	//  [TL, TL,     3 ],
	//  [TL, [TL,BR],BR],
	//  [7,  BR,    BR]]

	// Tc will be O(n * m), Sc is O(1)

	rows := len(A)
	columns := len(A[0])
	submatricesSum := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			submatricesSum += A[i][j] * (i + 1) * (j + 1) * (rows - i) * (columns - j)
		}
	}

	return submatricesSum
}
