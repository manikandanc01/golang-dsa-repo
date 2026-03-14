package arraysonedimensional

/*
Problem Description
You are given an array of N integers, A1, A2, .... AN.
Return the maximum value of f(i, j) for all 1 ≤ i, j ≤ N. f(i, j) is defined as |A[i] - A[j]| + |i - j|, where |x| denotes absolute value of x.

Problem Constraints
1 <= N <= 100000
-109 <= A[i] <= 109

Input Format
First argument is an integer array A of size N.

Output Format
Return an integer denoting the maximum value of f(i, j).

Example Input
Input 1:
A = [1, 3, -1]
Input 2:
A = [2]

Example Output
Output 1:
5
Output 2:
0

Example Explanation
Explanation 1:
f(1, 1) = f(2, 2) = f(3, 3) = 0
f(1, 2) = f(2, 1) = |1 - 3| + |1 - 2| = 3
f(1, 3) = f(3, 1) = |1 - (-1)| + |1 - 3| = 4
f(2, 3) = f(3, 2) = |3 - (-1)| + |2 - 3| = 5

So, we return 5.

Explanation 2:
Only possibility is i = 1 and j = 1. That gives answer = 0.

*/

/**
 * @input A : Integer array
 *
 * @Output Integer
 */

import "math"

func maxArr(A []int) int {
	//Brute force: Consider all pairs.
	// Tc is O(n^2), Sc is O(1)
	// TLE
	// maxi := int(math.MinInt64)
	// for i:=0; i< len(A); i++{
	//     for j:=i; j<len(A); j++{
	//         firstPart := A[i]-A[j]
	//         secondPart := i-j

	//         if firstPart < 0{
	//             firstPart *= -1
	//         }

	//         if secondPart < 0{
	//             secondPart *= -1
	//         }

	//         val := firstPart + secondPart
	//         if val > maxi{
	//             maxi = val
	//         }
	//     }
	// }

	// return maxi

	/*
	   Optimized approach:

	   |x| => x if x>0, -x if x<0
	   |A[i] - A[j]| => A[i] - A[j], for (A[i] - A[j]) > 0
	                 => A[j] - A[i], for (A[i] - A[j]) < 0

	   |i - j| => i - j, for (i - j) > 0
	           => j - i, for (i - j) < 0

	    1. A[i] - A[j] + i - j = (A[i] + i) - (A[j] + j)
	    2. A[i] - A[j] + j - i = (A[i] - i) - (A[j] - j)
	    3. A[j] - A[i] + i - j = (A[j] - j) - (A[i] - i)
	    4. A[j] - A[i] + j - i = (A[j] + j) - (A[i] + i)


	    In the above 4 possibilities, we have to think about 2 possibilities element + index, element - index
	    we have to find the max value and min value in the element + index sum to find the maximum possible difference
	    similarly we have to find the max and min value in the element - index sum
	    finally which differece is maximum we return that as answer.

	    Tc will be reduced to O(n), and Sc will be O(1)
	*/

	n := len(A)

	maxAbsoluteDifference := math.MinInt
	elementPlusIndexMax := math.MinInt
	elementMinusIndexMax := math.MinInt
	elementPlusIndexMin := math.MaxInt
	elementMinusIndexMin := math.MaxInt

	for i := 0; i < n; i++ {
		if (A[i] + i) > elementPlusIndexMax {
			elementPlusIndexMax = A[i] + i
		}

		if (A[i] + i) < elementPlusIndexMin {
			elementPlusIndexMin = A[i] + i
		}

		if (A[i] - i) > elementMinusIndexMax {
			elementMinusIndexMax = A[i] - i
		}

		if (A[i] - i) < elementMinusIndexMin {
			elementMinusIndexMin = A[i] - i
		}
	}

	if (elementPlusIndexMax - elementPlusIndexMin) > maxAbsoluteDifference {
		maxAbsoluteDifference = elementPlusIndexMax - elementPlusIndexMin
	}

	if (elementMinusIndexMax - elementMinusIndexMin) > maxAbsoluteDifference {
		maxAbsoluteDifference = elementMinusIndexMax - elementMinusIndexMin
	}

	return maxAbsoluteDifference
}
