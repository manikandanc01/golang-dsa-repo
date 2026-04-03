package bitmanipulations

/*
Problem Description
Given an array A. For every pair of indices i and j (i != j), find the maximum A[i] & A[j].

Problem Constraints
1 <= len(A) <= 105
1 <= A[i] <= 109

Input Format
The first argument is an integer array A.

Output Format
Return a single integer that is the maximum A[i] & A[j].

Example Input
Input 1:-
A = [53, 39, 88]

Input 2:-
A = [38, 44, 84, 12]

Example Output
Output 1:-
37

Output 2:-
36

Example Explanation
Explanation 1:-
53 & 39 = 37
39 & 88 = 0
53 & 88 = 16
Maximum among all these pairs is 37

Explanation 2:-
Maximum bitwise and among all pairs is (38, 44) = 36
*/

/**
 * @input A : Integer array
 *
 * @Output Integer
 */
func solveMaximumAndPair(A []int) int {
	/*
	   Brute Force:
	     - Consider all pairs.
	     - Take bitwise and of all and store the maximum.
	     - Tc will be O(n * n), Sc will be O(1)
	     - TLE

	   Optimal Approach:
	    - In order to get the maximum and value,
	    - We want to select 2 numbers, whose most significant bits(MSB's) are set.
	    - Start from the most significant bits, iterate over the entire array check the current bit is set or not.
	    - add the set bit numbers to the array.
	    - continue this process until the slice length is 2.
	    - If slice length is 2 we can break this loop.
	    - Tc will be O(n)
	*/

	// Optimal Approach Code
	// Here I used setBitElements as an extra space. We can avoid this.
	// In this problem, be careful about the operator priority.
	// (A[j] & 1<<i),
	// A[j] & (1<<i) Bit wise and & have more priority than the << left shift operator.
	// Tc will be O(n), Sc will be O(n)

	ans := 0
	for i := 32; i >= 0; i-- {
		setBitElements := make([]int, 0)
		for j := 0; j < len(A); j++ {
			if (A[j] & (1 << i)) > 0 {
				setBitElements = append(setBitElements, A[j])
			}
		}

		// fmt.Println(setBitElements, len(setBitElements))
		if len(setBitElements) >= 2 {
			ans += (1 << i)
			A = setBitElements
			/*
			   Instead of setting A to setbitelements.
			   we can iterate over A and set all the other elements to be 0.

			   for k := 0; k<n; k++{
			      if !(A[k]&(1<<i)) {
			        A[k]=0;
			      }
			    }
			*/
		}
	}

	return ans

	//  Brute Force Approach code:

	// n := len(A)
	// maximumAndValue := 0
	// for i:=0; i<n; i++{
	//     for j:= i+1; j<n; j++{
	//       andValue := (A[i] & A[j])
	//       if andValue > maximumAndValue{
	//           maximumAndValue = andValue
	//       }
	//     }
	// }

	// return maximumAndValue
}
