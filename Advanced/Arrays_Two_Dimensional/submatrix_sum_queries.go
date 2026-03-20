package arraystwodimensional

/*
Problem Description
Given a matrix of integers A of size N x M and multiple queries Q, for each query,
find and return the submatrix sum

Inputs to queries are top left (b, c) and bottom right (d, e) indexes of submatrix whose sum is to find out.

NOTE:
Rows are numbered from top to bottom, and columns are numbered from left to right.
The sum may be large, so return the answer mod 109 + 7.
Also, select the data type carefully, if you want to store the addition of some elements.
Indexing given in B, C, D, and E arrays is 1-based.
Top Left 0-based index = (B[i] - 1, C[i] - 1)
Bottom Right 0-based index = (D[i] - 1, E[i] - 1)

Problem Constraints
1 <= N, M <= 1000
-100000 <= A[i] <= 100000
1 <= Q <= 100000
1 <= B[i] <= D[i] <= N
1 <= C[i] <= E[i] <= M

Input Format
The first argument given is the integer matrix A.
The second argument given is the integer array B.
The third argument given is the integer array C.
The fourth argument given is the integer array D.
The fifth argument given is the integer array E.
(B[i], C[i]) represents the top left corner of the i'th query.
(D[i], E[i]) represents the bottom right corner of the i'th query.

Output Format
Return an integer array containing the submatrix sum for each query.

Example Input
Input 1:
 A = [   [1, 2, 3]
         [4, 5, 6]
         [7, 8, 9]   ]
 B = [1, 2]
 C = [1, 2]
 D = [2, 3]
 E = [2, 3]

Input 2:
 A = [   [5, 17, 100, 11]
         [0, 0,  2,   8]    ]
 B = [1, 1]
 C = [1, 4]
 D = [2, 2]
 E = [2, 4]

Example Output
Output 1:
 [12, 28]

Output 2:
 [22, 19]

Example Explanation
Explanation 1:
 For query 1: Submatrix contains elements: 1, 2, 4 and 5. So, their sum is 12.
 For query 2: Submatrix contains elements: 5, 6, 8 and 9. So, their sum is 28.

Explanation 2:
 For query 1: Submatrix contains elements: 5, 17, 0 and 0. So, their sum is 22.
 For query 2: Submatrix contains elements: 11 and 8. So, their sum is 19.
*/

/**
 * @input A : 2D integer array
 * @input B : Integer array
 * @input C : Integer array
 * @input D : Integer array
 * @input E : Integer array
 *
 * @Output Integer array.
 */
func solveSubmatrixSumQueries(A [][]int, B []int, C []int, D []int, E []int) []int {
	// Brute force: For each query, iterate from the range (B[i], C[i]) -> (D[i] -> E[i])
	// And find the sum.
	// Tc: O(Q * n * m), Sc is O(1) without considering the output else O(Q)
	// Iterating and finding the sum will give TLE

	// Better Approach: Use prefix sum to find the sum
	// psum[i][j] => sum of elements from (0,0) to (i,j)
	// let's say submatrix tl (a0,b0) and br (a1,b1)
	// sum = pfSum[a1][b1]-pfSum[a0-1][b1] - pfSum[a1][b0-1] + pfSum[a0-1][b0-1]
	// Handle the edge cases

	queriesSize := len(B)
	result := make([]int, queriesSize)
	prefixSum := calculatePrefixSum(len(A), len(A[0]), A)
	mod := int(1e9 + 7)

	for a := 0; a < queriesSize; a++ {
		topLeftRow := B[a] - 1
		topLeftColumn := C[a] - 1
		bottomRightRow := D[a] - 1
		bottomRightColumn := E[a] - 1

		sum := prefixSum[bottomRightRow][bottomRightColumn]

		if topLeftColumn > 0 {
			sum = (sum%mod - prefixSum[bottomRightRow][topLeftColumn-1]%mod) % mod
		}

		if topLeftRow > 0 {
			sum = (sum%mod - prefixSum[topLeftRow-1][bottomRightColumn]%mod) % mod
		}

		if topLeftColumn > 0 && topLeftRow > 0 {
			sum = (sum%mod + prefixSum[topLeftRow-1][topLeftColumn-1]%mod) % mod
		}

		result[a] = (sum + mod) % mod // to avoid negative value
	}

	return result
}

func calculatePrefixSum(n, m int, A [][]int) [][]int {
	pfSum := make([][]int, n)
	mod := int(1e9 + 7)
	for i := 0; i < n; i++ {
		pfSum[i] = make([]int, m)
	}

	// row wise sum
	for i := 0; i < n; i++ {
		sum := 0
		for j := 0; j < m; j++ {
			sum = (sum%mod + A[i][j]%mod) % mod
			pfSum[i][j] = sum
		}
	}

	// column wise prefix sum
	for j := 0; j < m; j++ {
		sum := 0
		for i := 0; i < n; i++ {
			sum = (sum%mod + pfSum[i][j]%mod) % mod
			pfSum[i][j] = sum
		}
	}

	return pfSum
}

/*

Brute Force Approach:

func solve(A [][]int , B []int , C []int , D []int , E []int )  ([]int) {
     Brute force: For each query, iterate from the range (B[i], C[i]) -> (D[i] -> E[i])
     And find the sum.
     Tc: O(Q * n * m), Sc is O(1) without considering the output else O(Q)

    queriesSize := len(B)
    result := make([]int,queriesSize)

    for a:=0; a<queriesSize; a++{
        topLeftRow := B[a]-1
        topLeftColumn := C[a]-1
        bottomRightRow := D[a]-1
        bottomRightColumn := E[a]-1

        sum := 0
        mod := int(1e9 + 7)
        for i:=topLeftRow; i<=bottomRightRow; i++{
            for j:=topLeftColumn; j<=bottomRightColumn; j++{
                sum = (sum % mod+A[i][j] %mod)%mod
            }
        }

        result[a] = sum
    }

    return result
}

*/
