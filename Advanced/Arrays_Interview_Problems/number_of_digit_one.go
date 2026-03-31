package arraysinterviewproblems

/*
Problem Description
Given an integer A, find and return the total number of digit 1 appearing in all
non-negative integers less than or equal to A.

Problem Constraints
0 <= A <= 109

Input Format
The only argument given is the integer A.

Output Format
Return the total number of digit 1 appearing in all non-negative integers less than or equal to A.

Example Input
Input 1:
  A = 10

Input 2:
 A = 11

Example Output
Output 1:
 2
Output 2:
 4

Example Explanation
Explanation 1:
Digit 1 appears in 1 and 10 only and appears one time in each. So the answer is 2.
Explanation 2:
Digit 1 appears in 1(1 time) , 10(1 time) and 11(2 times) only. So the answer is 4.
*/

/**
 * @input A : Integer
 *
 * @Output Integer
 */
func solve(A int) int {
	// Brute Force: Iterate from 1 to A,
	// Check the each digit of a number and count no.of one's.
	// Tc will be O(A * log (maxNumber of digits)
	// For the given contraints, 0<=A<=10^9 this will give TLE.

	/*
	     Optimal Solution:
	      - Count no of 1's in one's place, ten's place, hundred's place and so on.

	      one's place:
	       - [1 -> 10] -> no of one's in one's place is 1.
	       - [2 -> 10] -> no of one's in one's place is 1.
	       - ...and so on.

	       Let's say if we want to find no of one's in one's place for the given number A.
	       = [A/10] * 1 + (A%10 > 0 -> add 1) (eg:  A = 88, A / 10 = 8, but there is one more extra one in 81, 88 % 10 = 8, this is > 0 so add 1. )

	      ten's place:
	        - [1 -> 100] -> no of one's comes in ten's place is 10. (10,11,12,13,14,15,16,17,18,19)
	        - [101 -> 200] -> no of one's comes in ten's place is 10. (110, 111, 112, 113, 114, .......119)

	        For the given number A, no of one's in ten's place is
	         = [A / 100] * 10 + min(max((A % 100 - (10-1)) ,0),10)


	      Eg:
	         A = 830 -> 830 / 100 = 8 * 10 = 80
	         => 830 % 100  -> 30 - 9 = 21, max(21, 0) => 21, min(21, 10) is 10.
	         => At max 10 is the maximum one's will come at the ten's place. So that we took min of max
	         => And then [10 -> 19] this range itself 10 number's will come.
	         => if A%100 => 19 => 19-9 = 10, [we can take 10 one's],
	         => if A%100 => 13 => 13-9 = 4, [we can add 4 one's]
	         => if A%100 => 10 => 10-9 = 1
	         => if A%100 => 8 => 8-9 = -1, so we do => max(A%100 - 9, 0) => max(-1, 0) => 0
	         => So to avoid negative values we took max.
	         => So to avoid values more than 10, we minimize it to 10.

	   General Formula:
	   i = 1
	       = [A/10] * 1 + min(max(A%10 - (1-1), 0), 1)

	   i = 10
	       = [A/100] * 10 + min(max(A%100 - (10-1), 0), 10)

	    in general:
	      = [A/i*10] * i + min(max(A%(i*10) - (i-1), 0), i)

	*/

	// Tc will be O(log A), Sc will be O(1)

	ans := 0
	for i := 1; i <= A; i = i * 10 {
		ans += (A/(i*10))*i + min(max(A%(i*10)-(i-1), 0), i)
	}

	return ans

}

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}

// 	return b
// }
