package recursion2

/*
Problem Description

Implement pow(A, B) % C.
In other words, given A, B and C, Find (AB % C).
Note: The remainders on division cannot be negative. In other words, make sure the answer you return is non-negative.


Problem Constraints

-109 <= A <= 109
0 <= B <= 109
1 <= C <= 109


Input Format

Given three integers A, B, C.


Output Format

Return an integer.


Example Input

Input 1:
A = 2
B = 3
C = 3
Input 2:
A = 3
B = 3
C = 1


Example Output

Output 1:
2
Output 2:
0


Example Explanation

Explanation 1:
23 % 3 = 8 % 3 = 2
Explanation 2:
33 % 1 = 27 % 1 = 0

*/

/*
Do not write code to include libraries, main() function or accept any input from the console.
Initialization code is already written and hidden from you. Do not write code for it again.
*/
/**
 * @input A : Integer
 * @input B : Integer
 * @input C : Integer
 *
 * @Output Integer
 */
 
func pow(A int, B int, C int) int {
	/*
	   Just write your code below to complete the function. Required input is available to you as the function arguments.
	   Do not print the result or any output. Just return the result via this function.
	*/

	if A == 0 {
		return 0
	}

	if B == 0 {
		return 1
	}

	x := pow(A, B/2, C) % C
	var ans int

	if B%2 == 0 {
		ans = (x % C * x % C) % C
	} else {
		ans = (x % C * x % C * A % C) % C
	}

	if ans < 0 {
		ans = (ans + C) % C
	}

	return ans

}
