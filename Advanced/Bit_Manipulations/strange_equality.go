package bitmanipulations

/**
 * @input A : Integer
 *
 * @Output Integer

Problem Description
Given an integer A.
Two numbers, X and Y, are defined as follows:

X is the greatest number smaller than A such that the XOR sum of X and A is the same as the sum of X and A.
Y is the smallest number greater than A, such that the XOR sum of Y and A is the same as the sum of Y and A.
Find and return the XOR of X and Y.

NOTE 1: XOR of X and Y is defined as X ^ Y where '^' is the BITWISE XOR operator.
NOTE 2: Your code will be run against a maximum of 100000 Test Cases.

Problem Constraints
1 <= A <= 109

Input Format
First and only argument is an integer A.

Output Format
Return an integer denoting the XOR of X and Y.

Example Input
A = 5

Example Output
10

Example Explanation
The value of X will be 2 and the value of Y will be 8. The XOR of 2 and 8 is 10.
*/

// import "math"
func solveStrangeEquality(A int) int {
	/*
			   Optimal solution: After observation,
			   First have to find the last set bit.
			   2^(lastSetBit + 1) is the value for y, which is the smallest number greater than A.
			   and toggle the bits in A will give the greatest number which is less than A.
			   return x^y.

			   Tc will be O(1), Sc is O(1)

			   A + B = (A ^ B) + 2*(A & B), If A&B = 0, then A + B = A ^ B.

			   When you add two numbers in binary:

		          First, you add bits → this is XOR
		          Then, you handle carry → this comes from AND
		          Carry moves left → multiplying by 2

				  Sum = (sum without carry) + (carry shifted left)

				  A    B    A^B   A&B   Contribution
				  0    0     0     0      0
				  0    1     1     0      1
				  1    0     1     0      1
				  1    1     0     1      carry


				  Whenever we are adding 2 binary numbers: A^B -> will gives the actual sum without carry.
				  A & B -> Will gives the carry.
				  2 * (A ^ B) -> Will shift the Carry towards the left.


				  Let's say: A&B =  13 -> 1101, 13*2 = 26 -> 11010 Which is (13<<1)


	*/

	lastSetBit := 0
	x := A
	y := A
	for i := 31; i >= 0; i-- {
		// check i-th bit is set or not
		if (A & (1 << i)) > 0 {
			lastSetBit = i
			break
		}
	}

	for i := 0; i <= lastSetBit; i++ {
		// Toggle each bit. Parenthesis is important to ensure precedence of the operator.
		x = (x ^ (1 << i))
	}

	y = (1 << (lastSetBit + 1))

	return x ^ y

	/*
	   Brute Force: Try out all the possibilities
	*/

	// x := A-1
	// y := A+1
	// maxIntValue := math.MaxInt

	// for x > 0{
	//     if x^A == x+A{
	//         break
	//     }
	//     x--
	// }

	// for y <= maxIntValue {
	//     if y^A == y+A{
	//         break
	//     }
	//     y++
	// }

	// return x^y

}
