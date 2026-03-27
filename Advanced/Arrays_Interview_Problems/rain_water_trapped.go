package arraysinterviewproblems

/*
Problem Description

Imagine a histogram where the bars' heights are given by the array A. Each bar is of uniform width, which is 1 unit. When it rains, water will accumulate in the valleys between the bars.

Your task is to calculate the total amount of water that can be trapped in these valleys.

Example:

The Array A = [5, 4, 1, 4, 3, 2, 7] is visualized as below. The total amount of rain water trapped in A is 11.

Problem Constraints

1 <= |A| <= 105
0 <= A[i] <= 105



Input Format

First and only argument is the Integer Array, A.



Output Format

Return an Integer, denoting the total amount of water that can be trapped in these valleys



Example Input

Input 1:

 A = [0, 1, 0, 2]
Input 2:

A = [1, 2]


Example Output

Output 1:

1
Output 2:

0
*/

/**
 * @input A : Integer array
 *
 * @Output Integer
 */
func trap(A []int) int {
	/*
	   Brute Force: For each element, iterate over left side to find max,
	   and iterate over right side to find the maximum on the right.
	   because if there is nothing maximum than the current height the water won't trap.
	   Tc will be O(n^2), Sc will be O(1), TLE

	   Approach 2:
	    Instead of iterating and finding out the leftMax and rightMax, prestore that
	    Tc will be O(n), Sc will be O(n)

	   Approach 3: (Space Optimization)
	    - Using 2 pointer technique. (left -> 0, right -> n-1)
	    - leftMax -> indicates the leftMax of left pointer,
	    - rightMax -> indicates the rightMax of right pointer
	    - at any point, we only know the above thing,
	    - (we don't know the rightMax of the left pointer and leftMax of right pointer)

	     Now case 1:
	        leftMax <= rightMax
	           - rightMax is already greater than the leftMax of the left.
	           - which means rightMax will only be greater further when we iterate from right to left.
	           - we need min(leftMax[i], rightMax[i])
	           - so we can simply confirm that leftMax is the minimum height of the maximum''s
	           - and we can add  leftMax-A[left] to the ans

	      leftMax > rightMax
	          - leftMax is greater than the right max of right.
	          - leftMax will further increase in the range left -> right
	          - so at this point we can calculate the ans  rightMax - A[right]

	      Tc will be O(n) , Sc will be O(1)

	*/

	n := len(A)
	leftMax := 0
	rightMax := 0
	right := n - 1
	left := 0
	total := 0

	for left <= right {
		if leftMax <= rightMax {
			if leftMax-A[left] > 0 {
				total += leftMax - A[left]
			} else {
				leftMax = A[left]
			}
			left++
		} else {
			if rightMax-A[right] > 0 {
				total += rightMax - A[right]
			} else {
				rightMax = A[right]
			}
			right--
		}
	}

	return total

	/*n:= len(A)
	  leftMax := make([]int, n)
	  rightMax := make([]int, n)
	  totalWaterTrapped := 0

	  for i:=0; i<n; i++{
	      if i == 0{
	          leftMax[i] = A[i]
	      }else{
	          leftMax[i] = max(leftMax[i-1],A[i])
	      }
	  }

	  for i:=n-1; i>=0; i--{
	      if i==n-1{
	          rightMax[i] = A[i]
	      }else{
	          rightMax[i] = max(rightMax[i+1],A[i])
	      }
	  }

	  for i:=0; i<n; i++{
	      minOfMaximum := min(leftMax[i], rightMax[i])
	      totalWaterTrapped += minOfMaximum - A[i]
	  }

	  return totalWaterTrapped
	*/

	/*
	   // Approach 1: Code
	   n := len(A)
	   totalWaterTrapped := 0

	   for i:=0; i<n; i++{
	       leftMax := A[i]
	       rightMax := A[i]

	       for j:=i; j>=0; j--{
	           if A[j] > leftMax{
	               leftMax = A[j]
	           }
	       }

	       for j:=i; j<n; j++{
	           if A[j] > rightMax{
	               rightMax = A[j]
	           }
	       }

	       minOfMaximum := leftMax
	       if rightMax < minOfMaximum{
	           minOfMaximum = rightMax
	       }

	       totalWaterTrapped += minOfMaximum - A[i]
	   }

	   return totalWaterTrapped

	*/
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}
