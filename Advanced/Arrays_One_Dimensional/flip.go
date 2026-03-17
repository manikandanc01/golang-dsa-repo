package arraysonedimensional

/**
 * @input A : String
 *
 * @Output Integer array.
 */

/*
Problem Description:

You are given a binary string A(i.e., with characters 0 and 1) consisting of characters A1, A2, ..., AN. In a single operation, you can choose two indices, L and R, such that 1 ≤ L ≤ R ≤ N and flip the characters AL, AL+1, ..., AR. By flipping, we mean changing character 0 to 1 and vice-versa.
Your aim is to perform ATMOST one operation such that in the final string number of 1s is maximized.
If you don't want to perform the operation, return an empty array. Else, return an array consisting of two elements denoting L and R. If there are multiple solutions, return the lexicographically smallest pair of L and R.
NOTE: Pair (a, b) is lexicographically smaller than pair (c, d) if a < c or, if a == c and b < d.

Problem Constraints
1 <= size of string <= 100000

Input Format:
First and only argument is a string A.

Output Format
Return an array of integers denoting the answer.

Example Input
Input 1:
A = "010"

Input 2:
A = "111"

Example Output
Output 1:
[1, 1]

Output 2:
[]

Example Explanation
Explanation 1:
A = "010"
Pair of [L, R] | Final string
_______________|_____________
[1 1]          | "110"
[1 2]          | "100"
[1 3]          | "101"
[2 2]          | "000"
[2 3]          | "001"

We see that two pairs [1, 1] and [1, 3] give same number of 1s in final string. So, we return [1, 1].

Explanation 2:
No operation can give us more than three 1s in final string. So, we return empty array [].
*/

import "math"

func flip(A string) []int {
	/*
	   3. Best Approach

	   eg: [1,1,0,0,0,0,1,1,1,0,0,1]

	   - We need to find the linear solution.
	   - Greedy about maximum no of zero subarrays will not give us the correct answer.
	   - We need finally total gain's of 1's. If 0 -> (+1), 1 -> (-1)
	   - SO finally we need a subarray with maximum sum.
	   - We need the starting and end index of the subarray as well.

	   Kadane's algorithm, when the sum is less than 0, set the sum to -1
	   Finding the subarray with maximum gains

	   Tc is O(n), Sc is O(1)
	*/

	isAllOne := true
	for _, num := range A {
		if num == '0' {
			isAllOne = false
			break
		}
	}

	if isAllOne {
		return []int{}
	}

	n := len(A)
	maximumSum := math.MinInt
	startIdx := -1
	endIdx := -1
	sum := 0
	resetIdx := -1

	for i := 0; i < n; i++ {
		if A[i] == '0' {
			sum += 1
		} else {
			sum -= 1
		}

		if sum > maximumSum {
			maximumSum = sum
			endIdx = i
			startIdx = resetIdx + 1
		}

		if sum < 0 {
			sum = 0
			resetIdx = i
		}
	}

	return []int{startIdx + 1, endIdx + 1}

	/*
	   2. Better Approach:
	     1. First find the total number of ones in the given string array
	     2. And consider all possible L and R pair points.
	     3. If A[i] == 0, that will increase the total one's count in the total count else reduce one because of the flip.
	     4. And everytime record maximumOnesCount.

	     Tc will be O(n^2), Sc is O(1)
	*/

	// totalOnes := 0
	// for _, num := range A{
	//     if num == '1'{
	//         totalOnes++
	//     }
	// }

	// result := []int{}
	// maximumOnesCount := totalOnes
	// for i:=0; i<len(A); i++{
	//     onesCount := totalOnes
	//     for j:=i; j<len(A); j++{
	//         if A[j] == '0'{
	//             onesCount++
	//         }else{
	//             onesCount--
	//         }

	//         if onesCount > maximumOnesCount{
	//             result = []int{i+1,j+1}
	//             maximumOnesCount = onesCount
	//         }
	//     }
	// }

	// return result

	/*
	   1.  Brute Force Approach: Consider all substrings, and try flipping them in order to get the maximum ones
	     Tc will be: O(N^3) and Sc will be O(n), for that extra byte slice.
	     TLE
	*/

	// totalOnesCount := 0
	// for _, num := range A{
	//     if num == '1'{
	//        totalOnesCount++
	//     }
	// }

	// maximumOnesCount := totalOnesCount
	// leftIdx := -1
	// rightIdx := -1

	// for i:=0; i<len(A); i++{
	//     for j:=i; j<len(A); j++{
	//         byteSlice := []byte(A)
	//         for k:=i; k<=j; k++{
	//             if byteSlice[k] == '0'{
	//                 byteSlice[k] = '1'
	//             }else{
	//                 byteSlice[k] = '0'
	//             }
	//         }

	//         onesCount := 0
	//         for _, elem := range byteSlice{
	//            if elem == '1'{
	//                onesCount++
	//            }
	//         }

	//         if onesCount > maximumOnesCount{
	//             leftIdx = i+1
	//             rightIdx = j+1
	//             maximumOnesCount = onesCount
	//         }
	//     }
	// }

	// if leftIdx == -1 && rightIdx == -1{
	//     return []int{}
	// }

	// return []int{leftIdx, rightIdx}
}
