package arraystwodimensional

/*
Problem Description
Given a 2D integer matrix A of size N x N, find a B x B submatrix where B<= N and B>= 1, such that the sum of all the elements in the submatrix is maximum.

Problem Constraints
1 <= N <= 10^3.
1 <= B <= N
-10^2 <= A[i][j] <= 10^2.

Input Format
First arguement is an 2D integer matrix A.
Second argument is an integer B.

Output Format
Return a single integer denoting the maximum sum of submatrix of size B x B.

Example Input
Input 1:
 A = [
        [1, 1, 1, 1, 1]
        [2, 2, 2, 2, 2]
        [3, 8, 6, 7, 3]
        [4, 4, 4, 4, 4]
        [5, 5, 5, 5, 5]
     ]
 B = 3

Input 2:
 A = [
        [2, 2]
        [2, 2]
    ]
 B = 2

Example Output
Output 1:
 48

Output 2:
 8


Example Explanation
Explanation 1:
    Maximum sum 3 x 3 matrix is
    8 6 7
    4 4 4
    5 5 5
    Sum = 48

Explanation 2:
 Maximum sum 2 x 2 matrix is
  2 2
  2 2
  Sum = 8
*/

/**
 * @input A : 2D integer array
 * @input B : Integer
 *
 * @Output Integer
 */
import "math"

func solveMaximumSumSquareSubmatrix(A [][]int, B int) int {
	/*
	    Approach 1:
	     - Brute Force: For each Top left, there is only one possible bottom right.
	     - Consider all possible submatrices.
	     - Iterate over the submatrices and find the sum.
	     - Tc will be O(n^2 * m^2), Sc will be O(1), But TLE possibility

	   Approach 2:
	      - Find the sum in O(1) using Prefix sum.
	      - Tc will be reduced to O(n * m), Sc will be O(n * m)
	*/

	n := len(A)
	m := len(A[0])
	maximumSum := math.MinInt
	pfSum := buildPrefixSumArray(A)

	// top left corner
	for r1 := 0; r1 < n-B+1; r1++ {
		for c1 := 0; c1 < m-B+1; c1++ {

			// Bottom right corner Br
			r2 := r1 + B - 1
			c2 := c1 + B - 1

			sum := pfSum[r2][c2]
			if c1 != 0 {
				sum -= pfSum[r2][c1-1]
			}

			if r1 != 0 {
				sum -= pfSum[r1-1][c2]
			}

			if r1 != 0 && c1 != 0 {
				sum += pfSum[r1-1][c1-1]
			}

			if sum > maximumSum {
				maximumSum = sum
			}
		}
	}

	return maximumSum

	/*
	   // Approach 1: Code

	   n := len(A)
	   m := len(A[0])
	   maximumSum := math.MinInt

	   // top left corner
	   for r1:=0; r1<n-B+1; r1++{
	       for c1:=0; c1<m-B+1; c1++{

	          // Bottom right corner Br
	          r2 := r1+B-1
	          c2 := c1+B-1

	          // Iterate and find sum
	          sum := 0
	          for a:=r1; a<=r2; a++{
	              for b:=c1; b<=c2; b++{
	                  sum += A[a][b]
	              }
	          }

	          if sum > maximumSum{
	              maximumSum = sum
	          }
	       }
	   }

	   return maximumSum
	*/
}

func buildPrefixSumArray(A [][]int) [][]int {
	pfSum := make([][]int, len(A))
	for i := 0; i < len(A); i++ {
		pfSum[i] = make([]int, len(A[0]))
	}

	n := len(A)
	m := len(A[0])

	// row wise sum
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			pfSum[i][j] = A[i][j]
			if j != 0 {
				pfSum[i][j] += pfSum[i][j-1]
			}
		}
	}

	// column wise sum
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			if i != 0 {
				pfSum[i][j] += pfSum[i-1][j]
			}
		}
	}

	return pfSum
}
