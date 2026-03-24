package arraystwodimensional

/*
Problem Description
Given a row-wise and column-wise sorted matrix A of size N * M.
Return the maximum non-empty submatrix sum of this matrix.

Problem Constraints
1 <= N, M <= 1000
-109 <= A[i][j] <= 109

Input Format
The first argument is a 2D integer array A.

Output Format
Return a single integer that is the maximum non-empty submatrix sum of this matrix.

Example Input
Input 1:-
    -5 -4 -3
A = -1  2  3
     2  2  4

Input 2:-
    1 2 3
A = 4 5 6
    7 8 9


Example Output
Output 1:-
12

Output 2:-
45


Example Explanation
Expanation 1:-
The submatrix with max sum is
-1 2 3
 2 2 4
 Sum is 12.

Explanation 2:-
The largest submatrix with max sum is
1 2 3
4 5 6
7 8 9
The sum is 45.

*/

/**
 * @input A : 2D integer array
 *
 * @Output Long.
 */

import "math"

func solveMaximumSubmatrixSum(A [][]int) int64 {
	// Approach 1
	// Brute Force: Consider all submatrices.
	// for each submatrix, we need one top left corner and bottom right corner to uniquely identify one sub matrix
	// fix the top left (r1,c1) and bottom right(r2,c2) corner we need 4 loops, to find sum we need 2 more loops.
	// Tc will be (n^3 * m^3),Sc will be O(1) -> TLE

	// Approach 2:
	// Better Approach: To find sum we can use prefix sum.
	// Tc will be O(n^2 * m^2), Sc will be O(n * m) -> TLE

	// Approach 3:
	// Further Better approach: TL -> Starts from any row (0->n)
	// Ends at any row. (BR) -> Starts from start row.
	// Apply kadane's to find the maximum sum.
	// Tc is O(N^2 * M), and Sc is O(m)

	// Approach 4:
	// Since the given matrix is row wise and column wise sorted.
	// Let's say in 1d array if the array is sorted, maximum subarray sum will be the sum of one of the suffixes
	// Because the array is sorted the maximum values lies in the last.
	// But if we sum from beginning, that may give wrong answer because of the negative values present in the beginning.
	// Similarly we can do for 2D matrix.
	// Suffix matrix sum -> Suffix matrix is the matrix whose bottom right corner is always N*M th element.
	// Because the maximum element lies in the n-1, m-1 th place.
	// Iterate for the possibilities of Top left corner.
	// Tc will be O(N * M),

	// Approach 4 code
	n := len(A)
	m := len(A[0])

	suffixMatrix := make([][]int, n)
	for i := 0; i < n; i++ {
		suffixMatrix[i] = make([]int, m)
	}

	suffixMatrix[n-1][m-1] = A[n-1][m-1]
	ans := A[n-1][m-1]

	for j := m - 2; j >= 0; j-- {
		suffixMatrix[n-1][j] = suffixMatrix[n-1][j+1] + A[n-1][j]
		if suffixMatrix[n-1][j] > ans {
			ans = suffixMatrix[n-1][j]
		}
	}

	for i := n - 2; i >= 0; i-- {
		suffixMatrix[i][m-1] = suffixMatrix[i+1][m-1] + A[i][m-1]
		if suffixMatrix[i][m-1] > ans {
			ans = suffixMatrix[i][m-1]
		}
	}

	for i := n - 2; i >= 0; i-- {
		for j := m - 2; j >= 0; j-- {
			suffixMatrix[i][j] = A[i][j] + suffixMatrix[i+1][j] + suffixMatrix[i][j+1] - suffixMatrix[i+1][j+1]
			if suffixMatrix[i][j] > ans {
				ans = suffixMatrix[i][j]
			}
		}
	}

	return int64(ans)

	/*
	   // Approach 3 code

	    n := len(A)
	    m := len(A[0])

	    maxSum := int64(math.MinInt64)

	    for st := 0; st < n; st++{
	        a := make([]int,m)
	        for end := st; end<n; end++{
	            // add all elements to the array
	            for j:=0; j<m; j++{
	                a[j]+=A[end][j]
	            }

	            currMax := kadanesAlgo(a)
	            if currMax > maxSum{
	                maxSum = currMax
	            }
	        }
	    }

	    return maxSum

	*/

	/*
	   // Approach 2 code:

	    n := len(A)
	    m := len(A[0])

	    maxSum := int64(math.MinInt64)
	    prefixSum := buildPrefixSum(A)

	    // (r1,c1) -> top left
	    // (r2,c2) -> bottom right
	    for r1:=0; r1<n; r1++{
	       for c1:=0; c1<m; c1++{

	           sum := int64(0)
	           for r2:=r1; r2<n; r2++{
	               for c2:=c1; c2<m; c2++{

	                   sum = int64(prefixSum[r2][c2])
	                   if c1 !=0{
	                       sum -= int64(prefixSum[r2][c1-1])
	                   }

	                   if r1 !=0 {
	                       sum -= int64(prefixSum[r1-1][c2])
	                   }

	                   if r1!=0 && c1!=0{
	                       sum += int64(prefixSum[r1-1][c1-1])
	                   }


	                   if sum > maxSum{
	                       maxSum = sum
	                   }
	               }
	           }
	       }
	    }

	    return maxSum

	*/

	/*

	   // Approach 1: Code
	   n := len(A)
	   m := len(A[0])

	   maxSum := int64(math.MinInt64)

	   // (r1,c1) -> top left
	   // (r2,c2) -> bottom right
	   for r1:=0; r1<n; r1++{
	      for c1:=0; c1<m; c1++{

	          for r2:=r1; r2<n; r2++{
	              for c2:=c1; c2<m; c2++{

	                 sum := int64(0)
	                 for r:=r1; r<=r2; r++{
	                     for c:=c1; c<=c2; c++{
	                         sum += int64(A[r][c])
	                     }
	                 }

	                 if sum > maxSum{
	                     maxSum = sum
	                 }
	              }
	          }
	      }
	   }

	   return maxSum
	*/
}

func kadanesAlgo(A []int) int64 {
	ans := int64(math.MinInt64)
	sum := int64(0)

	for i := 0; i < len(A); i++ {
		sum += int64(A[i])
		if sum > ans {
			ans = sum
		}

		if sum < int64(0) {
			sum = int64(0)
		}
	}

	return ans
}

func buildPrefixSum(A [][]int) [][]int {
	n := len(A)
	m := len(A[0])

	pfSum := make([][]int, n)
	for i := 0; i < n; i++ {
		pfSum[i] = make([]int, m)
	}

	//  row sum
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			pfSum[i][j] = A[i][j]
			if j != 0 {
				pfSum[i][j] += pfSum[i][j-1]
			}
		}
	}

	// column sum
	for j := 0; j < m; j++ {
		for i := 0; i < n; i++ {
			if i != 0 {
				pfSum[i][j] += pfSum[i-1][j]
			}
		}
	}

	return pfSum
}
