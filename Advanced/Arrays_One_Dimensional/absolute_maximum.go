package arraysonedimensional

/**
 * @input A : Integer array
 * @input B : Integer array
 * @input C : Integer array
 * @input D : Integer array
 *
 * @Output Integer
 */

/*
Problem Description
Given 4 array of integers A, B, C and D of same size, find and return the maximum value of
| A [ i ] - A [ j ] | + | B [ i ] - B [ j ] | + | C [ i ] - C [ j ] | + | D [ i ] - D [ j ] | + | i - j|
where i != j and |x| denotes the absolute value of x.

Problem Constraints
2 <= length of the array A, B, C, D <= 100000
-106 <= A[i] <= 106

Input Format
The arguments given are the integer arrays A, B, C, D.

Output Format
Return the maximum value of
| A [ i ] - A [ j ] | + | B [ i ] - B [ j ] | + | C [ i ] - C [ j ] | + | D [ i ] - D [ j ] | + | i - j|

Example Input
Input 1:
 A = [1, 2, 3, 4]
 B = [-1, 4, 5, 6]
 C = [-10, 5, 3, -8]
 D = [-1, -9, -6, -10]

Input 2:
 A = [1, 2]
 B = [3, 6]
 C = [10, 11]
 D = [1, 6]

Example Output
Output 1:
 30

Output 2:
 11

Example Explanation
Explanation 1:
 Maximum value occurs for i = 0 and j = 1.

Explanation 2:
There is only one possibility for i, j as there are only two elements in the array.

*/

import "math"

func solveAbsoluteMaximum(A []int, B []int, C []int, D []int) int {
	/*
	   Brute Force: Total 32 Possibilities. Because of absolute.
	   Try all the possibilities find the min and maximum value and take the difference.

	   Tc O(n), Sc is O(1), because space is always 16 constant, it is not growing based on the input size.
	*/

	storage := make([][]int, 16)
	maxi := make([]int, 16)
	mini := make([]int, 16)

	for i := 0; i < 16; i++ {
		maxi[i] = math.MinInt
		mini[i] = math.MaxInt
	}

	for i := 0; i < len(A); i++ {
		storage[0] = append(storage[0], A[i]+B[i]+C[i]+D[i]+i)
		storage[1] = append(storage[1], A[i]+B[i]+C[i]+D[i]-i)
		storage[2] = append(storage[2], A[i]+B[i]+C[i]-D[i]+i)
		storage[3] = append(storage[3], A[i]+B[i]+C[i]-D[i]-i)
		storage[4] = append(storage[4], A[i]+B[i]-C[i]+D[i]+i)
		storage[5] = append(storage[5], A[i]+B[i]-C[i]+D[i]-i)
		storage[6] = append(storage[6], A[i]+B[i]-C[i]-D[i]+i)
		storage[7] = append(storage[7], A[i]+B[i]-C[i]-D[i]-i)
		storage[8] = append(storage[8], A[i]-B[i]+C[i]+D[i]+i)
		storage[9] = append(storage[9], A[i]-B[i]+C[i]+D[i]-i)
		storage[10] = append(storage[10], A[i]-B[i]+C[i]-D[i]+i)
		storage[11] = append(storage[11], A[i]-B[i]+C[i]-D[i]-i)
		storage[12] = append(storage[12], A[i]-B[i]-C[i]+D[i]+i)
		storage[13] = append(storage[13], A[i]-B[i]-C[i]+D[i]-i)
		storage[14] = append(storage[14], A[i]-B[i]-C[i]-D[i]+i)
		storage[15] = append(storage[15], A[i]-B[i]-C[i]-D[i]-i)

		if storage[0][len(storage[0])-1] > maxi[0] {
			maxi[0] = storage[0][len(storage[0])-1]
		}

		for j := 0; j < 16; j++ {
			if storage[j][len(storage[j])-1] > maxi[j] {
				maxi[j] = storage[j][len(storage[j])-1]
			}

			if storage[j][len(storage[j])-1] < mini[j] {
				mini[j] = storage[j][len(storage[j])-1]
			}
		}
	}

	result := math.MinInt
	for i := 0; i < 16; i++ {
		if (maxi[i] - mini[i]) > result {
			result = maxi[i] - mini[i]
		}
	}

	return result

}

/*
We can do the possibilities building by using the below logic. But the below one won't change the time or space complexity.


import "math"
func solve(A []int , B []int , C []int , D []int )  (int) {
    storage := make([][]int, 16)
    maxi := make([]int,16)
    mini := make([]int,16)

    for i:=0; i<16; i++{
        maxi[i] = math.MinInt
        mini[i] = math.MaxInt
    }

    for i:=0; i<len(A); i++{
        for j:=0; j<16; j++{
            currValue := A[i]
            for k:=0; k<5; k++{
                temp := 0
                if k==0{
                    temp = i
                }else if k==1{
                    temp = D[i]
                }else if k==2{
                    temp = C[i]
                }else if k==3{
                    temp = B[i]
                }

                if (j & (1<<k)) ==  0 {
                    currValue += temp
                }else{
                    currValue -= temp
                }
            }

            storage[j] = append(storage[j], currValue)

            if storage[j][len(storage[j])-1] > maxi[j]{
                maxi[j] = storage[j][len(storage[j])-1]
            }

            if storage[j][len(storage[j])-1] < mini[j]{
                mini[j] = storage[j][len(storage[j])-1]
            }
        }
    }

    result := math.MinInt
    for i:=0; i<16; i++{
        if (maxi[i] - mini[i]) > result{
            result = maxi[i] - mini[i]
        }
    }

    return result

}

*/
