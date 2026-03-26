package arraysinterviewproblems

/*
Problem Description
Given an unsorted integer array, A of size N. Find the first missing positive integer.
Note: Your algorithm should run in O(n) time and use constant space.

Problem Constraints
1 <= N <= 1000000
-109 <= A[i] <= 109

Input Format
First argument is an integer array A.

Output Format
Return an integer denoting the first missing positive integer.

Example Input
Input 1:
[1, 2, 0]

Input 2:
[3, 4, -1, 1]

Input 3:
[-8, -7, -6]

Example Output
Output 1:
3

Output 2:
2

Output 3:
1

Example Explanation
Explanation 1:
A = [1, 2, 0]
First positive integer missing from the array is 3.

Explanation 2:
A = [3, 4, -1, 1]
First positive integer missing from the array is 2.

Explanation 3:
A = [-8, -7, -6]
First positive integer missing from the array is 1.
*/

/**
 * @input A : Integer array
 *
 * @Output Integer
 */
func firstMissingPositive(A []int) int {
	/*
	   Approach 1:
	     Brute Force: Iterate from 1->n, search for the current element in the given array.
	     If it is not present return i.
	     If all elements from 1 to n is present in the array. then return n+1.

	     Tc will be O(n*n), Sc will be O(1)

	   Approach 2:
	     We can do search in O(1) using Hashset.
	     Tc will be reduced to O(n), Sc will be O(n)


	   Approach 3:
	     - Swap and put the elements in their respective position
	     - Iterate and find the missing integer.
	     - Each swap put atleast one element in its correct position.
	     - Iterate and figure out the missing elements.
	     - Tc will be O(n), Sc will be O(1)

	*/

	// Approach 3: Code
	n := len(A)

	i := 0
	for i < n {
		expectedElement := i + 1
		if A[i] == expectedElement || A[i] <= 0 || A[i] >= n {
			i++
			continue
		}

		temp := A[i]
		// To avoid infinite loop -> to handle duplicates -> Eg: [2,3,1,2]
		if A[i] == A[temp-1] {
			i++
			continue
		}
		A[i] = A[temp-1]
		A[temp-1] = temp
	}

	for i := 0; i < n; i++ {
		if A[i] != i+1 {
			return i + 1
		}
	}

	return n + 1

	/*
	   // Approach 2: Code
	   n := len(A)
	   hashSet := make(map[int]struct{})

	   for i:=0; i<n; i++{
	       hashSet[A[i]] = struct{}{}
	   }

	   for i:=1; i<=n; i++{
	       if _,ok := hashSet[i]; !ok{
	           return i
	       }
	   }

	   return n+1

	*/

	/*
	   // Approach 1: Code
	   n := len(A)
	   for i:=1; i<=n; i++{
	       //Search i in the given array.
	       isPresent := false
	       for j:=0; j<len(A); j++{
	          if A[j] == i{
	              isPresent = true
	              break
	          }
	       }

	       if !isPresent{
	           return i
	       }
	   }

	   return n+1
	*/
}
